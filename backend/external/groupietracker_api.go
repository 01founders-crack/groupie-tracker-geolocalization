package external

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/backend/models"
)

func GetArtists() ([]models.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func GetArtistsWithRelations() (*models.CombinedData, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}

	relationsMap := make(map[int]*models.Relations)
	relationsChannel := make(chan *models.Relations, len(artists))
	errorChannel := make(chan error, len(artists))

	for i := range artists {
		go func(artist models.Artist) {
			relations, err := GetRelations(artist.Relations)
			if err != nil {
				errorChannel <- err
				return
			}
			relationsChannel <- relations
		}(artists[i])
	}

	for i := 0; i < len(artists); i++ {
		select {
		case err := <-errorChannel:
			return nil, err
		case relations := <-relationsChannel:
			relationsMap[artists[i].ID] = relations
		}
	}

	combinedData := &models.CombinedData{
		Artists:       artists,
		RelationsData: relationsMap,
	}

	return combinedData, nil
}


func GetRelations(url string) (*models.Relations, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var relations models.Relations
	if err := json.NewDecoder(response.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return &relations, nil
}
