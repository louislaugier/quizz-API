package score

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/louislaugier/quizz-API/database"
	response "github.com/louislaugier/quizz-API/src"
)

type score struct {
	Username *string `json:"username"`
	Score    *int    `json:"score"`
}

// POST new score
func POST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	score := &score{}
	json.NewDecoder(r.Body).Decode(&score)
	txn, _ := database.DB.Begin()
	txn.Exec("INSERT INTO scores (username, score) VALUES ($1, $2);", &score.Username, &score.Score)
	txn.Commit()
}

// GET scores with/without limit or get score for a username
func GET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	p, hasParam := mux.Vars(r)["param"]
	param := ""
	if hasParam {
		_, err := strconv.Atoi(p)
		if err != nil {
			param = " where username = '" + p + "'"
		} else {
			param = " limit " + p
		}
	}
	rows, err := database.DB.Query("select username, score from scores" + param + ";")
	defer rows.Close()
	res := response.Response{
		Error: err,
	}
	for rows.Next() {
		s := score{}
		rows.Scan(&s.Username, &s.Score)
		res.Data = append(res.Data, &s)
	}
	json.NewEncoder(w).Encode(res)
}
