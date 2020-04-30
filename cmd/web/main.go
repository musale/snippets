package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	addrString := fmt.Sprintf(":%s", *addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	staticFileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))

	log.Printf("Starting server listening on port %s", addrString)
	err := http.ListenAndServe(addrString, mux)
	log.Fatal(err)
}
