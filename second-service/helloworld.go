package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloServer)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(rw http.ResponseWriter, req *http.Request) {
        // log.Println("got request")

	fmt.Fprint(rw, "Hello World!!")
}

