package dataBase 

import (
    "database/sql"
    "encoding/json"
)

// курсовая работа по дисциплине
type CourseProject struct {
    Id []uint8
    Subject string
    Description string
	Hours int
	StartDate string
	Deadline string
	Db *sql.DB
}

func (p CourseProject) Get_id() []uint8 { return p.Id }
func (p CourseProject) Get_subject() string { 
	return p.Subject
}
func (p CourseProject) Get_description() string{ return p.Description }
func (p CourseProject) Get_hours() int { return p.Hours }
func (p CourseProject) Get_start_date() string { return p.StartDate }
func (p CourseProject) Get_deadline() string { return p.Deadline }

func (p *CourseProject) Set_subject(name string) {
	p.Subject = name
    _, err := p.Db.Exec("update CourseProject SET Subject = $1 where ProjectID = $2", p.Subject, p.Id)
    if err != nil {
        panic(err)
    }
}

func (p *CourseProject) Set_description(text string) {
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

var supervisor_roles_counter = 2
var supervisor_roles = map[string]int {
    "Старший руководитель курсовыми проектами": 1,
	"Научный руководитель студента": 2,
}

// руководитель по курсовой работе
type Supervisor struct {
    TeacherId []uint8
	ProjectId []uint8
	Role int
	Db *sql.DB
}

func (s Supervisor) Get_id() ([]uint8, []uint8) { return s.TeacherId, s.ProjectId }
func (s Supervisor) Get_role() string {
	r, ok := mapkey(supervisor_roles, s.Role)
	if !ok {
  		panic("there is no such supervisor role in database")
	}
	return r
}

func (s *Supervisor) Set_role(name string) {
    if _, ok := supervisor_roles[name]; !ok {
		supervisor_roles_counter += 1
		supervisor_roles[name] = supervisor_roles_counter
	}
	s.Role = supervisor_roles[name]
	_, err := s.Db.Exec("update Supervisor SET SupervisorRole = $1 where teacher_id = $2 and project_id = $3", s.Role, s.TeacherId, s.ProjectId)
    if err != nil {
        panic(err)
    }
}

// курсовой проект студента
type StudentCourseProject struct {
    StudentId []uint8
	ProjectId []uint8
    Assignment json.RawMessage
    Title string
	Score int
    Date string
	Db *sql.DB
}

func (s StudentCourseProject) Get_id() ([]uint8, []uint8) { return s.StudentId, s.ProjectId }
func (s StudentCourseProject) Get_assignment() string {
    j, err := json.Marshal(s.Assignment)
	if err != nil {
		panic(err)
	}
	return string(j) 
}
func (s StudentCourseProject) Get_title() string { return s.Title }
func (s StudentCourseProject) Get_score() int { return s.Score }
func (s StudentCourseProject) Get_date() string { return s.Date }

func (s *StudentCourseProject) Set_assignment(text []byte) {
    s.Assignment = json.RawMessage(text)
	_, err := s.Db.Exec("update StudentCourseProject SET ProjAssignment = $1 where student_id = $2 and project_id =$3", s.Assignment, s.StudentId, s.ProjectId)
    if err != nil {
        panic(err)
    }
}
func (s *StudentCourseProject) Set_title(name string) { 
    s.Title = name
	_, err := s.Db.Exec("update StudentCourseProject SET TitleOfProject = $1 where student_id = $2 and project_id =$3", s.Title, s.StudentId, s.ProjectId)
    if err != nil {
        panic(err)
    }
}
func (s *StudentCourseProject) Set_score(score int) { 
    s.Score = score
	_, err := s.Db.Exec("update StudentCourseProject SET RecievedScore = $1 where student_id = $2 and project_id =$3", s.Score, s.StudentId, s.ProjectId)
    if err != nil {
        panic(err)
    }
}
func (s *StudentCourseProject) Set_date(date string) {
    s.Date = date
	_, err := s.Db.Exec("update StudentCourseProject SET DateOfPassing = $1 where student_id = $2 and project_id =$3", s.Date, s.StudentId, s.ProjectId)
    if err != nil {
        panic(err)
    }
}