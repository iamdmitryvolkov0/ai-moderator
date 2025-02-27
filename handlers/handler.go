package handlers

import (
	"encoding/json"
	"fmt"
	"moderator/moderation"
	"net/http"
)

func ModerateComment(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	text, exists := requestBody["text"]
	if !exists || text == "" {
		http.Error(w, "Text field is required", http.StatusBadRequest)
		return
	}

	result, err := moderation.ModerateComment(text)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
