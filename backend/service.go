package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/message", handleMessage)

	// Создаем новый CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Добавьте все нужные origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Оборачиваем наш mux в CORS middleware
	handler := c.Handler(mux)

	log.Println("Запуск бэкенд-сервера на порту :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Привет от бэкенда!"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
