package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LLMRequest struct {
	Message string `json:"message"`
}

type LLMResponse struct {
	Reply  string `json:"reply"`
	Action string `json:"action"`
}

func handleLLM(w http.ResponseWriter, r *http.Request) {
	var req LLMRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	// TODO: Call real LLM here
	reply := "NPC heard: " + req.Message
	action := "build_structure"

	resp := LLMResponse{Reply: reply, Action: action}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/llm", handleLLM)
	log.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
