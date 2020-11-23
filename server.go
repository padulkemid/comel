package main

import (
	"fmt"
	"log"
	"net/http"
)

func comelAvailable(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Printf("path: %s \n", r.URL.Path)
	fmt.Printf("scheme: %s \n", r.URL.Scheme)

	fmt.Fprintf(w, "What's up?")
}

func main() {
	// set router
	http.HandleFunc("/", comelAvailable)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
