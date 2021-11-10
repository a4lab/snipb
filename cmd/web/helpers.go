package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"a4lab2.net/snipb/pkg/models"
	"github.com/justinas/nosurf"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.cache[name]

	if !ok {
		app.serverError(w, fmt.Errorf("Template:(%s)  not found", name))
	}
	buf := new(bytes.Buffer)
	err := ts.Execute(w, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CurrentYear = time.Now().Year()
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.session.PopString(r, "flash")
	return td
}
func (app *application) authenticatedUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}
