package main

import (
	"log"
	"net/http"
)

func main() {

	const addr = "127.0.0.1:8080"
	log.Println("listen at ", addr)

	http.HandleFunc("/", fileHandler)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

}
