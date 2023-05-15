package dataBase

import (
	"database/sql"
	"encoding/json"
)

// рубежный контроль в модуле
type Bc struct {
    Id []uint8
	ModuleID []uint8
	Theme string
	Questions json.RawMessage
	MaxScore int
	MinScore int
	Db *sql.DB
}


func (b Bc) Get_id() ([]uint8, []uint8) { return b.Id, b.ModuleID }
func (b Bc) Get_theme() string { return b.Theme }
func (b Bc) Get_questions() string { 
	j, err := json.Marshal(b.Questions)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (b Bc) Get_max_score() int { return b.MaxScore }
func (b Bc) Get_min_score() int { return b.MinScore }

func(b *Bc) Set_theme(text string) {
	b.Theme = text
	_, err := b.Db.Exec("update BC SET Theme = $1 where BCID = $2", b.Theme, b.Id)
    if err != nil {
        panic(err)
    }
}

func (b *Bc) Set_questions(text []byte) {
    b.Questions = json.RawMessage(text)
	_, err := b.Db.Exec("update BC SET Questions = $1 where BCID = $2", b.Questions, b.Id)
    if err != nil {
        panic(err)
    }
}

func (b *Bc) Set_max_score(score int) {
    b.MaxScore = score
	_, err := b.Db.Exec("update BC SET MaxScore = $1 where BCID = $2", b.MaxScore, b.Id)
    if err != nil {
        panic(err)
    }
}

func (b *Bc) Set_min_score(score int) {
    b.MinScore = score
	_, err := b.Db.Exec("update BC SET MinScore = $1 where BCID = $2", b.MinScore, b.Id)
    if err != nil {
        panic(err)
    }
}