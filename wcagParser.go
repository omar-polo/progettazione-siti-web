// Qui sono presenti gli algoritmi e le strutture dati per effettuare
// il parsing della pagina del WCAG.

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"regexp"
)

// Parser definisce un semplice tipo `parser' su cui sono definite
// alcune funzioni utili. Con questa astrazione verrà poi costruito il
// parser per la pagina del WCAG.
type Parser struct {
	lineNumber int
	line       string
	lastLine   string
	reader     *bufio.Reader
	filename   string
}

func trim(s string) string {
	return strings.TrimRight(s, " \t\n\r")
}

func (p *Parser) ReadLine() (string, error) {
	p.lineNumber++
	if p.line == "" {
		l, err := p.reader.ReadString('\n')
		l = trim(l)
		if err == nil {
			p.lastLine = l
		}
		return l, err
	} else {
		l := p.line
		p.line = ""
		p.lastLine = l
		return l, nil
	}
}

func (p *Parser) Rollback() {
	if p.lastLine == "" {
		panic(fmt.Errorf("Parser: cannot Rollback()!"))
	}

	p.lineNumber--
	p.line = p.lastLine
	p.lastLine = ""
}

func (p *Parser) SyntaxError(reasons ...interface{}) error {
	return fmt.Errorf("%v:%v %v", p.filename, p.lineNumber, fmt.Sprint(reasons...))
}

// La pagina del WCAG ha una grammatica abbastanza semplice:
//
//	document = Section*
//
//	Section = Title '\n' Metadata? '\n' content '\n' Subsections*
//	Subsections = Section*
//
//	Title = '#'+ + [[:print:]]*
//
//	Metadata = key ':' value '\n' Metadata?
//
// con il vincolo aggiunto che le sottosezioni abbiano i titoli di un
// livello inferiore a quello della sezione che li contiene.
//
// L'intero documento è racchiuso poi in una Section{} vuota che
// contiene le varie sezioni come SubSections (per comodità)

// Definiamo il tipo `Section' e un paio di funzioni di utility.

// Section rappresenta una sezione del report WCAG
type Section struct {
	Title       string            `json:"title"`
	Metadata    map[string]string `json:"metadata"`
	Quote       string            `json:"-"`
	Content     string            `json:"-"`
	SubSections []Section         `json:"subsections"`
}

func (s Section) Source() string {
	return s.Metadata["source"]
}

func (s Section) Level() string {
	return s.Metadata["livello"]
}

func (s Section) IsPassed() bool {
	return s.Metadata["outcome"] == "yes"
}

func (s Section) IsFailed() bool {
	return s.Metadata["outcome"] == "no"
}

func (s Section) IsNa() bool {
	return s.Metadata["outcome"] == "na"
}

// OutcomeClassName ritorna l'esito come stringa da usare come classe
// CSS.
func (s Section) OutcomeClassName() string {
	if s.IsPassed() {
		return "passed"
	} else if s.IsFailed() {
		return "failed"
	} else {
		return "na"
	}
}

// OutcomeNL ritorna l'esito come stringa in linguaggio naturale
// (italiano).
func (s Section) OutcomeNL() string {
	if s.IsPassed() {
		return "Positivo"
	} else if s.IsFailed() {
		return "Negativo"
	} else {
		return "Non applicabile"
	}
}

// Di seguito il recursive descent parser per la pagina del WCAG

// Estende il parser generale aggiungendo l'informazione relativa al
// livello corrente.
type WCAGParser struct {
	p     Parser
	Level int
}

func (w *WCAGParser) ReadLine() (string, error)          { return (&w.p).ReadLine() }
func (w *WCAGParser) Rollback()                          { (&w.p).Rollback() }
func (w *WCAGParser) SyntaxError(e ...interface{}) error { return (&w.p).SyntaxError(e...) }

// Effettua il parsing dell'intera pagina. La pagina viene vista come
// una sezione (vuota) contenente la lista di subsections.
func WCAGParse(p *WCAGParser) (Section, error) {
	/* // leggi tutte le sezioni
	s := Section{}

	//content, err := WCAGContent(p)
	//if err != nil {
	//	return s, err
	//}
	//s.Content = content

	//p.Level = 2

	for {
		ss, err := WCAGSection(p)

		s.SubSections = append(s.SubSections, ss)

		if err == io.EOF {
			return s, nil
		} else if err != nil {
			return s, err
		}
	}*/

	s, err := WCAGSection(p)
	if err != nil && err != io.EOF {
		return s, err
	}
	return s, nil
}

