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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s := &score{}
	json.NewDecoder(r.Body).Decode(&s)
	rows, _ := database.DB.Query("select count(*) from scores where username='" + *s.Username + "';")
	defer rows.Close()
	var existingUser *int
	for rows.Next() {
		rows.Scan(&existingUser)
	}
	txn, _ := database.DB.Begin()
	if *existingUser == 1 {
		txn.Exec("update scores set score = $1 where username = $2;", &s.Score, &s.Username)
	} else {
		txn.Exec("insert into scores (username, score) values ($1, $2);", &s.Username, &s.Score)
	}
	txn.Commit()
	json.NewEncoder(w).Encode(s)
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
			param = " where username = '" + p + "' order by score desc;"
		} else {
			param = " order by score desc limit " + p + ";"
		}
	}
	rows, err := database.DB.Query("select username, score from scores" + param)
	defer rows.Close()
	res := response.Response{
		Error: err,
	}
	for rows.Next() {
		s := score{}
		rows.Scan(&s.Username, &s.Score)
		res.Scores = append(res.Scores, &s)
	}
	json.NewEncoder(w).Encode(res)
}
