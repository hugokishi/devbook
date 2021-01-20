package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"
)

func main() {
	config.Load()
	cookies.Load()
	utils.LoadTemplates()

	r := router.GenerateRouter()
	fmt.Printf("Web rodando na porta: %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
