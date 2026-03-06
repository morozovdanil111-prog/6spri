package main

import (
    "log"
    "net/http"
    "6spri/internal/server"  
)


func main() {
	// Создаем логгер
	logger := log.New()

	// Создаем сервер
	srv := server.CreateServer(logger)

	// Запускаем сервер
	log.Fatal(srv.ListenAndServe())
}
