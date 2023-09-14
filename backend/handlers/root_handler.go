package handlers

import (
	"fmt"
	"groupie-tracker/backend/external"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		combinedData, err := external.GetArtistsWithRelations()
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}
		renderTemplate(w, "index", combinedData)
	} else {
		data := struct{}{}
		renderTemplate(w, "404", data)
	}
}
