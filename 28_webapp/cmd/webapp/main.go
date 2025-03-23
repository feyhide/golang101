package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/feyhide/golang101/28_webapp/internal/config"
	"github.com/feyhide/golang101/28_webapp/internal/http/handlers/student"
	"github.com/feyhide/golang101/28_webapp/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/v1/students", student.New(storage))
	router.HandleFunc("GET /api/v1/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/v1/students", student.GetList(storage))
	router.HandleFunc("PATCH /api/v1/students", student.Update(storage))
	router.HandleFunc("DELETE /api/v1/students/{id}", student.Delete(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address", cfg.HTTPServer.Addr))

	//----WITH GRACEFULLY SHUTDOWN----
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Warning: Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down The server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown gracefully")

	//----WITHOUT GRACEFULLY SHUTDOWN----
	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("Warning: Failed to start server")
	// }
}
