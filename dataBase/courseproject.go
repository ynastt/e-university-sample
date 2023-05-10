package dataBase 

import (
    "database/sql"
)

var subjects_counter = 1
var subjects = map[string]int {
    "базы данных": 1,
}

// функция нахождения ключа по значению
func mapkey(mapa map[string]int, value int) (key string, ok bool) {
	for k, v := range mapa {
	  if v == value { 
		key = k
		ok = true
		return
	  }
	}
	return
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
func (p CourseProject) Get_subject() string { 
	subj, ok := mapkey(subjects, p.Subject)
	if !ok {
  		panic("there is no such subject in database")
	}
	return subj
}
func (p CourseProject) Get_description() []byte { return p.Description }
func (p CourseProject) Get_hours() int { return p.Hours }
func (p CourseProject) Get_start_date() string { return p.StartDate }
func (p CourseProject) Get_deadline() string { return p.Deadline }

func (p *CourseProject) Set_subject(name string) {
	if _, ok := subjects[name]; !ok {
		subjects_counter += 1
		subjects[name] = subjects_counter
	}
	p.Subject = subjects[name]
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