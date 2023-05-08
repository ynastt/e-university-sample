package dataBase

import (
	"database/sql"
	"encoding/json"
)

type Lecture struct {
    Id []uint8
	ModuleID []uint8
	Theme string
	Text json.RawMessage
	Db *sql.DB
}


func (l Lecture) Get_id() ([]uint8, []uint8) { return l.Id, l.ModuleID}
func (l Lecture) Get_text() string { 
	j, err := json.Marshal(l.Text)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (l *Lecture) Get_theme() string {return l.Theme}


func (l *Lecture) Set_text(text []byte) {
    l.Text = json.RawMessage(text)
	_, err := l.Db.Exec("update Lecture SET Text = $1 where LectureID = $2", l.Text, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lecture) Set_theme(t string) {
    l.Theme = t
	_, err := l.Db.Exec("update Lecture SET Theme = $1 where LectureID = $2", l.Theme, l.Id)
    if err != nil {
        panic(err)
    }
}
