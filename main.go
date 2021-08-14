package main

import (
	"fileserver/controller"
)

const filefolder = "files/"

func main() {
	router := controller.SetupRouter(filefolder)
	controller.RunGin(router)

}
