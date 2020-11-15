package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louislaugier/quizz-API/src/question"
	"github.com/louislaugier/quizz-API/src/score"
	"github.com/rs/cors"
)

// Init router
func Init() http.Handler {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	r.HandleFunc("/api/scores", score.POST).Methods("POST")
	r.HandleFunc("/api/scores", score.GET).Methods("GET")

	// param can be :limit (number) or :username (string)
	r.HandleFunc("/api/scores/{param}", score.GET).Methods("GET")

	r.HandleFunc("/api/questions", question.GET).Methods("GET")

	// param can be :limit (number) or "random"
	r.HandleFunc("/api/questions/{param}", question.GET).Methods("GET")

	h := c.Handler(r)
	return h
}
