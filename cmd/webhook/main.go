package webhook
package main

import (
	"fmt"
	"log"
	"net/http"











































}	log.Fatal(http.ListenAndServe(addr, nil))	log.Printf("Webhook corriendo en http://%s/deploy\n", addr)	addr := "0.0.0.0:" + port	})		fmt.Fprint(w, "Pull + PM2 restart OK")		w.WriteHeader(http.StatusOK)		w.Header().Set("Content-Type", "text/plain")		}			return			http.Error(w, "PM2 restart failed", http.StatusInternalServerError)			log.Printf("PM2 restart error: %v\n", err)		if err := cmd.Run(); err != nil {		cmd = exec.Command("pm2", "restart", "mi-servidor")		// PM2 restart		}			return			http.Error(w, "Git pull failed", http.StatusInternalServerError)			log.Printf("Git pull error: %v\n", err)		if err := cmd.Run(); err != nil {		cmd := exec.Command("git", "-C", "/home/ubuntu/cloud_3", "pull")		// Git pull		log.Println("Deploy iniciado: git pull + pm2 restart")		}			return			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)		if r.Method != "POST" {	http.HandleFunc("/deploy", func(w http.ResponseWriter, r *http.Request) {	}		port = "3000"	if port == "" {	port := os.Getenv("PORT")func main() {)	"os/exec"	"os"