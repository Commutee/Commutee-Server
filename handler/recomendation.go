package handler

import (
	"commutee/types"
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

func LoadRecomendations() ([]types.Recomendation, error) {
	content, err := os.ReadFile("./data/recom.json")
	if err != nil {
		log.Fatal(err)
	}
	var dest []types.Recomendation
	err = json.Unmarshal(content, &dest)
	if err != nil {

		return nil, err
	}
	n := len(dest)

	indices := rand.Perm(n)

	subSliceSize := 4
	subSlice := make([]types.Recomendation, subSliceSize)
	for i := 0; i <= n-subSliceSize; i += subSliceSize {

		for j := 0; j < subSliceSize; j++ {
			subSlice[j] = dest[indices[i+j]]
		}
	}
	return subSlice, nil
}