// WCAGSection legge (e ritorna) solo una sezione
func WCAGSection(p *WCAGParser) (Section, error) {
	s := Section{}

	// 1. Read title
	t, l, err := WCAGParseTitle(p)
	if err != nil {
		return s, err
	} else if l != p.Level {
		return s, p.SyntaxError("expected a title of level ", p.Level)
	}
	s.Title = t

	// 2. Read metadata
	m, err := WCAGMetadata(p)
	s.Metadata = m
	if err != nil {
		return s, err
	}
	
	// 3. Read quote
	q, err := WCAGQuote(p)
	s.Quote = q
	if err != nil {
		return s, err
	}

	// 4. Read content
	content, err := WCAGContent(p)
	s.Content = content
	if err != nil {
		return s, err
	}

	// 5. (try to) read subsections
	for {
		line, err := p.ReadLine()
		if err != nil {
			return s, nil
		}

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "#") {
			return s, p.SyntaxError("expected title")
		}

		p.Rollback()
		if level(line) > p.Level {
			//fmt.Println("subtitle-> recursion", line)
			p.Level++
			ss, err := WCAGSection(p)
			p.Level--

			s.SubSections = append(s.SubSections, ss)
			if err != nil {
				return s, err
			}
		} else {
			//fmt.Println("subtitle-> breaking", line)
			break
		}
	}

	return s, nil
}

// Assume s is a (markdown) title
func level(s string) int {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != '#' {
			return i
		}
	}
	return -1
}

// Assume s is a (markdown) title
func getTitle(s string) string {
	i := 0
	l := len(s)
	for ; i < l; i++ {
		if s[i] != '#' {
			break
		}
	}
	return s[i:]
}

func WCAGParseTitle(p *WCAGParser) (string, int, error) {
	for {
		l, err := p.ReadLine()
		if err != nil {
			return "", 0, err
		}

		if l == "" {
			continue
		}

		if !strings.HasPrefix(l, "#") {
			return "", 0, p.SyntaxError("expected a title")
		}

		return getTitle(l), level(l), nil
	}
}

func WCAGMetadata(p *WCAGParser) (map[string]string, error) {
	m := make(map[string]string)

	defer func() {
		if _, ok := m["source"]; !ok {
			fmt.Println(p.SyntaxError("missing `source'"))
		}

		if p.Level == 4 {
			if _, ok := m["livello"]; !ok {
				fmt.Println(p.SyntaxError("missing `livello'"))
			}
			if _, ok := m["outcome"]; !ok {
				fmt.Println(p.SyntaxError("missing `outcome'"))
			}
		}
	}()

	for {
		line, err := p.ReadLine()
		if err != nil {
			return m, err
		}

		if line == "" {
			return m, nil
		}

		f := strings.Split(line, ":")
		if len(f) != 2 {
			return m, p.SyntaxError("metadata should have two field:", len(f), "found")
		}

		k := trim(f[0])
		v := trim(f[1])
		m[k] = v
	}
}

var quoteRe = regexp.MustCompile(`\s*>\s*`)

func WCAGQuote(p *WCAGParser) (string, error) {
	var l []string
	var err error
	var line string

	for {
		line, err = p.ReadLine()
		if err != nil {
			break
		}

		if line == "" {
			l = append(l, line)
			continue
		}

		// stop when it's not a quote!
		if !quoteRe.MatchString(line) {
			p.Rollback()
			break
		}
		
		line = quoteRe.ReplaceAllString(line, "")
		l = append(l, line)
	}

	return strings.Join(l, "\n"), err
}

func WCAGContent(p *WCAGParser) (string, error) {
	var l []string
	var err error
	var line string

	for {
		line, err = p.ReadLine()
		if err != nil {
			break
		}

		// stops on title of level 1,2,3,4 but not 5 and onwards!
		if strings.HasPrefix(line, "#") && level(line) <= 4 {
			p.Rollback()
			break
		}

		l = append(l, line)
	}

	return strings.Join(l, "\n"), err
}
