package main

import (
	"log"
	"net/http"
)

// home handles the homepage
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from portland"))
}

// showSnippet displays a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a specific request"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting portland server")
	err := http.ListenAndServe(":4050", mux)
	log.Fatal(err)
}
