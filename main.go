package main

import (
	"bootcamp/routes"
	"fmt"
)

func main() {
	router := routes.Router()

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
