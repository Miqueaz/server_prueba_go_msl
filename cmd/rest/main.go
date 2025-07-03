package main

import (
	"fmt"
	config "main/config"
	"main/source/helpers/router"
	modules "main/source/modules/core"
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
