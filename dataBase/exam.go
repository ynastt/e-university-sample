package dataBase

import (
	"database/sql"
)

// экзамен по дисциплине
type Exam struct {
    Id []uint8
	SubjectID []uint8
	Questions string
	MaxScore int
	MinScore int
	Date string // like that '2005-01-01'
	Db *sql.DB
}


func (e Exam) Get_id() ([]uint8, []uint8) { return e.Id, e.SubjectID }
func (e Exam) Get_questions() string { 
	return e.Questions 
}
func (e Exam) Get_max_score() int { return e.MaxScore }
func (e Exam) Get_min_score() int { return e.MinScore }
func (e Exam) Get_date() string { return e.Date }


func (e *Exam) Set_questions(text string) {
    e.Questions = text
	_, err := e.Db.Exec("update Exam SET Questions = $1 where ExamID = $2", e.Questions, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_max_score(score int) {
    e.MaxScore = score
	_, err := e.Db.Exec("update Exam SET Maxscore = $1 where ExamID = $2", e.MaxScore, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_min_score(score int) {
    e.MinScore = score
	_, err := e.Db.Exec("update Exam SET MinScore = $1 where ExamID = $2", e.MinScore, e.Id)
    if err != nil {
        panic(err)
    }
}

func (e *Exam) Set_date(date string) {
    e.Date = date
	_, err := e.Db.Exec("update Exam SET ExamDate = $1 where ExamID = $2", e.Date, e.Id)
    if err != nil {
        panic(err)
    }
}

// экзамен студента
type ExamInstance struct {
    StudentId []uint8
	ExamID []uint8
    Date string 
	NumOfInstance int
	Score int
    Ticket int
	Db *sql.DB
}

func (i ExamInstance) Get_id() ([]uint8, []uint8) { return i.StudentId, i.ExamID }
func (i ExamInstance) Get_date() string { return i.Date }
func (i ExamInstance) Get_num_of_instance() int { return i.NumOfInstance }
func (i ExamInstance) Get_score() int { return i.Score }
func (i ExamInstance) Get_ticket() int { return i.Ticket }

func (i *ExamInstance) Set_date(date string) {
    i.Date = date
	_, err := i.Db.Exec("update ExamInstance SET DateOfPassing = $1 where student_id = $2 and exam_id = $3", i.Date, i.StudentId, i.ExamID)
    if err != nil {
        panic(err)
    }
}
func (i *ExamInstance) Set_num_of_instance(num int)  {
    i.NumOfInstance = num
	_, err := i.Db.Exec("update ExamInstance SET NumOfInstance = $1 where student_id = $2 and exam_id = $3", i.NumOfInstance, i.StudentId, i.ExamID)
    if err != nil {
        panic(err)
    }
}
func (i *ExamInstance) Set_score(score int)  {
    i.Score = score
	_, err := i.Db.Exec("update ExamInstance SET RecievedScore = $1 where student_id = $2 and exam_id = $3", i.Score, i.StudentId, i.ExamID)
    if err != nil {
        panic(err)
    }
}
func (i *ExamInstance) Set_ticket(v int)  {
    i.Ticket = v
	_, err := i.Db.Exec("update LExamInstance SET TicketNumber = $1 where student_id = $2 and exam_id = $3", i.Ticket, i.StudentId, i.ExamID)
    if err != nil {
        panic(err)
    }
}
