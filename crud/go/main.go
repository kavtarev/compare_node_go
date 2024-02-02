package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "compare_node_go"
)

func run(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hi"))
}

func main () {
	db, errCon := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/compare_node_go?sslmode=disable")

	if errCon != nil {
		log.Println("cant connect to db")
	}
	defer db.Close()

	rows, qErr := db.Query("select count(1) from users")

	if qErr != nil {
		log.Println("query error", qErr)
	}
	defer rows.Close()

	var res int;

	for rows.Next() {
		err := rows.Scan(&res)

		if err != nil {
			fmt.Println("scan error")
		}

		fmt.Println(6666, res)
	}

	// http.HandleFunc("/", run)
	// log.Println("before init")

	// err := http.ListenAndServe(":3000", nil)

	// if err != nil {
	// 	log.Println("not working")
	// }

}
