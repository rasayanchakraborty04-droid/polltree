package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rasayanchakraborty04-droid/polltree/polltree-backend/internal/handlers"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(w, "Hello from go!")
	// })

	// http.HandleFunc("/api/health", handlers.Health)

	// http.HandleFunc("/api/users", handlers.GetUsers)

	// fmt.Println("Server started on http://localhost:8080")

	// err := http.ListenAndServe(":8080", nil)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello from Go + Gin",
		})
	})

	router.GET("/api/hello", handlers.Health)

	router.GET("/api/users", handlers.GetUsers)

	api := router.Group("/api")
	{
		users := api.Group("/users")

		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetUsers)
	}

	router.GET("/api/poll/:id", handlers.GetPolls)
	router.POST("/api/poll", handlers.CreatepollRequest)

	fmt.Println("Server has started running on http://ocalhost:8080")

	router.Run(":8080")
}