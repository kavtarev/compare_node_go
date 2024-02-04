package handlers

import (
	"encoding/json"
	"net/http"
)

func JsonStringifyResponse(w http.ResponseWriter, req *http.Request) {
	var myMap map[string]any;

	mar, err := json.Marshal(myMap)

	if err != nil {
		panic("cant stringify")
	}

	w.Write(mar)
	// fmt.Fprintf(w, string(mar))
}