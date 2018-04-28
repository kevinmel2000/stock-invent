package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := httprouter.New()

	router.GET("/", index)

	log.Println("running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("API working"))
}
