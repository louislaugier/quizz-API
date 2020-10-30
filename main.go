package main

import (
	"log"

	"github.com/louislaugier/quizz-API/database"
	"github.com/louislaugier/quizz-API/router"
)

func main() {
	router.Start().Run()
	log.Println(database.DB)
}
