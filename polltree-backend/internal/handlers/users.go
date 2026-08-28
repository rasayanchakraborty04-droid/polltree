package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rasayanchakraborty04-droid/polltree/polltree-backend/internal/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []models.User{
		{
			ID: 1,
			NAME: "User1",
			EMAIL: "user1@gmail.com",
		},
		{
			ID: 2,
			NAME: "User2",
			EMAIL: "user2@gmail.com",
		},
	}

	w.Header().Set("Cntent-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}