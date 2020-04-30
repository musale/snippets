package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	addrString := fmt.Sprintf(":%s", *addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	staticFileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))

	server := &http.Server{
		Addr:     addrString,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server listening on port %s", addrString)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
