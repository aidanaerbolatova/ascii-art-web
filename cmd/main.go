package main

import (
	"ascii/server"
	"log"
)

func main() {
	err := server.Run()
	log.Fatal(err)
}
