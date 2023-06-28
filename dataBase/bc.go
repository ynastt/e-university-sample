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
	Questions string
	MaxScore int
	MinScore int
	Db *sql.DB
}


func (b Bc) Get_id() ([]uint8, []uint8) { return b.Id, b.ModuleID }
func (b Bc) Get_theme() string { return b.Theme }
func (b Bc) Get_questions() string { 
	return b.Questions
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

func (b *Bc) Set_questions(text string) {
    b.Questions = text
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


// рубежный контроль студента
type BCInstance struct {
    StudentId []uint8
	BCID []uint8
    Date string 
	NumOfInstance int
	Score int
    Variant int
    Remarks json.RawMessage
	Db *sql.DB
}

func (i BCInstance) Get_id() ([]uint8, []uint8) { return i.StudentId, i.BCID }
func (i BCInstance) Get_date() string { return i.Date }
func (i BCInstance) Get_num_of_instance() int { return i.NumOfInstance }
func (i BCInstance) Get_score() int { return i.Score }
func (i BCInstance) Get_variant() int { return i.Variant }
func (i BCInstance) Get_remarks() string { 
	j, err := json.Marshal(i.Remarks)
	if err != nil {
		panic(err)
	}
	return string(j) 
}

func (i *BCInstance) Set_date(date string) {
    i.Date = date
	_, err := i.Db.Exec("update BCInstance SET DateOfPassing = $1 where student_id = $2 and bc_id = $3", i.Date, i.StudentId, i.BCID)
    if err != nil {
        panic(err)
    }
}
func (i *BCInstance) Set_num_of_instance(num int)  {
    i.NumOfInstance = num
	_, err := i.Db.Exec("update BCInstance SET NumOfInstance = $1 where student_id = $2 and bc_id = $3", i.NumOfInstance, i.StudentId, i.BCID)
    if err != nil {
        panic(err)
    }
}
func (i *BCInstance) Set_score(score int)  {
    i.Score = score
	_, err := i.Db.Exec("update BCInstance SET RecievedScore = $1 where student_id = $2 and bc_id = $3", i.Score, i.StudentId, i.BCID)
    if err != nil {
        panic(err)
    }
}
func (i *BCInstance) Set_variant(v int)  {
    i.Variant = v
	_, err := i.Db.Exec("update LBCInstance SET Variant = $1 where student_id = $2 and bc_id = $3", i.Variant, i.StudentId, i.BCID)
    if err != nil {
        panic(err)
    }
}

func (i *BCInstance) Set_remarks(text []byte) {
    i.Remarks = json.RawMessage(text)
	_, err := i.Db.Exec("update BCInstance SET Remarks = $1 where student_id = $2 and bc_id = $3", i.Remarks, i.StudentId, i.BCID)
    if err != nil {
        panic(err)
    }
}