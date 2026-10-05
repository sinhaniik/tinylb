// Package server provides HTTP server handlers.
package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello Server")

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
