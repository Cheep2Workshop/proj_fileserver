package main

import "fileserver/cmd/client"

func main() {
	// file := "1628312808803.gif"
	file := "large.txt"
	client.Download("http://localhost:8080/downloadlarge/"+file, "./"+file)
}
