package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func longProcess(ctx context.Context, resultChan chan<- string) {
	select {
	case <-time.After(5 * time.Second):
		resultChan <- "Process Completed!"
	case <-ctx.Done():
		resultChan <- "Request Cancelled!"
	}
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	resultChan := make(chan string, 1)

	go longProcess(ctx, resultChan)

	result := <-resultChan

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": result})
}

func runEx3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/process", processHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	fmt.Println("Server running on :3000")
	server.ListenAndServe()
}
