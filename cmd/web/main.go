package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from portland"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting portland server")
	err := http.ListenAndServe(":4050", mux)
	log.Fatal(err)
}
