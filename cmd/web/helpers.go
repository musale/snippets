package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
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
