// if youre running this example
// and on cmd use this (for /L %i in (1,1,10) do curl http://localhost:3000/)

package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// always set (tokens/sec, burst) same for smoothness

var limiter = rate.NewLimiter(2, 2) // 3 requests/sec

func rateLimitedHandler(w http.ResponseWriter, r *http.Request) {
	if limiter.Allow() {
		fmt.Fprintf(w, "Request allowed at %s\n", time.Now().Format("15:04:05"))
		fmt.Println("Allowed at", time.Now().Format("15:04:05"))
	} else {
		http.Error(w, "❌ Too Many Requests", http.StatusTooManyRequests)
		fmt.Println("❌ Not allowed at", time.Now().Format("15:04:05"))
	}
}

func main() {
	http.HandleFunc("/", rateLimitedHandler)

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}
