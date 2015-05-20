package main

import (
	"log"
	"net/http"
)


func main() {

	router := NewRouter()

	log.Println("Server is running on 0.0.0.0:9090")

	log.Fatal(http.ListenAndServe(":9090", router))

}
