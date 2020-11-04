package router

import (
	"github.com/gorilla/mux"
	"github.com/louislaugier/quizz-API/src/question"
	"github.com/louislaugier/quizz-API/src/score"
)

// Init router
func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/scores", score.POST).Methods("POST")
	r.HandleFunc("/api/scores", score.GET).Methods("GET")

	// param can be :limit (number) or :username (string)
	r.HandleFunc("/api/scores/{param}", score.GET).Methods("GET")

	r.HandleFunc("/api/questions", question.GET).Methods("GET")

	// param can be :limit (number) or "random"
	r.HandleFunc("/api/questions/{param}", question.GET).Methods("GET")

	return r
}
