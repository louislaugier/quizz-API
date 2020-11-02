package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/louislaugier/quizz-API/src/question"
	"github.com/louislaugier/quizz-API/src/score"
)

// Start router
func Start() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	r.GET("/api/scores/:limit", score.TopGET())
	r.GET("/api/score/:username", score.GET())
	r.GET("/api/questions/:limit", question.GET())
	r.GET("/api/question/random", question.RandomGET())
	return r
}
