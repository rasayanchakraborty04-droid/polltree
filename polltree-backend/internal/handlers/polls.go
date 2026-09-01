package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rasayanchakraborty04-droid/polltree/polltree-backend/internal/models"
)

func GetPolls(c *gin.Context) {

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"poll_id": id,
	})

	// poll := models.Poll{
	// 	ID: 1,
	// 	Question: "What is your favorite programming language?",
	// 	Options: []models.Option{
	// 		{
	// 			ID: 1,
	// 			Text: "Go",
	// 		},
	// 		{
	// 			ID: 2,
	// 			Text: "Python",
	// 		},
	// 	},
	// }

	// c.JSON(http.StatusOK, poll)
}

func CreatepollRequest(c *gin.Context) {

	var request models.CreatePollRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid request body",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Poll create successfully",
		"question": request.Question,
		"options": request.Options,
	})

}