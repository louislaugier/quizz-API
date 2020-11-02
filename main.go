package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louislaugier/quizz-API/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router.Start().Run()
}
