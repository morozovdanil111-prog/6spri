package server

import (
	"net/http"
	"log"
	"time"
	"yourproject/handlers"
)

// CreateServer создаёт и настраивает HTTP сервер
func CreateServer(logger *log.Logger) *http.Server {
	// Создаем HTTP роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем сервер
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}