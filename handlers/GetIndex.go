package handlers

import (
	"html/template"
	"net/http"

	"github.com/sogilis/support-formation-docker/templates"
)

type GetIndex struct {
	BasePath string
}

func (i GetIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").Parse(templates.IndexPage)
	if err != nil {
		http.Error(w, "Unable to load index template", http.StatusInternalServerError)
		return
	}
	if err = t.ExecuteTemplate(w, "index", i.BasePath); err != nil {
		http.Error(w, "Unable execute index template", http.StatusInternalServerError)
		return
	}
}
