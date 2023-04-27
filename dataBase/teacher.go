package dataBase 

import (
    "database/sql"
)

type Teacher struct{
    Id []uint8
    Name string
	Surname string
	Patronymic string
	Email string
	User_id []uint8
	Db *sql.DB
}

func (t Teacher) Get_id() []uint8 { return t.Id }
func (t Teacher) Get_name() string { return t.Name }
func (t Teacher) Get_surname() string { return t.Surname }
func (t Teacher) Get_patronymic() string{ return t.Patronymic }
func (t Teacher) Get_email() string{ return t.Email }

func (t *Teacher) Set_name(name1 string) {
    t.Name = name1
    _, err := t.Db.Exec("update Teacher SET TeacherName = $1 where TeacherId = $2", t.Name, t.Id)
    if err != nil {
        panic(err)
    }
}

func (t *Teacher) Set_surname(name1 string) {
    t.Surname = name1
    _, err := t.Db.Exec("update Teacher SET Surname = $1 where TeacherId = $2", t.Surname, t.Id)
    if err != nil {
        panic(err)
    }
}

func (t *Teacher) Set_Patronymic(name1 string) {
    t.Patronymic = name1
    _, err := t.Db.Exec("update Teacher SET Patronymic = $1 where TeacherId = $2", t.Patronymic, t.Id)
    if err != nil {
        panic(err)
    }
}

func (t *Teacher) Set_Email(name1 string) {
    t.Email = name1
    _, err := t.Db.Exec("update Teacher SET Email = $1 where TeacherId = $2", t.Email, t.Id)
    if err != nil {
        panic(err)
    }
}