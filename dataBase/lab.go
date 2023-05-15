package dataBase

import (
	"database/sql"
	"encoding/json"
)

// лабораторная работа в модуле
type Lab struct {
    Id []uint8
	ModuleID []uint8
	Name string
	Text json.RawMessage
	MaxScore int
	MinScore int
	Date string     // таким образом '2005-01-01'
	Deadline string
	Db *sql.DB
}


func (l Lab) Get_id() ([]uint8, []uint8) { return l.Id, l.ModuleID }
func (l Lab) Get_text() string { 
	j, err := json.Marshal(l.Text)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (l Lab) Get_max_score() int { return l.MaxScore }
func (l Lab) Get_min_score() int { return l.MinScore }
func (l Lab) Get_date() string { return l.Date }
func (l Lab) Get_dealine() string { return l.Deadline }

func (l *Lab) Set_text(text []byte) {
    l.Text = json.RawMessage(text)
	_, err := l.Db.Exec("update Lab SET Text = $1 where LabID = $2", l.Text, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_max_score(score int) {
    l.MaxScore = score
	_, err := l.Db.Exec("update Lab SET Maxscore = $1 where LabID = $2", l.MaxScore, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_min_score(score int) {
    l.MinScore = score
	_, err := l.Db.Exec("update Lab SET MinScore = $1 where LabID = $2", l.MinScore, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_date(date string) {
    l.Date = date
	_, err := l.Db.Exec("update Lab SET LabDate = $1 where LabID = $2", l.Date, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_deadline(date string) {
    l.Deadline = date
	_, err := l.Db.Exec("update Lab SET Deadline = $1 where LabID = $2", l.Deadline, l.Id)
    if err != nil {
        panic(err)
    }
}