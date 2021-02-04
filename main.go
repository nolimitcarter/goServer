package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/serve", serve)
	http.HandleFunc("/form", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//	fmt.Printf("Hello")
	//	if err := http.ListenAndServe(":8080", nil); err != nil {
	//		log.Fatal(err)
	//}
}

// Sending request
func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formHandler(w, r)
		return
	}
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, wd+r.URL.Path[1:])
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jsn, _ := json.MarshalIndent(r.Form, "", " ")
	fmt.Fprintf(w, string(jsn))
}

//func formHandler(w http.ResponseWriter, r *http.Request) {
//	if err := r.ParseForm(); err != nil {
//		fmt.Fprintf(w, "error: %v", err)
//		return
//	}
//	fmt.Fprintf(w, "SUCCESS")
//	name := r.FormValue("name")
//
//	fmt.Fprintf(w, "Name = %s\n", name)
//}
