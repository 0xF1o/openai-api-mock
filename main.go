package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MessageOut struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Choice struct {
	Index   int        `json:"index"`
	Message MessageOut `json:"message"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func chatCompletionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	reqSize := len(bodyBytes)

	response := Response{
		Choices: []Choice{
			{
				Index: 0,
				Message: MessageOut{
					Content: fmt.Sprintf("I'm just a blind mock service. Have a nice day! (Request size: %d bytes)", reqSize),
					Role:    "assistant",
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/v1/chat/completions", chatCompletionsHandler)

	log.Println("Mock OpenAI server running on port 1323")
	log.Fatal(http.ListenAndServe(":1323", nil))
}
