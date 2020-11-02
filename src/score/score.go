package score

import (
	"log"

	"github.com/gin-gonic/gin"
)

type score struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}

// TopGET gets top scores
func TopGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println(c.Param("limit"))
	}
}

// GET score by username
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println(c.Param("username"))
	}
}
