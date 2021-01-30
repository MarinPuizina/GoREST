package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from handler func 1")
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from handler func 2")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))

	/*s := &server{}
	http.HandleFunc("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))*/
}

/*
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r * http.Reqeust) {
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte('{"message": "Hello Lasma"}'))
}*/
