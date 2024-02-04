package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var myMap map[string]any

func res(w http.ResponseWriter, req *http.Request) {
		mar, err := json.Marshal(myMap)

		if err != nil {
			panic("no file found")
		}

		w.Write(mar)
    // fmt.Fprintf(w, string(mar))
}


func main() {
		fmt.Println(os.Getpid())

		var builder strings.Builder
		builder.WriteString("../common/json")
		builder.WriteString(os.Args[1])
		builder.WriteString(".json")

		file, err := os.ReadFile(builder.String())

		if err != nil {
			panic("no file found")
		}

		jErr := json.Unmarshal(file, &myMap)

		if jErr != nil {
			panic(jErr)
		}

    http.HandleFunc("/", res)

    http.ListenAndServe(":3000", nil)
}