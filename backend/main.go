// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"groupie-tracker/backend/mapboxgeo"
	"groupie-tracker/backend/models"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))

	http.HandleFunc("/", handleNotFound)
	http.HandleFunc("/group", handleID)
	http.HandleFunc("/500", handle500)

	port := "3000"
	println("Server listening on port http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
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

	var tempRelations map[string][]string
	var tempArtist models.Artist
	for _, artist := range combinedData.Artists {
		if strconv.Itoa(artist.ID) == artistID {
			tempRelations = combinedData.RelationsData[artist.ID].DatesLocations
			tempArtist = artist
		}
	}

 // Load environment variables from .env file
 if err := loadEnv(".env"); err != nil {
	fmt.Println("Error loading .env file:", err)
	return
}
	accessToken := os.Getenv("ACCESS_TOKEN")
	gMapsToken := os.Getenv("GMAPS_TOKEN")

    // Check if the access token is empty or not set
    if accessToken == "" || gMapsToken == "" {
        fmt.Println("Access token not found in environment variable ACCESS_TOKEN or GMAPS_TOKEN")
        // Handle the case where the access token is missing or empty
        return
    }

	CoordinatesArr := mapboxgeo.ReturnLocationCoordinates(tempRelations,accessToken)

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
		CoordinatesArr []mapboxgeo.Location
		GMapsToken string
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
		CoordinatesArr: CoordinatesArr,
		GMapsToken: gMapsToken,
	}


	// Pass data to the 'group.html' template
	renderTemplateGroup(w, "group", data)
}


// Load environment variables from a file
func loadEnv(envFile string) error {
    content, err := os.ReadFile(envFile)
    if err != nil {
        return err
    }

    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        parts := strings.SplitN(line, "=", 2)
        if len(parts) == 2 {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            os.Setenv(key, value)
        }
    }

    return nil
}
