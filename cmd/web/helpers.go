package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
	"github.com/musale/snippets/pkg/models"
)

// serverError writes an error to the errorLog then
// sends a generic 500 Internal Server Error response
func (app *webApp) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

// clientError sends a specific status code and description
func (app *webApp) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	return
}

// notFound for 404
func (app *webApp) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *webApp) render(w http.ResponseWriter, r *http.Request,
	name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

// addDefaultData injects common data into the templates
func (app *webApp) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.FlashMessage = app.session.PopString(r, "flash")
	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CSRFToken = nosurf.Token(r)
	return td
}

// The authenticatedUser method returns the ID of the current user from the
// session, or zero if the request is from an unauthenticated user.
func (app *webApp) authenticatedUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}
