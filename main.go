package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("servers started at :8080")
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	if err != nil {
		log.Fatal(err)
	}
}
