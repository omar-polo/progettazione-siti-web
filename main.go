package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"encoding/json"

	"github.com/fsnotify/fsnotify"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// flags
var (
	port    = flag.Int("p", 8000, `La porta da bindare per il server web.`)
	verbose = flag.Bool("v", false, `Stampa più log`)
)

var pages = make(map[string]*template.Template)

type Page struct {
	Title   string
	Content interface{}
	WCAGResult Section
}

type Section struct {
	Title       string            `json:"title"`
	Metadata    map[string]string `json:"metadata"`
	Content     string            `json:"-"`
	SubSections []Section         `json:"subsections"`
}

func (s Section) Source() string {
	return s.Metadata["source"]
}

func (s Section) Level() string {
	return s.Metadata["livello"]
}

func (s Section) Outcome() bool {
	return s.Metadata["outcome"] == "yes"
}

func (p Page) Results() string {
	j, err := json.Marshal(p.WCAGResult)
	if err != nil {
		log.Println("Cannot generate json", err)
		return ""
	}
	return string(j)
}

func initParser(file string) (Page, *bufio.Reader, *os.File, error) {
	p := Page{}

	f, err := os.Open(file)
	if err != nil {
		return p, nil, nil, err
	}

	reader := bufio.NewReader(f)
	title, err := reader.ReadString('\n')
	if err != nil {
		return p, reader, f, err
	}

	q := strings.Split(title, ":")
	if len(q) != 2 {
		return p, reader, f, fmt.Errorf("Cannot parse %q", title)
	}
	p.Title = strings.Trim(q[1], " \t\n")
	return p, reader, f, nil
}

