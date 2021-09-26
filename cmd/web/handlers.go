package main

import (
	"fmt"
	template2 "html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)

		return
	}

	templates := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	template, err := template2.ParseFiles(templates...)
	if err != nil {
		app.serverError(w, err)

		return
	}

	err = template.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)

		return
	}

	_, err = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	if err != nil {
		return
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}

	_, err := w.Write([]byte("createSnippet"))
	if err != nil {
		app.errorLog.Println(err.Error())
	}
}
