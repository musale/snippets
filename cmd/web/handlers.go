package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/musale/snippets/pkg/forms"
	"github.com/musale/snippets/pkg/models"
)

// signupUserForm is used to render register a new user form
func (app *webApp) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign up page")
}

// signupUser is used to register a new user
func (app *webApp) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign up")
}

// loginUserForm is used to render a login form for a user
func (app *webApp) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log in page")
}

// loginUser is used to allow a registered user app access
func (app *webApp) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log in page")
}

// logoutUser is used to revoke auth
func (app *webApp) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log out page")
}

// home handles the homepage
func (app *webApp) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippets: s}
	app.render(w, r, "home.page.html", data)
}

// showSnippet displays a specific snippet
func (app *webApp) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil && id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippet: s}
	app.render(w, r, "show.page.html", data)
}

func (app *webApp) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")
	if !form.Valid() {
		data := &templateData{Form: form}
		app.render(w, r, "create.page.html", data)
		return
	}

	id, err := app.snippets.Insert(
		form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
	}

	msg := fmt.Sprintf("Snippet with id %d was created successfully!", id)
	app.session.Put(r, "flash", msg)
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}

func (app *webApp) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", &templateData{Form: forms.New(nil)})
}
