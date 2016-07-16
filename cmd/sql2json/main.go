package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gdey/sql2json"
)

var conn string
var sqlStr string
var dbName string

func init() {
	flag.StringVar(&sqlStr, "sql", "", "The select sql to run.")
	flag.StringVar(&conn, "conn", "", "Connection string for mysql db.")
	dbs := strings.Join(sql.Drivers(), ",")
	flag.StringVar(&dbName, "db", "mysql", fmt.Sprintf("The Database to use, available databases are: [%v]", dbs))
}

func main() {
	flag.Parse()
	if conn == "" {
		log.Println("Connection string is required.")
		os.Exit(1)
	}
	if sqlStr == "" {
		log.Println("SQL is required.")
		os.Exit(1)
	}
	db, err := sql.Open(dbName, conn)
	if err != nil {
		log.Printf("Error connecting(%v) to db: %v", conn, err)
		os.Exit(2)
	}
	json, err := sql2json.Query(db, sqlStr)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(3)
	}
	fmt.Println(string(json))
}
