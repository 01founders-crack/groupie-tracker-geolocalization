// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"groupie-tracker/backend/models"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))

	http.HandleFunc("/", handleNotFound)
	http.HandleFunc("/group", handleID)
	http.HandleFunc("/500", handle500)

	go func() {
		_, err := handlers.GetArtists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Artists Fetched")
	}()

	port := "3000"
	println("Server listening on port http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

/*
func printRelatedDatesLocations(url string) {
	relations, err := handlers.GetRelations(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for location, relatedDatesLocations := range relations.DatesLocations {
		fmt.Printf("location: %s\n", location)
		fmt.Println("Related DatesLocations:")
		for _, relatedDateL := range relatedDatesLocations {
			fmt.Println(relatedDateL)
		}
	}
}
*/

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

func handleNotFound(w http.ResponseWriter, r *http.Request) {
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
	type RelationsData struct {
		ID             int
		DatesLocations map[string][]string
	}

	// Find the artist's relations by their ID
	// artistRelations, found := getArtistRelationsByID(combinedData, artistID)
	// if !found {
	//     fmt.Println("Artist relations not found for ID:", artistID)
	//     http.Error(w, "Artist relations not found", http.StatusNotFound)
	//     return
	// }
	var tempRelations map[string][]string
	var tempArtist models.Artist
	for _, artist := range combinedData.Artists {
		if strconv.Itoa(artist.ID) == artistID {
			//fmt.Println(combinedData.RelationsData[artist.ID], "::::ERKEKLERLE GEZDIM ABI")
			tempRelations = combinedData.RelationsData[artist.ID].DatesLocations

			//fmt.Println(artist, "::::ERKEKLERLE GEZDIM ABI")
			tempArtist = artist
		}
	}
	//fmt.Println(":::::ADSAADS",tempRelations,":::::::::ADSDSADS")
	// Pass the artist relations data to the template
	data := struct {
		GroupID        string
		Image          string
		Name           string
		Members        []string
		CreationDate   int
		FirstAlbum     string
		Locations      string
		ConcertDates   string
		Relations      string
		DatesLocations map[string][]string // Add a field for dates and locations
	}{
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
	}

	// Pass data to the 'group.html' template
	renderTemplateGroup(w, "group", data)
}
