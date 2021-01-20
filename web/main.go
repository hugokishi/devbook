package main

import (
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func init() {
	utils.LoadTemplates()
}

func main() {

	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
