package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Env struct {
	dbHost, dbUser, dbPass, dbName string
}

func main() {
	router := mux.NewRouter()
	dbCon := Env{
		dbHost: "192.168.0.100",
		dbUser: "db_user",
		dbPass: "Proj123",
		dbName: "go_web_api",
	}
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		dbCon.connectDb()
		fmt.Fprintf(rw, "Welcome to Golang home page...")
	})

	router.HandleFunc("/books/{title}/page/{page}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(rw, "You've requested the book: %s on page %s\n", title, page)
	})

	fmt.Printf("Starting server at port 8085\n")

	if err := http.ListenAndServe(":8085", router); err != nil {
		log.Fatal(err)
	}
}

func (env Env) connectDb() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true", env.dbUser, env.dbPass, env.dbHost, env.dbName))

	if db == nil && err != nil {
		fmt.Printf(" Connection error ")
	}
}
