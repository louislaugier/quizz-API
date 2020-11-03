package main

import (
	"log"
	"net/http"
	"os"

	"github.com/louislaugier/quizz-API/router"
)

func main() {
	r := router.Init()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
