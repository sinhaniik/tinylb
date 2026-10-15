package main

import (
	"log"
	"tinylb/internal/server"
)

func main() {
	err := server.Server()

	if err != nil {
		log.Fatal(err)
	}

}
