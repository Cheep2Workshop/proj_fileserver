package main

import (
	"fileserver/controller"
)

func main() {
	router := controller.SetupRouter()
	controller.RunGin(router)

}
