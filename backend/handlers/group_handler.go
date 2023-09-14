package handlers

import (
	"fmt"
	"groupie-tracker/backend/external"
	"groupie-tracker/backend/helpers"
	"groupie-tracker/backend/mapboxgeo"
	"groupie-tracker/backend/models"
	"net/http"
	"strconv"
)

func HandleGroup(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID query parameter from the URL
	artistID := r.URL.Query().Get("id")
	intArtistID, err := strconv.Atoi(artistID)
	if err != nil {
		data := struct{}{}
		renderTemplate(w, "404", data)
	}
	if intArtistID > 0 && intArtistID < 53 {
		// Fetch the artist's data (if needed)
		combinedData, err := external.GetArtistsWithRelations()
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
		renderTemplate(w, "group", data)
	} else {
		data := struct{}{}
		renderTemplate(w, "404", data)
	}

}
