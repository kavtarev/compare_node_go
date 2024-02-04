package main

import (
	"os"
	"strings"
)

func ReadJson(name string) []byte {
	var builder strings.Builder
	builder.WriteString("../common/json/")
	builder.WriteString(name)
	builder.WriteString(".json")

	file, err := os.ReadFile(builder.String())

	if err != nil {
		panic("no file found")
	}

	return file
}