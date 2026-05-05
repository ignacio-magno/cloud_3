package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}

func deploy(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Deploy iniciado: git pull + pm2 restart")

	// Git pull
	cmd := exec.Command("git", "-C", "/home/ubuntu/cloud_3", "pull")
	if err := cmd.Run(); err != nil {
		log.Printf("Git pull error: %v\n", err)
		http.Error(w, "Git pull failed", http.StatusInternalServerError)
		return
	}

	// PM2 restart
	cmd = exec.Command("pm2", "restart", "mi-servidor")
	if err := cmd.Run(); err != nil {
		log.Printf("PM2 restart error: %v\n", err)
		http.Error(w, "PM2 restart failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pull + PM2 restart OK")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/deploy", deploy)

	addr := "0.0.0.0:" + port
	log.Printf("Servidor corriendo en http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
