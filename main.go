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

	userHandler := func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["key"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url param key is missing")
			return
		}

		key := keys[0]

		io.WriteString(w, "Url param key is="+string(key))
		log.Println("Url param key is=" + string(key))
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	http.HandleFunc("/user", userHandler)

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
