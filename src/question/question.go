package question

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louislaugier/quizz-API/database"
	response "github.com/louislaugier/quizz-API/src"
)

type question struct {
	ID       *int       `json:"questionID"`
	Question *string    `json:"question"`
	Answers  [4]*answer `json:"answers"`
}

type answer struct {
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"isCorrect"`
}

// GET questions with/without limit or get a random question
func GET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	p, hasParam := mux.Vars(r)["param"]
	param := ""
	if hasParam {
		param = " limit " + p
	}
	if p == "random" {
		param = " order by random() limit 1"
	}
	query := "select id, question, first_answer, second_answer, third_answer, fourth_answer, answer from questions" + param + ";"
	rows, err := database.DB.Query(query)
	defer rows.Close()
	res := response.Response{
		Error: err,
	}
	for rows.Next() {
		var a int
		q := question{
			Answers: [4]*answer{{}, {}, {}, {}},
		}
		rows.Scan(&q.ID, &q.Question, &q.Answers[0].Answer, &q.Answers[1].Answer, &q.Answers[2].Answer, &q.Answers[3].Answer, &a)
		q.Answers[a-1].IsCorrect = true
		res.Data = append(res.Data, &q)
	}
	json.NewEncoder(w).Encode(res)
}
