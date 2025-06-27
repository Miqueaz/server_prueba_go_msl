package main

import (
	"fmt"
	"log"
	config "main/config"
	modules "main/modules/core/cmd"
	router "main/modules/core/router"
	"net/http"
)

func main() {
	fmt.Print("\033[H\033[2J")
	config.Execute()
	modules.Execute()
	Execute()
}

func Execute() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r.Router))
}
