package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	// "compare/handlers"
)

var myMap map[string]any

func readFileHandlerChunk(w http.ResponseWriter, r *http.Request) {
	const size = 50000;

	f, err := os.Open("../common/files/small.txt")

	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	bytes := make([]byte, size)

	for {
		bytesread, err := f.Read(bytes)

		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
		}

		w.Write(bytes[:bytesread])

	}

}

func main() {
		fmt.Println(os.Getpid())
		// file := ReadJson(os.Args[1])

		// jErr := json.Unmarshal(file, &myMap)

		// if jErr != nil {
		// 	panic(jErr)
		// }

		// jsonHandler := handlers.ClosureJson(myMap)

    // http.HandleFunc("/json-stringify", jsonHandler)
    http.HandleFunc("/download-files", readFileHandlerChunk)

    http.ListenAndServe(":3000", nil)
}