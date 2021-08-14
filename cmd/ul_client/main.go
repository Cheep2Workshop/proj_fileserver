package main

import (
	"fileserver/cmd/client"
	"log"
)

func main() {
	url := "http://localhost:8080/upload"
	// file := "1628312808803.gif"
	file := "2021-05-01 02-44-36.mkv"
	err := client.Upload(url, file)
	if err != nil {
		log.Fatal(err.Error())
	}
}
