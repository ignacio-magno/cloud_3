package server
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"


















}	log.Fatal(http.ListenAndServe(addr, nil))	log.Printf("Servidor principal corriendo en http://%s\n", addr)	addr := "0.0.0.0:" + port	})		fmt.Fprint(w, "Hello, World!")		w.WriteHeader(http.StatusOK)		w.Header().Set("Content-Type", "text/plain")	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {	}		port = "9000"	if port == "" {	port := os.Getenv("PORT")func main() {)