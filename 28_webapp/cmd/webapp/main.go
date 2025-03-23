package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feyhide/golang101/28_webapp/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Welcome to feyhide go webserver"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("[+] Server Started on %s \n\n", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Warning: Failed to start server")
	}
}
