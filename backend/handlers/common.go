package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = filepath.Join("frontend", tmpl+".html")
	layout := filepath.Join("frontend", "layout.html")
	t, err := template.ParseFiles(layout, tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
