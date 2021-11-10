package main

import (
	"html/template"
	"path/filepath"
	"time"

	"a4lab2.net/snipb/pkg/forms"
	"a4lab2.net/snipb/pkg/models"
)

type templateData struct {
	AuthenticatedUser *models.User
	CSRFToken         string
	Flash             string
	CurrentYear       int
	Form              *forms.Form
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02,Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tpl"))

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
