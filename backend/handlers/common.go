package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	//Template variables
	var t *template.Template
	var err error

	// Define the paths to the layout, header, footer, and the specific template
	layoutPath := filepath.Join("frontend", "layout.html")
	headerPath := filepath.Join("frontend", "components/header.html")
	footerPath := filepath.Join("frontend", "components/footer.html")
	templatePath := filepath.Join("frontend", tmpl+".html")

	// Add new one to here with if else $If you come from main go and added new page$
	// If you created only one html file for your page no need this area for more components
	if tmpl == "index" {
		musicCardPath := filepath.Join("frontend", "components/music_card.html")
		t = templateParseFiles(w, layoutPath, headerPath, footerPath, templatePath, musicCardPath)
	} else if tmpl == "group" {
		googleMapsPath := filepath.Join("frontend", "components/google_maps.html")
		groupInfoPath := filepath.Join("frontend", "components/group_info.html")
		t = templateParseFiles(w, layoutPath, headerPath, footerPath, templatePath, googleMapsPath, groupInfoPath)
	} else {
		t = templateParseFiles(w, layoutPath, headerPath, footerPath, templatePath)
	}

	// Execute the composed template
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func templateParseFiles(w http.ResponseWriter, filenames ...string) *template.Template {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// return
	}
	return t
}
