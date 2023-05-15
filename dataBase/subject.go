package dataBase 

import (
    "database/sql"
    "encoding/json"
)

// дисциплина
type Subject struct {
    Id []uint8
    Description json.RawMessage
	Program json.RawMessage
	Hours int
	Credits int
	Db *sql.DB
}

func (s Subject) Get_id() []uint8 { return s.Id }
func (s Subject) Get_description() string {
    j, err := json.Marshal(s.Description)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (s Subject) Get_program() string {
    j, err := json.Marshal(s.Program)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (s Subject) Get_hours() int { return s.Hours }
func (s Subject) Get_credits() int { return s.Credits }

func (s *Subject) Set_description(name1 []byte) {
    s.Description = json.RawMessage(name1)
    _, err := s.Db.Exec("update Subject SET Description = $1 where SubjectID = $2", s.Description, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Subject) Set_program(name1 []byte) {
    s.Program = json.RawMessage(name1)
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

var teacher_roles_counter = 3
var teacher_roles = map[string]int {
    "Лектор": 1,
	"Семинарист": 2,
    "Лаборант": 3,
}

// преподаватель по дисциплине
type TeacherSubject struct {
    TeacherId []uint8
	SubjectId []uint8
	Role int
	Db *sql.DB
}

func (t TeacherSubject) Get_id() ([]uint8, []uint8) { return t.TeacherId, t.SubjectId }
func (t TeacherSubject) Get_role() string {
	r, ok := mapkey(teacher_roles, t.Role)
	if !ok {
  		panic("there is no such teacher role in database")
	}
	return r
}

func (t *TeacherSubject) Set_role(name string) {
    if _, ok := teacher_roles[name]; !ok {
		teacher_roles_counter += 1
		teacher_roles[name] = teacher_roles_counter
	}
	t.Role = teacher_roles[name]
	_, err := t.Db.Exec("update TeacherSubject SET TeacherRole = $1 where teacher_id = $2 and subject_id = $3", t.Role, t.TeacherId, t.SubjectId)
    if err != nil {
        panic(err)
    }
}