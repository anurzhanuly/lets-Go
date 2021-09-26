package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLogger,
		infoLog:  infoLogger,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	infoLogger.Printf("Starting server on : ", *addr)
	err := server.ListenAndServe()
	errorLogger.Println(err)
}
