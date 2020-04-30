package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home handles the homepage
func (app *webApp) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// showSnippet displays a specific snippet
func (app *webApp) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Snippet is of id %d", id)
}

func (app *webApp) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
