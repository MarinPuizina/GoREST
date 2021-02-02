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

	// Specifing the logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Reference to the handler
	hh := handlers.NewHello(l)

	// Create a new serve mux
	sm := http.NewServeMux()
	// Register a handler on a serve mux
	// And link it to our hh handler
	sm.Handle("/echo", hh)

	// Converting function to a handler type
	// Registering it to the default serve mux
	http.HandleFunc("/", readRequest)
	http.HandleFunc("/user", userHandler)

	// if we don't specify handler it will use default serve mux
	//log.Fatal(http.ListenAndServe(":8080", nil))

	// Specifying our own serve mux
	http.ListenAndServe(":8080", sm)
}
