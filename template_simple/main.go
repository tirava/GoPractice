package main

import (
	"html/template"
	"net/http"
)

type page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &page{
		Title:   "My title",
		Content: "My content",
	}
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
