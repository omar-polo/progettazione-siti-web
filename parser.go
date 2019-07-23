// Qui sono presenti gli algoritmi e le strutture dati per effettuare
// il parsing della pagina del WCAG.

package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Parser definisce un semplice tipo `parser' su cui sono definite
// alcune funzioni utili. Con questa astrazione verrà poi costruito il
// parser per la pagina del WCAG.
type Parser struct {
	lineNumber int
	line       string // used for rollback
	lastLine   string // last line read
	reader     *bufio.Reader
	filename   string
	rollback   bool // a rollback was done
}

func trim(s string) string {
	return strings.TrimRight(s, " \t\n\r")
}

func trimm(s string) string {
	return strings.Trim(s, " \t\n\r")
}

func (p *Parser) ReadLine() (string, error) {
	p.lineNumber++
	p.rollback = false
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
	if p.rollback {
		panic(fmt.Errorf("Parser: cannot Rollback()!"))
	}

	p.lineNumber--
	p.line = p.lastLine
	p.lastLine = ""
	p.rollback = true
}

func (p *Parser) SyntaxError(reasons ...interface{}) error {
	return fmt.Errorf("%v:%v %v", p.filename, p.lineNumber, fmt.Sprint(reasons...))
}

func (p *Parser) ParseMetadata() (map[string]string, error) {
	m := make(map[string]string)

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

		k := trimm(f[0])
		v := trimm(f[1])
		m[k] = v
		//fmt.Println(p.SyntaxError(k, " -> ", v))
	}
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

// Questo workaround è necessario in quanto devo poter giocare con i
// puntatori, cosa che altrimenti non potrei fare.
func (w *WCAGParser) ReadLine() (string, error)                 { return (&w.p).ReadLine() }
func (w *WCAGParser) Rollback()                                 { (&w.p).Rollback() }
func (w *WCAGParser) SyntaxError(e ...interface{}) error        { return (&w.p).SyntaxError(e...) }
func (w *WCAGParser) ParseMetadata() (map[string]string, error) { return (&w.p).ParseMetadata() }

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
	m, err := p.ParseMetadata()

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

	return m, err
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

// User Testing

// Un Task rappresenta un task richiesto all'utente
type Task struct {
	Background string
	Timing     string
	Path       string
}

func TasksParse(p *Parser) ([]Task, error) {
	var t []Task

	// skip the intro
	TaskReadUntilSep(p)

	for {
		tt, err := TaskParse(p)
		t = append(t, tt)
		if err == io.EOF {
			return t, nil
		} else if err != nil {
			return t, err
		}
	}
}

func TaskParse(p *Parser) (Task, error) {
	t := Task{}
	tt := make(map[string]string)

	var err error
	var k, v string

	// 1. read all the key-value pairs
	for i := 0; i < 3; i++ {
		k, v, err = TaskKV(p)
		tt[k] = v
		if err != nil {
			break
		}
	}

	// 2. read up until the sep
	TaskReadUntilSep(p)

	// 3. reassemble the task and return

	var ok bool

	// 3.1 get the background
	t.Background, ok = tt["storia"]
	if !ok {
		return t, p.SyntaxError("missing `story' field!")
	}

	// 3.2 get the timing
	t.Timing, ok = tt["tempo"]
	if !ok {
		return t, p.SyntaxError("missing `tempo' field!")
	}

	// 3.3 get the path
	t.Path, ok = tt["percorso"]
	if !ok {
		return t, p.SyntaxError("missing `path' field!")
	}

	return t, err
}

var sepRe = regexp.MustCompile("^#+.*$")

func TaskKV(p *Parser) (key string, val string, err error) {
	var line string
	var l []string

	kRe := regexp.MustCompile(`^[a-zA-Z]+:$`)

	// 1. get the key
	for {
		line, err = p.ReadLine()

		if err != nil {
			break
		}

		if line == "" {
			continue
		}

		if kRe.MatchString(line) {
			// save the key, but drop the ':'
			key = line[:len(line)-1]
			break
		}
	}

	// 2. get the value
	for {
		line, err = p.ReadLine()

		if err != nil {
			break
		}

		if kRe.MatchString(line) || sepRe.MatchString(line) {
			p.Rollback()
			break
		}

		l = append(l, line)
	}

	// 3. join the value and return
	val = strings.Join(l, "\n")

	return
}

func TaskReadUntilSep(p *Parser) (string, error) {
	var l []string
	var err error
	var line string

	for {
		line, err = p.ReadLine()
		if err != nil {
			break
		}

		if sepRe.MatchString(line) {
			break
		}

		l = append(l, line)
	}

	return strings.Join(l, "\n"), err
}

