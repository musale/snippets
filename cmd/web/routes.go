package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *webApp) routes() http.Handler {
	mainMiddleware := alice.New(app.recoverPanic, app.logRequest,
		secureHeaders)
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	staticFileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))
	return mainMiddleware.Then(mux)
}
