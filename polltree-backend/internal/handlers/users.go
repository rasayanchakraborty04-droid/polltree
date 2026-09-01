package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rasayanchakraborty04-droid/polltree/polltree-backend/internal/models"
)

func GetUsers(c *gin.Context) {

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

	// userss, err := userServices.GetAllUsers()

	// if err != nil {
	// 	c.JSON(500, g.H{
	// 		"error": "failed to fetch error",
	// 	})
	// }

	c.JSON(200, users)
}