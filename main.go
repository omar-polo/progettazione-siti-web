package main

import (
	"bufio"
	"encoding/json"
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

// Definisce alcuni parametri da riga di comando
var (
	port    = flag.Int("p", 8000, `La porta da bindare per il server web.`)
	verbose = flag.Bool("v", false, `Stampa più log`)
)

// pages contiene i template
var pages = make(map[string]*template.Template)

// Page rappresenta una pagina Web
type Page struct {
	Title      string
	Content    interface{}
	WCAGResult Section
	HasSearch  bool
}

// Results ritorna una stringa contenente una rappresentazione Json
// dei possibili risultati della ricerca
func (p Page) Results() string {
	j, err := json.Marshal(p.WCAGResult)
	if err != nil {
		log.Println("Cannot generate json", err)
		return ""
	}
	return string(j)
}

// SkipCR implements io.Reader and automatically skips ONLY \r in the
// src reader.  Pretend it's UNIX everywhere!
type SkipCR struct {
	src io.Reader
}

// Read fills the p slice with data from SkipCR. n is the numer of bytes
// read and err is the eventual error.
func (s SkipCR) Read(p []byte) (n int, err error) {
	b := make([]byte, 1)
	n = 0

	for i := 0; i < len(p); i++ {
		r, err := s.src.Read(b)
		if err != nil {
			return n, err
		}

		if b[0] == '\r' { /* skip carriage returns */
			i--
			continue
		}

		if r == 0 {
			return n, nil
		}

		n += r
		p[i] = b[0]
	}
	return n, nil
}

// Inizializza il parser
func initParser(file string) (Page, *bufio.Reader, *os.File, error) {
	p := Page{}

	f, err := os.Open(file)
	if err != nil {
		return p, nil, nil, err
	}

	reader := bufio.NewReader(SkipCR{src: f})
	title, err := reader.ReadString('\n')
	if err != nil {
		return p, reader, f, err
	}

	q := strings.Split(title, ":")
	if len(q) != 2 {
		return p, reader, f, fmt.Errorf("Cannot parse %q", title)
	}
	p.Title = strings.Trim(q[1], " \t\n\r")
	return p, reader, f, nil
}

// Ritorna la pagina index
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

// Ritorna la pagina della mappa
func mappaPage() (Page, error) {
	p, reader, f, err := initParser("src/mappa.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}
	content, err := ioutil.ReadAll(reader)
	p.Content = string(content)
	return p, err
}

// Ritorna la pagina di info
func infoPage() (Page, error) {
	p, reader, f, err := initParser("src/info.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}
	content, err := ioutil.ReadAll(reader)
	p.Content = string(content)
	return p, err
}

// Ritorna la pagina dell'user testing
func utPage() (Page, error) {
	p, reader, f, err := initParser("src/ut.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}
	content, err := ioutil.ReadAll(reader)
	p.Content = string(content)
	return p, err
}

// Inizializza e ritorna una stuct section
func newSection() Section {
	s := Section{}
	s.Metadata = make(map[string]string)
	return s
}

// Ritorna la pagina del WCAG
func wcagPage() (Page, error) {
	p, reader, f, err := initParser("src/wcag.md")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return p, err
	}

	parser := WCAGParser{
		p: Parser{
			reader:     reader,
			filename:   "src/wcag.md",
			lineNumber: 1,
		},
		Level: 1,
	}

	s, err := WCAGParse(&parser)
	p.Content = s
	p.HasSearch = true
	return p, err
}

// Prende come input una stringa con testo markdown e ritorna la rappresentazione in HTML
func markdown(c string) template.HTML {
	return template.HTML(blackfriday.Run([]byte(c)))
}

// Carica il template con il nome dato
func loadTemplate(name string) *template.Template {
	fmap := template.FuncMap{
		"markdown": markdown,
	}

	return template.Must(template.New("main").Funcs(fmap).ParseFiles("template.gohtml", name))
}

// Carica tutti i template
func loadTemplates() {
	ps := []string{"pages/index.gohtml", "pages/wcag.gohtml", "pages/info.gohtml", "pages/mappa.gohtml", "pages/ut.gohtml"}

	for _, p := range ps {
		pages[p] = loadTemplate(p)
	}
}

// Funzione di utility che copia un file
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

// Funzione di utility che copia una Directory.
// Non copia ricorsivamente le sotto Directory
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

// Renderizza una pagina
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

// Compila tutto il sito
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

	// render mappa.html
	if *verbose {
		log.Println("Compiling pages/mappa.gohtml")
	}
	p, err = mappaPage()
	if err != nil {
		log.Println("Cannot render pages/mappa.gohtml", err)
		return
	}
	render(p, "mappa.gohtml", results)

	// render info.html
	if *verbose {
		log.Println("Compiling pages/info.gohtml")
	}
	p, err = infoPage()
	if err != nil {
		log.Println("Cannot render pages/info.gohtml", err)
		return
	}
	render(p, "info.gohtml", results)

	// render ut.html
	if *verbose {
		log.Println("Compiling pages/ut.gohtml")
	}
	p, err = utPage()
	if err != nil {
		log.Println("Cannot render pages/ut.gohtml", err)
		return
	}
	render(p, "ut.gohtml", results)
}

// Crea un Server Web che serve la cartella build
func serve() {
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)
	log.Println("listening on", *port, "...")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Invoca la funzione build quando i sorgenti vengono modificati
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

// Avvia sia il server Web che la funzione watch
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
	// Effettua il parsing degli argomenti da riga di comando
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
