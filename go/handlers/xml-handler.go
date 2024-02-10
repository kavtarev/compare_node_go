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

	var obj SmallXml

	res := func(w http.ResponseWriter, r *http.Request) {
		xml.Unmarshal(file, &obj)
		fmt.Fprint(w, obj)
	}

	return res
}

type TinyXml struct {
	XMLName xml.Name `xml:"note"`
	To string`xml:"to"`
	From 	string `xml:"from"`
	Heading 	string `xml:"heading"`
	Body 	string `xml:"body"`
}

type SmallXml struct {
	XMLName xml.Name `xml:"breakfast_menu"`
	Food []Food `xml:"food"`
}

type Food struct {
	Name string `xml:"name"`
	Price string `xml:"price"`
	Description string `xml:"description"`
	Calories string `xml:"calories"`
}
