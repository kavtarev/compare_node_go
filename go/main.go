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

    http.HandleFunc("/json-stringify", handlers.JsonStringifyResponse)

    http.ListenAndServe(":3000", nil)
}