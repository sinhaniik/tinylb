// Package server provides HTTP server handlers.
package server

import (
	"fmt"
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server")
	})
	return mux
}

func Server() error {
	return http.ListenAndServe(":8080", Handler())
}
