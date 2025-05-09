package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AliNajafi2104/vvm_server/middleware"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	envUsername := "hej"
	envPassword := "123"

	if credentials.Username != envUsername || credentials.Password != envPassword {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, expiration, err := middleware.GenerateToken(credentials.Username)

	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":      token,
		"expiration": fmt.Sprintf("%d", expiration),
	})

}
