package handlers

import (
	"encoding/json"
	"net/http"
)

func ClosureJson(obj map[string]any) func(w http.ResponseWriter, req *http.Request) {
	res := func(w http.ResponseWriter, req *http.Request) {

		mar, err := json.Marshal(obj)
		if err != nil {
			panic("cant stringify")
		}

		w.Write(mar)
	}

	return res
}
