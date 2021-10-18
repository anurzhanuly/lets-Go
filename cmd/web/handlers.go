package main

import (
	"anurzhanuly/snippetbox/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: s}

	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Create an instance of a templateData struct holding the snippet data.
	data := &templateData{Snippet: s}

	app.render(w, r, "show.page.tmpl", data)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errorsMap := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errorsMap["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errorsMap["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errorsMap["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errorsMap["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errorsMap["expires"] = "This field is invalid"
	}

	if len(errorsMap) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errorsMap,
			FormData:   r.PostForm,
		})

		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}
