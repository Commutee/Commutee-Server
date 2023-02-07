package handler

import (
	"commutee/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var client = &http.Client{}

func CoodinateCalc(place string) (string, error) {
	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=2fde5f8bd5ff4f13a1a137bacbe0301c", place)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	var featureCollection types.FeatureCollection
	err = json.Unmarshal(body, &featureCollection)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	feature := featureCollection.Features[0]
	latitude := feature.Geometry.Coordinates[1]
	longitude := feature.Geometry.Coordinates[0]

	return fmt.Sprintf("%f,%f", latitude, longitude), nil
}
