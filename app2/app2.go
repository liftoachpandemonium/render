package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Response string `json:"response"`
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	// Чтение данных из запроса
	var data map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Формирование ответа
	message := data["message"]
	response := Response{Response: fmt.Sprintf("Processed by Container 2: %s", message)}

	// Ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/process", processRequest)
	fmt.Println("Container 2 listening on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