func indexPage() (Page, error) {
	p, reader, f, err := initParser("src/index.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}
	p.Content, err = ioutil.ReadAll(reader)
	return p, err
}

func newSection() Section {
	s := Section{}
	s.Metadata = make(map[string]string)
	return s
}

func wcagParseRec(reader *bufio.Reader, firstline string, level, lineNo int) (Section, string, int, error) {
	type status int
	const (
		init status = iota
		metadata
		content
	)

	s := init
	cs := newSection()

	for {
		var line string

		if firstline != "" {
			line, firstline = firstline, ""
		} else {
			l, err := reader.ReadString('\n')
			if err == io.EOF {
				return cs, line, lineNo, nil
			} else if err != nil {
				return cs, line, lineNo, err
			}

			line = strings.Trim(l, "\n\r")
			lineNo++

			/*if *verbose {
				log.Printf("src/wcag.md:%v\t%q\n", lineNo, line)
			}*/
		}

		switch s {
		case init:
			var l int

			if line == "" {
				break
			} else if strings.HasPrefix(line, "###") {
				l = 3
			} else if strings.HasPrefix(line, "##") {
				l = 2
			} else if strings.HasPrefix(line, "#") {
				l = 1
			} else {
				err := fmt.Errorf("src/wcag.md:%d expected a title or an empty line", lineNo)
				return cs, line, lineNo, err
			}

			if l == level {
				// same level, continue!
				//fmt.Println("heading same level:", line)

				if cs.Title != "" {
					return cs, line, lineNo, nil
				}

				s = metadata
				t := strings.Fields(line)
				cs.Title = strings.Join(t[1:], " ")
			} else if l > level {
				// recursion
				//fmt.Println("heading level lower -> recursion", line)
				var ns Section
				var l string
				var err error

				ns, l, lineNo, err = wcagParseRec(reader, line, level+1, lineNo)
				cs.SubSections = append(cs.SubSections, ns)
				if err != nil {
					return cs, l, lineNo, err
				}
				firstline = l
			} else {
				//fmt.Println("heading level upper -> end recursion", line)
				return cs, line, lineNo, nil
			}

		case metadata:
			if line == "" {
				s = content

				if level == 3 {
					if _, ok := cs.Metadata["livello"]; !ok {
						fmt.Printf("pages/wcag.md:%v manca il livello\n", lineNo)
					}

					if _, ok := cs.Metadata["source"]; !ok {
						fmt.Printf("pages/wcag.md:%v manca la source\n", lineNo)
					}

					if _, ok := cs.Metadata["outcome"]; !ok {
						fmt.Printf("pages/wcag.md:%v manca l'outcome\n", lineNo)
					}
				}

				break
			}

			if strings.HasPrefix(line, "#") { // it's a title!
				//return cs, line, lineNo, nil
				s = init
				firstline = line
				break
			}

			fields := strings.Split(line, ":")
			if len(fields) != 2 {
				return cs, line, lineNo, fmt.Errorf("src/wcag.md:%d cannot parse metadata block", lineNo)
				/*if *verbose {
					//log.Println("line", lineNo, "assuming it's content")
				}

				s = content
				firstline = line
				break*/
			}

			k := strings.Trim(fields[0], " \t")
			v := strings.Trim(fields[1], " \t")
			cs.Metadata[k] = v

		case content:
			if strings.HasPrefix(line, "#") { // it's a title!
				//return cs, line, lineNo, nil
				s = init
				firstline = line
				break
			}

			cs.Content = cs.Content + "\n" + line
		}
	}
}

func wcagPage() (Page, error) {
	p, scanner, f, err := initParser("src/wcag.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}

	content, l, _, err := wcagParseRec(scanner, "", 0, 1)
	if err != nil && *verbose {
		log.Println("Malformed line:", l)
	}
	p.Content = content
	return p, err
}

func markdown(c string) template.HTML {
	return template.HTML(blackfriday.Run([]byte(c)))
}

func loadTemplate(name string) *template.Template {
	fmap := template.FuncMap{
		"markdown": markdown,
	}

	return template.Must(template.New("main").Funcs(fmap).ParseFiles("template.gohtml", name))
}

func loadTemplates() {
	ps := []string{"pages/index.gohtml", "pages/wcag.gohtml"}

	for _, p := range ps {
		pages[p] = loadTemplate(p)
	}
}

func copyFile(dst, src string) {
	if *verbose {
		log.Println("coping", src, "to", dst)
	}

	stats, err := os.Stat(src)
	if err != nil {
		log.Println("Cannot copy", src, "because:", err)
		return
	}

	if !stats.Mode().IsRegular() {
		log.Println("Cannot copy non-regular file", src)
		return
	}

	s, err := os.Open(src)
	if err != nil {
		log.Println("Cannot open", src)
		return
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		log.Println("Cannot create", dst)
		return
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		log.Println("Error copying", src, "to", dst, err)
	}
}

func copyDir(dst, src string) {
	os.Mkdir(dst, 0700)

	g, err := filepath.Glob(filepath.Join(src, "*"))
	if err != nil {
		log.Println("Cannot copy dir", src, "to", dst)
		return
	}

	for _, s := range g {
		d := filepath.Join(dst, filepath.Base(s))
		copyFile(d, s)
	}
}

func render(p Page, page string, res Section) {
	p.WCAGResult = res

	src := fmt.Sprintf("pages/%v", page)
	d := fmt.Sprintf("build/%v", page)

	r := regexp.MustCompile(`\.gohtml`)
	dst := r.ReplaceAllString(d, ".html")

	f, err := os.Create(dst)
	if err != nil {
		log.Println("Cannot write", dst)
		return
	}
	defer f.Close()

	t, ok := pages[src]
	if !ok {
		log.Println("cannot render page", src, "because it hasn't been loaded!")
		return
	}

	err = t.Execute(f, p)
	if err != nil {
		log.Println("Erorr rendering", src, err)
	}
}

func build() {
	loadTemplates()

	copyDir("build/js", "js")
	copyDir("build/css", "css")
	copyDir("build/img", "img")

	var p Page
	var err error
	var results Section

	// first render wcag, so we load all results for the search function
	if *verbose {
		log.Println("Compiling pages/wcag.gohtml")
	}
	p, err = wcagPage()
	if err != nil {
		log.Println("Cannot render pages/wcag.gohtml", err)
		return
	}
	results = p.Content.(Section)
	render(p, "wcag.gohtml", results)

	// render index.html
	if *verbose {
		log.Println("Compiling pages/index.gohtml")
	}
	p, err = indexPage()
	if err != nil {
		log.Println("Cannot render pages/index.gohtml", err)
		return
	}
	render(p, "index.gohtml", results)
}

func serve() {
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
	log.Println("listening on", *port, "...")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
	loop:
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					break loop
				}

				if *verbose {
					log.Println("Event:", event)
				}

				/*if event.Op&fsnotify.Write == fsnotify.Write {
					build()
				}*/
				log.Println("building...")
				build()

			case err, ok := <-watcher.Errors:
				if !ok {
					break loop
				}
				log.Println(err)

			case <-done:
				break loop
			}
		}

		done <- true
	}()

	ds := []string{".", "./js", "./css", "./src", "./pages", "./img"}

	for _, d := range ds {
		err = watcher.Add(d)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-done
}

func dev() {
	// first round of build
	build()

	go serve()
	watch()
}

func usage() {
	fmt.Fprintf(os.Stderr, "USAGE: %v [-h] [-p port] [-v] <subcommand>\n", os.Args[0])
	fmt.Fprintln(os.Stderr, " dove `subcommand' è:")
	fmt.Fprintln(os.Stderr, " - build :: compila il sito ed esce")
	fmt.Fprintln(os.Stderr, " - serve :: avvia un server locale")
	fmt.Fprintln(os.Stderr, " - dev   :: avvia il server web e ricompila dopo ogni cambiamento")
	fmt.Fprintln(os.Stderr, " - help  :: questo messaggio di aiuto")
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	switch args[0] {
	case "build":
		build()

	case "serve":
		serve()

	case "dev":
		dev()

	case "help":
		usage()

	default:
		usage()
		os.Exit(1)
	}
}
