package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"compare/handlers"
	"compare/utils"
)

var myMap map[string]any

func main() {
		fmt.Println(os.Getpid())
		file := utils.ReadJson(os.Args[1])

		jErr := json.Unmarshal(file, &myMap)

		if jErr != nil {
			panic(jErr)
		}

		jsonHandler := handlers.ClosureJson(myMap)
		readFileHandler := handlers.ClosureReadFileHandlerChunk(os.Args[1])

    http.HandleFunc("/json-stringify", jsonHandler)
    http.HandleFunc("/download-files", readFileHandler)

    http.ListenAndServe(":3000", nil)
}