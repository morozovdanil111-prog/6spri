package main

import import (
	"log"
	"net/http"
	"yourproject/internal/server"
)


func main() {
	// Создаем логгер
	logger := log.New()

	// Создаем сервер
	srv := server.CreateServer(logger)

	// Запускаем сервер
	log.Fatal(srv.ListenAndServe())
}
