package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("server start :8080")

	http.ListenAndServe(":8080", nil)
}
