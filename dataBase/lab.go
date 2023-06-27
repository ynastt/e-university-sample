package dataBase

import (
	"database/sql"
)

// лабораторная работа в модуле
type Lab struct {
    Id []uint8
	ModuleID []uint8
    Number int
	Name string
	Text string
	MaxScore int
	MinScore int
	Date string     // таким образом '2005-01-01'
	Deadline string
	Db *sql.DB
}


func (l Lab) Get_id() ([]uint8, []uint8) { return l.Id, l.ModuleID }
func (l Lab) Get_number() int { return l.Number }
func (l Lab) Get_name() string { return l.Name }
func (l Lab) Get_text() string { return l.Text }
func (l Lab) Get_max_score() int { return l.MaxScore }
func (l Lab) Get_min_score() int { return l.MinScore }
func (l Lab) Get_date() string { return l.Date }
func (l Lab) Get_dealine() string { return l.Deadline }

func (l *Lab) Set_number(num int) {
    l.MaxScore = num
	_, err := l.Db.Exec("update Lab SET LabNumber = $1 where LabID = $2", l.Number, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_name( name1 string) { 
    l.Name = name1
	_, err := l.Db.Exec("update Lab SET Name = $1 where LabID = $2", l.Name, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lab) Set_text(text string) {
    l.Text = text
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

// лабораторная работa студента
type LabInstance struct {
    StudentId []uint8
	LabID []uint8
    Date string 
	NumOfInstance int
	Score int
    Variant sql.NullInt64
    Remarks string
    BonusScore sql.NullInt64
	Db *sql.DB
}

func (i LabInstance) Get_id() ([]uint8, []uint8) { return i.StudentId, i.LabID }
func (i LabInstance) Get_date() string { return i.Date }
func (i LabInstance) Get_num_of_instance() int { return i.NumOfInstance }
func (i LabInstance) Get_score() int { return i.Score }
func (i LabInstance) Get_variant() int { return int(i.Variant.Int64) }
func (i LabInstance) Get_remarks() string { return i.Remarks }
func (i LabInstance) Get_bonus_score() int { return int(i.BonusScore.Int64)}

func (i *LabInstance) Set_date(date string) {
    i.Date = date
	_, err := i.Db.Exec("update LabInstance SET DateOfPassing = $1 where student_id = $2 and lab_id = $3", i.Date, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}
func (i *LabInstance) Set_num_of_instance(num int)  {
    i.NumOfInstance = num
	_, err := i.Db.Exec("update LabInstance SET NumOfInstance = $1 where student_id = $2 and lab_id = $3", i.NumOfInstance, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}
func (i *LabInstance) Set_score(score int)  {
    i.Score = score
	_, err := i.Db.Exec("update LabInstance SET RecievedScore = $1 where student_id = $2 and lab_id = $3", i.Score, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}
func (i *LabInstance) Set_variant(v int)  {
    i.Variant = sql.NullInt64{Int64: int64(v), Valid: true}
	_, err := i.Db.Exec("update LabInstance SET Variant = $1 where student_id = $2 and lab_id = $3", i.Variant, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}

func (i *LabInstance) Set_remarks(text string) {
    i.Remarks = text
	_, err := i.Db.Exec("update LabInstance SET Remarks = $1 where student_id = $2 and lab_id = $3", i.Remarks, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}

func (i *LabInstance) Set_bonus_score(score int)  {
    i.BonusScore = sql.NullInt64{Int64: int64(score), Valid: true}
	_, err := i.Db.Exec("update LabInstance SET BonusScore = $1 where student_id = $2 and lab_id = $3", i.BonusScore, i.StudentId, i.LabID)
    if err != nil {
        panic(err)
    }
}