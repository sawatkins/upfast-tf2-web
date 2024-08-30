package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
)

const awsEndpoint = "https://bwdfgz2pbedm7ficoxqxbhfazi0ynfoh.lambda-url.us-west-1.on.aws"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/server-ips", handleServerIPs)
	http.HandleFunc("/api/server-info", handleServerInfo)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleServerIPs(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(awsEndpoint)
	if err != nil {
		http.Error(w, "Failed to fetch server IPs", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func handleServerInfo(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing IP parameter", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://%s:8000/server-info", ip)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch server info: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}