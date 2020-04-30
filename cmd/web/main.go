package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Snippet is of id %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
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
