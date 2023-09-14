package models

import "groupie-tracker/backend/mapboxgeo"

// Define the struct for artist data
type ArtistData struct {
	GroupID        string
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	Locations      string
	ConcertDates   string
	Relations      string
	DatesLocations map[string][]string
	CoordinatesArr []mapboxgeo.Location // Assuming mapboxgeo.Location is the correct type
	GMapsToken     string
}
