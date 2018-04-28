package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/routes"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := httprouter.New()

	apiv1 := routes.Route()
	apiv1.RegisterAPI(router)

	log.Println("running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
