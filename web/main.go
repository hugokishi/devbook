package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/router"
	"web/src/utils"
)

func main() {
	utils.LoadTemplates()
	config.Load()

	r := router.GenerateRouter()
	fmt.Printf("Web rodando na porta: %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
