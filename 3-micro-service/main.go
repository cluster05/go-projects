package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	PORT = ":3000"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")

		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Hello %s", d)
	})

	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Println("Error :", err)
	}

}
