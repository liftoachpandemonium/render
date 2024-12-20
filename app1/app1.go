package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	ResponseFromContainer2 string `json:"response_from_container2"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Чтение параметра запроса
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "Hello from User"
	}

	// Отправка запроса контейнеру №2
	resp, err := http.PostForm("http://container2:8081/process", url.Values{"message": {message}})
	if err != nil {
		http.Error(w, "Failed to connect to Container 2", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа от контейнера №2
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from Container 2", http.StatusInternalServerError)
		return
	}

	// Формирование ответа для пользователя
	var response Response
	response.ResponseFromContainer2 = fmt.Sprintf("Container 2 Response: %s", string(body))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/request", handleRequest)
	fmt.Println("Container 1 listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
