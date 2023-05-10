package dataBase 

import (
    "database/sql"
)

type Subject struct {
    Id []uint8
    Description []byte
	Program []byte
	Hours int
	Credits int
	Db *sql.DB
}

func (s Subject) Get_id() []uint8 { return s.Id }
func (s Subject) Get_description() []byte { return s.Description }
func (s Subject) Get_program() []byte { return s.Program }
func (s Subject) Get_hours() int { return s.Hours }
func (s Subject) Get_credits() int { return s.Credits }

func (s *Subject) Set_description(name1 []byte) {
    s.Description = name1
    _, err := s.Db.Exec("update Subject SET Description = $1 where SubjectID = $2", s.Description, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Subject) Set_program(name1 []byte) {
    s.Program = name1
    _, err := s.Db.Exec("update Subject SET SubjectProgram = $1 where SubjectID = $2", s.Program, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Subject) Set_hours(name1 int) {
    s.Hours = name1
    _, err := s.Db.Exec("update Subject SET NumberOfHours = $1 where SubjectID = $2", s.Hours, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Subject) Set_credits(name1 int) {
    s.Credits = name1
    _, err := s.Db.Exec("update Subject SET NumberOfCredits = $1 where SubjectID = $2", s.Credits, s.Id)
    if err != nil {
        panic(err)
    }
}