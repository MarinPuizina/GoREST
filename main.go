package main

import (
	"fmt"
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

	echoRequest := func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Reading request")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Sorry, a bad request", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
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
	http.HandleFunc("/echo", echoRequest)
	http.HandleFunc("/user", userHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
