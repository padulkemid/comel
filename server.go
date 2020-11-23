package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	path   string
	scheme string
}

func comelAvailable(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Printf("path: %s \n", r.URL.Path)
	fmt.Printf("scheme: %s \n", r.URL.Scheme)

	fmt.Fprintf(w, "What's up? %+v", Data{r.URL.Path, r.URL.Scheme})
}

func loginAja(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method: %s \n", r.Method)

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("username: %s", r.Form["username"])
		fmt.Printf("password: %s", r.Form["password"])
	}
}

func checkComel(w http.ResponseWriter, r *http.Request) {
	status := []string{"available", "not available", "unknown"}

	if r.Method == "POST" {
		for _, v := range status {
			if v == r.Form.Get("status") {
				fmt.Printf("status is: %s", r.Form.Get("status"))
			}
		}
	}

}

func main() {
	// set router
	http.HandleFunc("/", comelAvailable)
	http.HandleFunc("/login", loginAja)
	http.HandleFunc("/status", checkComel)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
