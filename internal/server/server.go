// Package server provides HTTP server handlers.
package server

import (
	"fmt"
	"net/http"
)

func Server() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello Server")

	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		// server failed to start
		return err
	}
	return nil
}
