package main

import (
	"bufio"
	"bytes"
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
}

type Section struct {
	Title       string
	Metadata    map[string]string
	Content     string
	SubSections []Section
}

func simplePage(file string) (Page, error) {
	p := Page{}

	f, err := os.Open(file)
	if err != nil {
		return p, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	title, err := reader.ReadString('\n')
	if err != nil {
		return p, err
	}

	q := strings.Split(title, ":")
	if len(q) != 2 {
		return p, fmt.Errorf("Cannot parse %q", title)
	}
	p.Title = strings.Trim(q[1], " \t\n")

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return p, err
	}

	p.Content = content

	return p, nil
}

func indexPage() (Page, error) {
	p, err := simplePage("src/index.md")
	if err != nil {
		return p, err
	}
	c := p.Content.([]byte)

	p.Content = blackfriday.Run(c)
	return p, nil
}

func newSection() Section {
	s := Section{}
	s.Metadata = make(map[string]string)
	return s
}

func wcagParseRec(scanner *bufio.Scanner, firstline string, level, lineNo int) ([]Section, string, int, error) {
	type status int
	const (
		init status = iota
		metadata
		content
	)

	s := init
	var ss []Section
	cs := newSection()

	for scanner.Scan() {
		var line string

		if firstline != "" {
			line = firstline
			firstline = ""
		} else {
			line = scanner.Text()
			lineNo++
		}

		if *verbose {
			log.Printf("src/wcag.md:%v\t%q\n", lineNo, line)
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
				return ss, line, lineNo, err
			}

			if l == level {
				// same level, continue!
				s = metadata
				t := strings.Fields(line)
				cs.Title = strings.Join(t[1:], " ")
			} else if l > level {
				// recursion
				var ns []Section
				var l string
				var err error

				ns, l, lineNo, err = wcagParseRec(scanner, line, level+1, lineNo)
				cs.SubSections = ns
				ss = append(ss, cs)
				cs = newSection()
				if err != nil {
					return ss, l, lineNo, err
				}
				firstline = l
			} else {
				return ss, line, lineNo, nil
			}

		case metadata:
			if line == "" {
				s = content
				break
			}

			if strings.HasPrefix(line, "#") { // it's a title!
				ss = append(ss, cs)
				firstline = line
				s = init
				break
			}

			fields := strings.Split(line, ":")
			if len(fields) != 2 {
				//return ss, line, lineNo, fmt.Errorf("src/wcag.md:%d cannot parse metadata block", lineNo)
				if *verbose {
					log.Println("line", lineNo, "assuming it's content")
				}

				s = content
				firstline = line
				break
			}

			k := strings.Trim(fields[0], " \t")
			v := strings.Trim(fields[1], " \t")
			cs.Metadata[k] = v

		case content:
			if strings.HasPrefix(line, "#") { // it's a title!
				ss = append(ss, cs)
				firstline = line
				s = init
				break
			}

			cs.Content = cs.Content + "\n" + line
		}
	}

	return ss, "", lineNo, scanner.Err()
}

func wcagPage() (Page, error) {
	p, err := simplePage("src/wcag.md")
	reader := bytes.NewReader(p.Content.([]byte))
	scanner := bufio.NewScanner(reader)

	var l string
	p.Content, l, _, err = wcagParseRec(scanner, "", 0, 1)
	if err != nil && *verbose {
		log.Println("Malformed line:", l)
	}
	return p, err
}

func loadTemplate(name string) *template.Template {
	return template.Must(template.New("main").ParseFiles("template.gohtml", name))
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

func render(p Page, page string) {
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
	copyDir("build/js", "js")
	copyDir("build/css", "css")
	//copyDir("build/img", "img")

	var p Page
	var err error

	// render index.html
	if *verbose {
		log.Println("Compiling pages/index.gohtml")
	}
	p, err = indexPage()
	if err != nil {
		log.Println("Cannot render pages/index.gohtml", err)
		return
	}
	render(p, "index.gohtml")

	// render wcag.html
	if *verbose {
		log.Println("Compiling pages/wcag.gohtml")
	}
	p, err = wcagPage()
	if err != nil {
		log.Println("Cannot render pages/wcag.gohtml", err)
		return
	}
	render(p, "wcag.gohtml")
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

				if event.Op&fsnotify.Write == fsnotify.Write {
					build()
				}

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

	ds := []string{".", "js", "css", "src", "pages"}

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
		loadTemplates()
		build()

	case "serve":
		serve()

	case "dev":
		loadTemplates()
		dev()

	case "help":
		usage()

	default:
		usage()
		os.Exit(1)
	}
}
