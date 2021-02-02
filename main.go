package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"marin.com/rest/handlers"
)

func main() {
	readRequest := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Reading request")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data=%s\n", d)
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

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/echo", hh)

	http.HandleFunc("/", readRequest)
	http.HandleFunc("/user", userHandler)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	http.ListenAndServe(":8080", sm)
}
