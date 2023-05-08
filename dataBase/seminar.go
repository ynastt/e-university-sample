package dataBase

import (
	"database/sql"
	"encoding/json"
)

type Seminar struct {
    Id []uint8
	ModuleID []uint8
	Theme string
	Text json.RawMessage
	Db *sql.DB
}


func (s Seminar) Get_id() ([]uint8, []uint8) { return s.Id, s.ModuleID}
func (s Seminar) Get_text() string { 
	j, err := json.Marshal(s.Text)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (s *Seminar) Get_theme() string {return s.Theme}


func (s *Seminar) Set_text(text []byte) {
    s.Text = json.RawMessage(text)
	_, err := s.Db.Exec("update Seminar SET Text = $1 where SeminarID = $2", s.Text, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Seminar) Set_theme(t string) {
    s.Theme = t
	_, err := s.Db.Exec("update Seminar SET Theme = $1 where SeminarID = $2", s.Theme, s.Id)
    if err != nil {
        panic(err)
    }
}
