package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":3333", r))
}
