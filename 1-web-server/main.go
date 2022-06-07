package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello world")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "POST" {
		http.Error(w, "method not found", http.StatusNotFound)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error : %v", err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %v\nAddress : %v", name, address)

}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting on PORT 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Error while starting server ", err)
	}

}
