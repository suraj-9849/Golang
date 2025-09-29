package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/suraj-9849/Golang.git/internal/config"
	"github.com/suraj-9849/Golang.git/internal/http/handlers/student"
	"github.com/suraj-9849/Golang.git/internal/storage/sqlite"
)

func main() {
	//load config:
	cfg := config.MustLoad()
	//database Setup:
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage Initialized", slog.String("env", cfg.Env))
	// setup the Router:
	router := http.NewServeMux()
	router.HandleFunc("POST /api/student", student.New(storage))
	router.HandleFunc("GET /api/student/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/student/{id}", student.UpdateById(storage))
	router.HandleFunc("DELETE /api/student/{id}", student.DeleteTheStudentByID(storage))
	//setup the server:
	server := http.Server{
		Addr:    cfg.Address + ":" + cfg.PORT,
		Handler: router,
	}
	slog.Info("Server starting at \n", slog.String("config", cfg.Address), slog.String("port", cfg.PORT))
	fmt.Printf("Server starting at %s:%s\n", cfg.Address, cfg.PORT)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start the Server")
		}
	}()

	<-done
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to ShutDown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutDown successfully!")
}
