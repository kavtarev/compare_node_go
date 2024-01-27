package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var myMap map[string]any

func hello(w http.ResponseWriter, req *http.Request) {
		mar, err := json.Marshal(myMap)

		if err != nil {
			panic("no file found")
		}

    fmt.Fprintf(w, string(mar))
}


func main() {
		var builder strings.Builder
		builder.WriteString("../common/")
		builder.WriteString(os.Args[1])
		builder.WriteString(".json")

		file, err := os.ReadFile(builder.String())

		if err != nil {
			panic("no file found")
		}

		jErr := json.Unmarshal(file, &myMap)

		if jErr != nil {
			fmt.Println(jErr)
			panic("cant parse")
		}


    http.HandleFunc("/", hello)

    http.ListenAndServe(":3000", nil)
}