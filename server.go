package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Data struct {
	path   string
	scheme string
}

type PersonData struct {
	Username string
	Password string
}

type SuccessResponse struct {
	Message   string
	CreatedAt time.Time
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

		p := PersonData{}

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			log.Fatal(err)
		}

		s := SuccessResponse{
			Message:   "Successed login with username: " + p.Username,
			CreatedAt: time.Now().Local(),
		}

		success, err := json.Marshal(&s)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("body: %+v", p)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(success)
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
