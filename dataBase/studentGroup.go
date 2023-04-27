package dataBase 

import (
    "database/sql"
)

type StudentGroup struct{
    Id []uint8
    Name string
    YearOfAdm int
    Course int
    Amount int
    Db *sql.DB
}

func (sg StudentGroup) Get_id() []uint8 { return sg.Id }
func (sg StudentGroup) Get_name() string { return sg.Name }
func (sg StudentGroup) Get_yearOfAdm() int { return sg.YearOfAdm }
func (sg StudentGroup) Get_course() int{ return sg.Course }
func (sg StudentGroup) Get_amount() int{ return sg.Amount }

func (sg *StudentGroup) Set_name(name1 string) {
    sg.Name = name1
    _, err := sg.Db.Exec("update StudentGroup SET GroupName = $1 where GroupID = $2", sg.Name, sg.Id)
    if err != nil {
        panic(err)
    }
}

func (sg *StudentGroup) Set_yearOfAdm(yoa int) {
    sg.YearOfAdm = yoa
    _, err := sg.Db.Exec("update StudentGroup SET YearOfAdmission = $1 where GroupID = $2", sg.YearOfAdm, sg.Id)
    if err != nil {
        panic(err)
    }
}

func (sg *StudentGroup) Set_course(c int) {
    sg.Course = c
    _, err := sg.Db.Exec("update StudentGroup SET Course = $1 where GroupID = $2", sg.Course, sg.Id)
    if err != nil {
        panic(err)
    }
}

func (sg *StudentGroup) Set_amount(a int) {
    sg.Amount = a
    _, err := sg.Db.Exec("update StudentGroup SET AmountOfStudents = $1 where GroupID = $2", sg.Amount, sg.Id)
    if err != nil {
        panic(err)
    }
}
