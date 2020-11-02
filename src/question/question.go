package question

import (
	"log"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID       int       `json:"questionID"`
	Question string    `json:"question"`
	Answers  []*answer `json:"answers"`
}

type answer struct {
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"isCorrect"`
}

// GET all questions
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println(c.Param("limit"))
	}
}

// RandomGET gets a random question
func RandomGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println("test")
	}
}
