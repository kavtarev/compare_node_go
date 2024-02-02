package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "compare_node_go"
)



























// func doSome(res http.ResponseWriter, req *http.Request) {
// 	res.Write([]byte("do some"))
// }

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, e := sql.Open("postgres", psqlInfo)

// 	if e != nil {
// 		fmt.Println(111)
// 	}
// 	defer db.Close()

// 	errPing := db.Ping()

// 	if errPing != nil {
// 		fmt.Println(errPing)
// 	}

// 	res, err := db.Query("select count(1) from users;")

// 	if err != nil {
// 		fmt.Println(123)
// 	}
// 	defer res.Close()

// 	var n []int

// 	for res.Next() {
// 		err :=res.Scan(&n)
// 		if err != nil {
// 			fmt.Println(8888, err)
// 		}

// 		fmt.Println(5566, n)
// 	}

// 	// http.HandleFunc("/", doSome)
// 	// log.Println("up on 3000")
// 	// log.Fatal(http.ListenAndServe(":3000", nil))

// }
