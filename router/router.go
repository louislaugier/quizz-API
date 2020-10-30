package router

import (
	"github.com/louislaugier/quizz-API/src/score"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgres driver
)

// Start the router
func Start() *gin.Engine {

	// gin.SetMode(gin.ReleaseMode)

	base := "/api/"

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST"},
		ExposeHeaders: []string{"Content-Type", "Date"},
		// AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		AllowAllOrigins: true,
	}))

	// Items
	r.GET(base+"/scores", score.GET())

	return r
}
