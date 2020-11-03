package main

import (
	"log"
	"net/http"

	"github.com/louislaugier/quizz-API/router"
)

func main() {
	r := router.Init()
	log.Fatal(http.ListenAndServe(":80", r))
}
