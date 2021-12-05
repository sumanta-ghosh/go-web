package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
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
