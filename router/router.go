package router

import (
	"github.com/gorilla/mux"
	"github.com/louislaugier/quizz-API/src/question"
	"github.com/louislaugier/quizz-API/src/score"
)

// Init router
func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/scores", score.POST).Methods("POST")
	r.HandleFunc("/scores", score.GET).Methods("GET")

	// param can be :limit (number) or :username (string)
	r.HandleFunc("/scores/{param}", score.GET).Methods("GET")

	r.HandleFunc("/questions", question.GET).Methods("GET")

	// param can be :limit (number) or "random"
	r.HandleFunc("/questions/{param}", question.GET).Methods("GET")

	return r
}
