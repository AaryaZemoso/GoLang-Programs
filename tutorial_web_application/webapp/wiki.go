package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		title, err := getTitle(res, *req)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		fn(res, req, title)
	}
}

func getTitle(w http.ResponseWriter, r http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, &r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func renderTemplate(res http.ResponseWriter, title string, p *Page) {

	err := templates.ExecuteTemplate(res, title+".html", p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

}

func saveHandler(res http.ResponseWriter, req *http.Request, title string) {
	b := req.FormValue("body")
	p := &Page{title, []byte(b)}
	err := p.save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/view/"+title, http.StatusFound)
}

func editHandler(res http.ResponseWriter, req *http.Request, title string) {

	page, err := loadPage(title)

	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(res, "edit", page)
}

func viewHandler(res http.ResponseWriter, req *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(res, req, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(res, "view", p)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{title, body}, nil
}

func main() {
	// p1 := &Page{"test1", []byte("This is a test body")}
	// p1.save()

	// p2, err := loadPage("test1")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(p2.Body))
	// }

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
