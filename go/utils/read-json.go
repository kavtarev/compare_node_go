package utils

import (
	"encoding/json"
	"os"
)

func ReadJson(name string) map[string]any {
	file, err := os.ReadFile("../common/json/" + name + ".json")

	if err != nil {
		panic("no file found")
	}

	var myMap map[string]any

	jErr := json.Unmarshal(file, &myMap)

	if jErr != nil {
		panic(jErr)
	}

	return myMap
}