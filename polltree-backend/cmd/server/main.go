package main

import (
	"fmt"
	"net/http"

	"github.com/rasayanchakraborty04-droid/polltree/polltree-backend/internal/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello from go!")
	})

	http.HandleFunc("/api/health", handlers.Health)

	http.HandleFunc("/api/users", handlers.GetUsers)

	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}