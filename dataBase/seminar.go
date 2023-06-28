package dataBase

import (
	"database/sql"
)

// семинар в модуле
type Seminar struct {
    Id []uint8
	ModuleID []uint8
	Theme string
	Text string
	Db *sql.DB
}


func (s Seminar) Get_id() ([]uint8, []uint8) { return s.Id, s.ModuleID }
func (s Seminar) Get_text() string { 
	return s.Text 
}
func (s Seminar) Get_theme() string { return s.Theme }


func (s *Seminar) Set_text(text string) {
    s.Text = text
	_, err := s.Db.Exec("update Seminar SET Text = $1 where SeminarID = $2", s.Text, s.Id)
    if err != nil {
        panic(err)
    }
}

func (s *Seminar) Set_theme(t string) {
    s.Theme = t
	_, err := s.Db.Exec("update Seminar SET Theme = $1 where SeminarID = $2", s.Theme, s.Id)
    if err != nil {
        panic(err)
    }
}

// посещение семинара студентом
type SeminarAttendance struct {
    StudentId []uint8
	SeminarID []uint8
	Attendance bool
	BonusScore int
	Db *sql.DB
}

func (a SeminarAttendance) Get_id() ([]uint8, []uint8) { return a.StudentId, a.SeminarID }
func (a SeminarAttendance) Get_attendance() bool { return a.Attendance }
func (a SeminarAttendance) Get_bonus_score() int { return a.BonusScore }

func (a *SeminarAttendance) Set_attendance(was bool) {
    a.Attendance = was
	_, err := a.Db.Exec("update SeminarAttendance SET WasAttended = $1 where student_id = $2 and seminar_id = $3", a.Attendance, a.StudentId, a.SeminarID)
    if err != nil {
        panic(err)
    }
}

func (a *SeminarAttendance) Set_bonus_score(score int) {
    a.BonusScore = score
	_, err := a.Db.Exec("update SeminarAttendance SET BonusScore = $1 where student_id = $2 and seminar_id = $3", a.BonusScore, a.StudentId, a.SeminarID)
    if err != nil {
        panic(err)
    }
}