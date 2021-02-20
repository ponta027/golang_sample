package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"regexp"

	page "ponta027.dip.jp/wiki"
)

/*
type Page struct {
	Title string
	Body  []byte
}
*/

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var templates = template.Must(template.ParseFiles("template/edit.html", "template/view.html"))

/**
 */
func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	templateFile := tmpl + ".html"
	//templateFile := "template/"+ tmpl + ".html"
	err := templates.ExecuteTemplate(w, templateFile, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil //
}

/*** */
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//	title := r.URL.Path[len("/view/"):]
	/*
		title, err := getTitle(w, r)
		if err != nil {
			return
		}
	*/
	p, err := page.LoadPage(string(title))
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//	title := r.URL.Path[len("/edit/"):]
	/*
		title, err := getTitle(w, r)
		if err != nil {
			return
		}
	*/
	p, err := page.LoadPage(string(title))
	if err != nil {
		p = &page.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//	title := r.URL.Path[len("/save/"):]
	/*
		title, err := getTitle(w, r)
		if err != nil {
			return
		}
	*/

	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/** */
/*
func (p *Page) save() error {
	filename := "log/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
*/

/*
func loadPage(title string) (*page.Page, error) {
	filename := "log/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page.Page{Title: title, Body: body}, nil
}
*/

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	/*
		p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
		p1.save()

		p2, _ := loadPage("TestPage")
		fmt.Println(string(p2.Body))
	*/
	http.HandleFunc("/view/", makeHandler(viewHandler))
	//http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", makeHandler(editHandler))
	//http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", makeHandler(saveHandler))
	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