// parser per i risultati

type Gravity int

const (
	Minor Gravity = iota
	Important
	Critic
	NA
)

func (g Gravity) String() string {
	switch g {
	case Minor:
		return "minore"
	case Important:
		return "importante"
	case Critic:
		return "critico"
	default:
		return "n/a"
	}
}

type Problem struct {
	Gravity Gravity
	Text    string
}

type UTResult struct {
	Outcome  string
	Level    string
	Timing   string
	Path     string
	Problems []Problem
	Opinions string
}

type UserResults struct {
	Metadata map[string]string
	Results  []UTResult
}

// parse the whole src/ut-results.md
func UTParse(p *Parser) ([]UserResults, error) {
	// skip the introduction
	TaskReadUntilSep(p)

	// Rewind so the next ReadLine will return "## utente X"
	p.Rollback()

	var urs []UserResults
	for {
		ur, err := UTUserResults(p)
		if err == io.EOF {
			return urs, nil
		} else if err != nil {
			return urs, err
		}
		urs = append(urs, ur)
	}
}

// parse a sigle user block
func UTUserResults(p *Parser) (UserResults, error) {
	ur := UserResults{}

	// the first line should be the heading
	line, err := p.ReadLine()
	if err != nil {
		return ur, err
	}

	if !sepRe.MatchString(line) {
		return ur, p.SyntaxError("Expecting a heading for the user")
	}

	// read the metadata
	ur.Metadata, err = p.ParseMetadata()
	if err != nil {
		return ur, err
	}

	for {
		// Peek next line
		line, err := p.ReadLine()
		if err != nil {
			return ur, err
		}
		if line == "" { // but skip empty lines
			continue
		}
		p.Rollback()

		// if it's a new user, return
		if strings.HasPrefix(line, "## ") {
			return ur, nil
		}

		// otherwise, read the next section
		r, err := UTParseResult(p)
		ur.Results = append(ur.Results, r)
		if err != nil {
			return ur, nil
		}
	}
}

// parse a single UTResult
func UTParseResult(p *Parser) (UTResult, error) {
	ut := UTResult{}
	m := make(map[string]string)

	// the first line should be the heading
	line, err := p.ReadLine()
	if err != nil {
		return ut, err
	}

	if !sepRe.MatchString(line) {
		return ut, p.SyntaxError("Expecting a heading for the task")
	}

	for {
		// Peek next line
		line, err = p.ReadLine()
		if err != nil {
			break
		}
		p.Rollback()

		// if it's a heading, exit loop
		if strings.HasPrefix(line, "##") {
			break
		}

		// otherwise read the KVs
		var k, v string
		k, v, err = TaskKV(p)
		if err != nil {
			break
		}
		m[k] = v
	}

	var ok bool

	if ut.Outcome, ok = m["esito"]; !ok {
		return ut, p.SyntaxError("missing `esito'")
	}

	if ut.Level, ok = m["livello"]; !ok {
		return ut, p.SyntaxError("missing `livello'")
	}

	if ut.Timing, ok = m["tempo"]; !ok {
		return ut, p.SyntaxError("missing `tempo'")
	}

	if ut.Path, ok = m["percorso"]; !ok {
		return ut, p.SyntaxError("missing `percorso'")
	}

	if _, ok = m["problemi"]; !ok {
		return ut, p.SyntaxError("missing `problemi'")
	}

	if ut.Opinions, ok = m["opinioni"]; !ok {
		return ut, p.SyntaxError("missing `opinioni'")
	}

	// fix "problemi"
	mre := regexp.MustCompile(`^(na|minore|importante|critico): (.*)$`)
	for _, l := range strings.Split(m["problemi"], "\n") {
		if l == "" {
			continue
		}

		if s := mre.FindStringSubmatch(l); s != nil {
			g := NA

			switch s[1] {
			case "minore":
				g = Minor
			case "importante":
				g = Important
			case "critico":
				g = Critic
			case "na":
				g = NA
			default:
				return ut, p.SyntaxError(fmt.Sprintf("invalid gravity %q", s[1]))
			}

			if g == NA {
				fmt.Println(p.SyntaxError("la gravità è n/a"))
			}

			ut.Problems = append(ut.Problems, Problem{
				Gravity: g,
				Text:    s[2],
			})
		} else {
			fmt.Println(p.SyntaxError("manca la gravità del problema"))
			ut.Problems = append(ut.Problems, Problem{
				Gravity: NA,
				Text:    l,
			})
		}
	}

	return ut, err
}
