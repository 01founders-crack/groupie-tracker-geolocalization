package mapboxgeo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MapboxAPIResponse represents the structure of the Mapbox Geocoding API response.
type MapboxAPIResponse struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

// GetLocationCenterCoordinates retrieves the "center" coordinates for a location from the Mapbox Geocoding API.
func GetLocationCenterCoordinates(location string, accessToken string) ([]float64, error) {
	// Construct the API endpoint URL
	apiURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?proximity=ip&access_token=%s", location, accessToken)

	// Make the HTTP GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %s", response.Status)
	}

	// Decode the JSON response
	var data MapboxAPIResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Extract and return the "center" coordinates
	if len(data.Features) > 0 {
		return data.Features[0].Geometry.Coordinates, nil
	}

	return nil, fmt.Errorf("Location not found")
}



func ReturnLocationCoordinates(tempRelations map[string][]string, accessToken string) [][]float64  {
	//MAP DATA
    //location := "London, UK"
    // Call the function from your package
	var LocationsArr []string 
	for key := range tempRelations {
		LocationsArr = append(LocationsArr, key)
		fmt.Println(key,":::::::")
	}
	var CoordinatesArr [][]float64

	for _, location := range LocationsArr {
		coordinates, err := GetLocationCenterCoordinates(location, accessToken)
    if err != nil {
        fmt.Println("Error:", err)  
    } else {
		CoordinatesArr =append(CoordinatesArr, coordinates)
	}
    fmt.Printf("Center coordinates for %s: [%f, %f]\n", location, coordinates[0], coordinates[1])
	}
    

	return CoordinatesArr
}