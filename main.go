package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("New request....")
		fmt.Fprintf(w, "Welcome to my website GoLang!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server at port 8085\n")

	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Started....")
	}
}
