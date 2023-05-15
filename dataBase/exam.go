package dataBase

import (
	"database/sql"
	"encoding/json"
)

// экзамен по дисциплине
type Exam struct {
    Id []uint8
	SubjectID []uint8
	Questions json.RawMessage
	MaxScore int
	MinScore int
	Date string // like that '2005-01-01'
	Db *sql.DB
}


func (e Exam) Get_id() ([]uint8, []uint8) { return e.Id, e.SubjectID }
func (e Exam) Get_questions() string { 
	j, err := json.Marshal(e.Questions)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (e Exam) Get_max_score() int { return e.MaxScore }
func (e Exam) Get_min_score() int { return e.MinScore }
func (e Exam) Get_date() string { return e.Date }


func (e *Exam) Set_questions(text []byte) {
    e.Questions = json.RawMessage(text)
	_, err := e.Db.Exec("update Exam SET Questions = $1 where ExamID = $2", e.Questions, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_max_score(score int) {
    e.MaxScore = score
	_, err := e.Db.Exec("update Exam SET Maxscore = $1 where ExamID = $2", e.MaxScore, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_min_score(score int) {
    e.MinScore = score
	_, err := e.Db.Exec("update Exam SET MinScore = $1 where ExamID = $2", e.MinScore, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_date(date string) {
    e.Date = date
	_, err := e.Db.Exec("update Exam SET ExamDate = $1 where ExamID = $2", e.Date, e.Id)
    if err != nil {
        panic(err)
    }
}
