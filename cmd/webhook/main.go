package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const repoPath = "/home/ubuntu/cloud_3"

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		log.Printf("%s %v output:\n%s", name, args, string(output))
	}
	return err
}

func deployHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Deploy started: pull -> build -> restart main service")

	if err := runCommand("git", "pull"); err != nil {
		log.Printf("git pull failed: %v", err)
		http.Error(w, "git pull failed", http.StatusInternalServerError)
		return
	}

	if err := runCommand("go", "build", "-o", "bin/server", "cmd/server/main.go"); err != nil {
		log.Printf("build failed: %v", err)
		http.Error(w, "build failed", http.StatusInternalServerError)
		return
	}

	if err := runCommand("systemctl", "restart", "servidor-main"); err != nil {
		log.Printf("restart failed: %v", err)
		http.Error(w, "restart failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deploy ok")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "webhook up")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/deploy", deployHandler)

	addr := "0.0.0.0:" + port
	log.Printf("Webhook listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
