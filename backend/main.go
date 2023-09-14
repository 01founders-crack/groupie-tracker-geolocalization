// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"groupie-tracker/backend/helpers"
	"groupie-tracker/backend/mapboxgeo"
	"groupie-tracker/backend/models"

	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type RelationsData struct {
	ID             int
	DatesLocations map[string][]string
}

func main() {
	// Serve static files and set up routes
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/group", handleID)
	http.HandleFunc("/500", handle500)

	port := "443"
	println("Server listening on port https://localhost:" + port)

	// Serve the application over HTTPS with HTTP/2 support
	err := http.ListenAndServeTLS(":"+port, "certificates/server.crt", "certificates/server.key", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

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

func renderTemplateGroup(w http.ResponseWriter, tmpl string, data interface{}) {
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

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		combinedData, err := handlers.GetArtistsWithRelations()
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

func handle500(w http.ResponseWriter, r *http.Request) {
	data := struct{}{}
	renderTemplate(w, "500", data)
}

func handleID(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID query parameter from the URL
	artistID := r.URL.Query().Get("id")

	// Fetch the artist's data (if needed)
	combinedData, err := handlers.GetArtistsWithRelations()
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	var tempRelations map[string][]string
	var tempArtist models.Artist
	for _, artist := range combinedData.Artists {
		if strconv.Itoa(artist.ID) == artistID {
			tempRelations = combinedData.RelationsData[artist.ID].DatesLocations
			tempArtist = artist
		}
	}
	
	//accessToken, gMapsToken from .env file
	accessToken, gMapsToken := helpers.InitEnv()

	//ReturnLocationCoordinates
	CoordinatesArr := mapboxgeo.ReturnLocationCoordinates(tempRelations, accessToken)

	// Pass the artist relations data to the template
	data := models.ArtistData{
		GroupID:        artistID,
		Image:          tempArtist.Image,
		Name:           tempArtist.Name,
		Members:        tempArtist.Members,
		CreationDate:   tempArtist.CreationDate,
		FirstAlbum:     tempArtist.FirstAlbum,
		Locations:      tempArtist.Locations,
		ConcertDates:   tempArtist.ConcertDates,
		Relations:      tempArtist.Relations,
		DatesLocations: tempRelations, // Access dates and locations from artistRelations
		CoordinatesArr: CoordinatesArr,
		GMapsToken:     gMapsToken,
	}
	// Pass data to the 'group.html' template
	renderTemplateGroup(w, "group", data)
}
