package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gdey/sql2json"
)

import _ "github.com/go-sql-driver/mysql"

var conn = flag.String("conn", "", "Connection string for mysql db.")
var sqlStr = flag.String("sql", "", "The select sql to run.")

func main() {
	flag.Parse()
	if *conn == "" {
		log.Println("Connection string is required.")
		os.Exit(1)
	}
	if *sqlStr == "" {
		log.Println("SQL is required.")
		os.Exit(1)
	}
	db, err := sql.Open("mysql", *conn)
	if err != nil {
		log.Printf("Error connecting(%v) to db: %v", conn, err)
		os.Exit(2)
	}
	json, err := sql2json.Query(db, *sqlStr)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(3)
	}
	fmt.Println(json)
}
