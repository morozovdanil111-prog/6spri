package main

import (
	"log"
	"net/http"
	"yourproject/server"
)

func main() {
	// Создаем логгер
	logger := log.New()

	// Создаем сервер
	srv := server.CreateServer(logger)

	// Запускаем сервер
	log.Fatal(srv.ListenAndServe())
}