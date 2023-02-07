package handler

import (
	"commutee/types"
	"encoding/json"
	"os"
)

func LoadConfig() (types.Config, error) {

	var config types.Config
	conf, err := os.ReadFile("./config.json")
	if err != nil {
		return types.Config{}, err
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		return types.Config{}, err

	}

	return config, nil
}
