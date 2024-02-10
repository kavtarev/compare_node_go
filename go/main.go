package main

import (
	"compare/handlers"
	"compare/utils"
	"fmt"
	"net/http"
	"os"
)

func main() {
		fmt.Println(os.Getpid())
		myMap := utils.ReadJson(os.Args[1])

		jsonHandler := handlers.ClosureJson(myMap)
		downloadFileHandler := handlers.ClosureReadFileHandlerChunk(os.Args[1])

    http.HandleFunc("/json-stringify", jsonHandler)
    http.HandleFunc("/download-file", downloadFileHandler)

    http.ListenAndServe(":3000", nil)
}