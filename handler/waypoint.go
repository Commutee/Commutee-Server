package handler

import (
	"commutee/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func WaypointCalc(waypoint string, mode string) (types.WaypointCollection, error) {
	url := fmt.Sprintf("https://api.geoapify.com/v1/routing?waypoints=%s&mode=%s&apiKey=2fde5f8bd5ff4f13a1a137bacbe0301c", waypoint, mode)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
		return types.WaypointCollection{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return types.WaypointCollection{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return types.WaypointCollection{}, err
	}
	var Instruction types.WaypointCollection
	err = json.Unmarshal(body, &Instruction)
	if err != nil {
		log.Fatal(err)
		return types.WaypointCollection{}, err
	}
	return Instruction, nil
}
