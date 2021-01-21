package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/serv", serve)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Hello")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Sending request
func serve(w http.ResponseWriter, r *http.Request) {
	log.Printf("whatever")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "SUCCESS")
	name := r.FormValue("name")

	fmt.Fprintf(w, "Name = %s\n", name)
}
