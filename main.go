package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", myapp.NewHttpHandler()))
}
