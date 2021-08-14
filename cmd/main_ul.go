package main

import (
	"fileserver/cmd/client"
	"log"
)

func main() {
	url := "http://localhost:8080/upload"
	err := client.Upload(url, "1628312808803.gif")
	if err != nil {
		log.Fatal(err.Error())
	}
}
