package dataBase 

import (
    "database/sql"
)

var c = 1
var m = map[int]string {
    1:"базы данных",
}

type CourseProject struct {
    Id []uint8
    Subject int
    Description []byte
	Hours int
	StartDate string
	Deadline string
	Db *sql.DB
}

func (p CourseProject) Get_id() []uint8 { return p.Id }
func (p CourseProject) Get_subject() string { return m[p.Subject] }
func (p CourseProject) Get_description() []byte { return p.Description }
func (p CourseProject) Get_hours() int { return p.Hours }
func (p CourseProject) Get_start_date() string { return p.StartDate }
func (p CourseProject) Get_deadline() string { return p.Deadline }

func (p *CourseProject) Set_subject(name string) {
	c += 1
	m[c] = name
    p.Subject = c
    _, err := p.Db.Exec("update CourseProject SET Subject = $1 where ProjectID = $2", p.Subject, p.Id)
    if err != nil {
        panic(err)
    }
}

func (p *CourseProject) Set_description(text []byte) {
    p.Description = text
    _, err := p.Db.Exec("update CourseProject SET Description = $1 where ProjectID = $2", p.Description, p.Id)
    if err != nil {
        panic(err)
    }
}

func (p *CourseProject) Set_hours(name1 int) {
    p.Hours = name1
    _, err := p.Db.Exec("update CourseProject SET Hours = $1 where ProjectID = $2", p.Hours, p.Id)
    if err != nil {
        panic(err)
    }
}

func (p *CourseProject) Set_start_date(date string) {
    p.StartDate = date
	_, err := p.Db.Exec("update CourseProject SET StartDate = $1 where ProjectID = $2", p.StartDate, p.Id)
    if err != nil {
        panic(err)
    }
}

func (p *CourseProject) Set_deadline(date string) {
    p.Deadline = date
	_, err := p.Db.Exec("update CourseProject SET Deadline = $1 where ProjectID = $2", p.Deadline, p.Id)
    if err != nil {
        panic(err)
    }
}