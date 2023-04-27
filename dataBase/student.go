package dataBase 

import (
    "database/sql"
)

type Student struct{
    Id []uint8
    Name string
	Surname string
	Patronymic string
	Email string
    Phone string
    Courses int
    Number int
    Year int
    Userid []uint8
    Groupid []uint8
    Db *sql.DB
}


func (s Student) Get_id() []uint8 { return s.Id }
func (s Student) Get_name() string { return s.Name }
func (s Student) Get_surname() string { return s.Surname }
func (s Student) Get_patronymic() string{ return s.Patronymic }
func (s Student) Get_email() string { return s.Email }
func (s Student) Get_phone() string { return s.Phone }
func (s Student) Get_courses() int { return s.Courses }
func (s Student) Get_number() int{ return s.Number }
func (s Student) Get_year() int{ return s.Year }

func (s *Student) Set_name(name1 string) {
    s.Name = name1
    _, err := s.Db.Exec("update Student SET StudentName = $1 where StudentID = $2", s.Name, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_surname(name1 string) {
    s.Surname = name1
    _, err := s.Db.Exec("update Student SET Surname = $1 where StudentID = $2", s.Surname, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_patronimic(p string) {
    s.Patronymic = p
    _, err := s.Db.Exec("update Student SET Patronymic = $1 where StudentID = $2", s.Patronymic, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_email(e string) {
    s.Email = e
    _, err := s.Db.Exec("update Student SET Email  = $1 where StudentID = $2)", s.Email, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_phone(p string) {
    s.Phone = p
    _, err := s.Db.Exec("update Student set Phone = $1 where StudentID = $2", s.Phone, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_courses(c int) {
    s.Courses = c
    _, err := s.Db.Exec("update Student set PassedCourses = $1 where StudentID = $2", s.Courses, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_number(c int) {
    s.Number = c
    _, err := s.Db.Exec("update Student set NumInGroup = $1 where StudentID = $2", s.Number, s.Id)
    if err != nil {
        panic(err)
    }
}
func (s *Student) Set_year(c int) {
    s.Year = c
    _, err := s.Db.Exec("update Student set YearOfAdmission = $1 where StudentID = $2", s.Year, s.Id)
    if err != nil {
        panic(err)
    }
}

