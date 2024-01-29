package main

import (
	"log"
	"net/http"
)

func doSome(res http.ResponseWriter,  req *http.Request ) {
	res.Write([]byte("do some"))
}
func main() {

	http.HandleFunc("/", doSome)
	log.Println("up on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}