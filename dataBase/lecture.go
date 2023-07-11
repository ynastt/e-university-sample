package dataBase

import (
	"database/sql"
    "github.com/google/uuid"
)

// лекция в модуле
type Lecture struct {
    Id uuid.UUID
	ModuleID uuid.UUID
	Theme string
	Text string
	Db *sql.DB
}

func (l Lecture) Get_id() (uuid.UUID, uuid.UUID) { return l.Id, l.ModuleID }
func (l Lecture) Get_text() string { 
	return l.Text
}
func (l Lecture) Get_theme() string { return l.Theme }


func (l *Lecture) Set_text(text string) {
    l.Text = text
	_, err := l.Db.Exec("update Lecture SET Text = $1 where LectureID = $2", l.Text, l.Id)
    if err != nil {
        panic(err)
    }
}

func (l *Lecture) Set_theme(t string) {
    l.Theme = t
	_, err := l.Db.Exec("update Lecture SET Theme = $1 where LectureID = $2", l.Theme, l.Id)
    if err != nil {
        panic(err)
    }
}

// посещение лекции студентом
type LectureAttendance struct {
    StudentId uuid.UUID
	LectureID uuid.UUID
	Attendance bool
	BonusScore int
	Db *sql.DB
}

func (a LectureAttendance) Get_id() (uuid.UUID, uuid.UUID) { return a.StudentId, a.LectureID }
func (a LectureAttendance) Get_attendance() bool { return a.Attendance }
func (a LectureAttendance) Get_bonus_score() int { return a.BonusScore }

func (a *LectureAttendance) Set_attendance(was bool) {
    a.Attendance = was
	_, err := a.Db.Exec("update LectureAttendance SET WasAttended = $1 where student_id = $2 and lecture_id = $3", a.Attendance, a.StudentId, a.LectureID)
    if err != nil {
        panic(err)
    }
}

func (a *LectureAttendance) Set_bonus_score(score int) {
    a.BonusScore = score
	_, err := a.Db.Exec("update LectureAttendance SET BonusScore = $1 where student_id = $2 and lecture_id = $3", a.BonusScore, a.StudentId, a.LectureID)
    if err != nil {
        panic(err)
    }
}

