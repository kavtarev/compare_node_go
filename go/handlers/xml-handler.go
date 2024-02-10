package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func ClosureXmlHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("../common/xml/" + name + ".xml")
	if err != nil {
		panic("cant open xml file")
	}

	var obj TinyXml

	res := func(w http.ResponseWriter, r *http.Request) {
		xml.Unmarshal(file, &obj)
		fmt.Fprint(w, obj)
	}

	return res
}



type TinyXml struct {

}

type SmallXml struct {

}
