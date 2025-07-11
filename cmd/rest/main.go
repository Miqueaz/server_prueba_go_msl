package main

import (
	"fmt"
	config "main/config"
	modules "main/source/core"
	"main/source/helpers/router"
)

func main() {
	fmt.Print("\033[H\033[2J")
	config.Execute()
	modules.Execute()
	Execute()
}

func Execute() {
	r := router.Router()
	r.Execute(":8080")
}
