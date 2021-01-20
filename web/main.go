package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {
	utils.LoadTemplates()

	r := router.GenerateRouter()
	fmt.Println("WebApp Rodando")
	log.Fatal(http.ListenAndServe(":3000", r))
}
