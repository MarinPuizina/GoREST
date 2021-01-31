package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	readRequest := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Reading request")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data=%s\n", d)
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

	http.HandleFunc("/", readRequest)
	http.HandleFunc("/endpoint", h2)
	http.HandleFunc("/user", userHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
