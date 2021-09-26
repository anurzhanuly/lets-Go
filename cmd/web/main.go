package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	listeningPort := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on : ", *listeningPort)
	err := http.ListenAndServe(*listeningPort, mux)
	log.Fatal(err)
}
