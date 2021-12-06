package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Env struct {
	dbHost, dbPort, dbUser, dbPass, dbName string
}

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost:3306"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("DB user name, password & db name env variable are not set properly !!!")
	}

	router := mux.NewRouter()
	dbCon := Env{
		dbHost: dbHost,
		dbPort: dbPort,
		dbUser: dbUser,
		dbPass: dbPass,
		dbName: dbName,
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
	fmt.Print(env)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", env.dbUser, env.dbPass, env.dbHost, env.dbPort, env.dbName))

	err = db.Ping()

	if err != nil {
		fmt.Println(" Connection error ")
	}
}
