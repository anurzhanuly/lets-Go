package main

import (
	"anurzhanuly/snippetbox/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"web:Project_mysql#65@/snippetbox?parseTime=true",
		"MySQL data source name",
	)
	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLogger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLogger,
		infoLog:  infoLogger,
		snippets: &mysql.SnippetModel{
			DB: db,
		},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	infoLogger.Printf("Starting server on : ", *addr)
	err = server.ListenAndServe()
	errorLogger.Println(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
