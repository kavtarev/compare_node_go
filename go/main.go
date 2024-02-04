package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"compare/handlers"
)

var myMap map[string]any

func main() {
		fmt.Println(os.Getpid())
		file := ReadJson(os.Args[1])

		jErr := json.Unmarshal(file, &myMap)

		if jErr != nil {
			panic(jErr)
		}

		jsonHandler := handlers.Closure(myMap)

    http.HandleFunc("/json-stringify", jsonHandler)

    http.ListenAndServe(":3000", nil)
}