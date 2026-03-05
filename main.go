package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	port := ":5000"
	fmt.Println("Server us running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
