package main

import (
	"fmt"
	"tinylb/internal/server"
)

func main() {
	server.Server()
	fmt.Println("Server started")
}
