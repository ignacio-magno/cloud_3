package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello, World 6")
	})

	addr := "0.0.0.0:" + port
	log.Printf("Main server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
