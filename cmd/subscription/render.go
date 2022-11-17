package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

var pathToTemplates = "./cmd/subscription/templates"

type TemplateData struct {
	StringMap         map[string]string
	IntMap            map[string]int
	FloatMap          map[string]float32
	Data              map[string]any
	Flash             string
	Warning           string
	Error             string
	AuthenticatedUser int
	Now               time.Time
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, name string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s.gohtml", pathToTemplates, name))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)

		if td == nil {
			td = &TemplateData{}
		}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	if app.isAuthenticated(r) {
		td.AuthenticatedUser = app.Session.GetInt(r.Context(), "authenticatedUser")
	}
	td.Now = time.Now()
	return td
}

func (app *Config) isAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "authenticatedUserID")
	return exists
}
