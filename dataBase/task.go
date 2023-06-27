package dataBase

import (
	"database/sql"
)

// задание студента по дисциплине
type Task struct {
    TaskId []uint8
	StudentID []uint8
	SubjectID []uint8
	Description string
	MaxScore int
	MinScore int
	Deadline string // like that '2005-01-01'
	RecievedScore int
	Db *sql.DB
}


func (t Task) Get_id() ([]uint8, []uint8, []uint8) { return t.TaskId, t.StudentID, t.SubjectID }
func (t Task) Get_description() string { return t.Description }
func (t Task) Get_max_score() int { return t.MaxScore }
func (t Task) Get_min_score() int { return t.MinScore }
func (t Task) Get_deadline() string { return t.Deadline }
func (t Task) Get_score() int { return t.RecievedScore }


func (t *Task) Set_description(text string) {
    t.Description = text
	_, err := t.Db.Exec("update Task SET Description = $1 where TaskID = $2 and student_id = $3 and subject_id = $4", t.Description, t.TaskId, t.StudentID, t.SubjectID)
    if err != nil {
        panic(err)
    }
}

func (t *Task) Set_max_score(score int) {
    t.MaxScore = score
	_, err := t.Db.Exec("update Task SET MaxScore = $1 where TaskID = $2 and student_id = $3 and subject_id = $4", t.MaxScore, t.TaskId, t.StudentID, t.SubjectID)
    if err != nil {
        panic(err)
    }
}

func (t *Task) Set_min_score(score int) {
    t.MinScore = score
	_, err := t.Db.Exec("update Task SET MinScore = $1 where TaskID = $2 and student_id = $3 and subject_id = $4", t.MinScore, t.TaskId, t.StudentID, t.SubjectID)
    if err != nil {
        panic(err)
    }
}

func (t *Task) Set_deadline(date string) {
    t.Deadline = date
	_, err := t.Db.Exec("update Task SET Deadline = $1 where TaskID = $2 and student_id = $3 and subject_id = $4", t.Deadline, t.TaskId, t.StudentID, t.SubjectID)
    if err != nil {
        panic(err)
    }
}

func (t *Task) Set_score(score int) {
    t.RecievedScore = score
	_, err := t.Db.Exec("update Task SET RecievedScore = $1 where TaskID = $2 and student_id = $3 and subject_id = $4", t.RecievedScore, t.TaskId, t.StudentID, t.SubjectID)
    if err != nil {
        panic(err)
    }
}

// задание экзамена
type TaskExam struct {
    TaskId []uint8
	ExamID []uint8
	Description string
	MaxScore int
	MinScore int
	Deadline string // like that '2005-01-01'
	RecievedScore int
	Db *sql.DB
}


func (t TaskExam) Get_id() ([]uint8, []uint8) { return t.TaskId, t.ExamID }
func (t TaskExam) Get_description() string { return t.Description }
func (t TaskExam) Get_max_score() int { return t.MaxScore }
func (t TaskExam) Get_min_score() int { return t.MinScore }
func (t TaskExam) Get_deadline() string { return t.Deadline }
func (t TaskExam) Get_score() int { return t.RecievedScore }


func (t *TaskExam) Set_description(text string) {
    t.Description = text
	_, err := t.Db.Exec("update TaskExam SET Description = $1 where TaskID = $2 and exam_id = $3", t.Description, t.TaskId, t.ExamID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskExam) Set_max_score(score int) {
    t.MaxScore = score
	_, err := t.Db.Exec("update TaskExam SET MaxScore = $1 where TaskID = $2 and exam_id = $3", t.MaxScore, t.TaskId, t.ExamID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskExam) Set_min_score(score int) {
    t.MinScore = score
	_, err := t.Db.Exec("update TaskExam SET MinScore = $1 where TaskID = $2 and exam_id = $3", t.MinScore, t.TaskId, t.ExamID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskExam) Set_deadline(date string) {
    t.Deadline = date
	_, err := t.Db.Exec("update TaskExam SET Deadline = $1 where TaskID = $2 and exam_id = $3", t.Deadline, t.TaskId, t.ExamID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskExam) Set_score(score int) {
    t.RecievedScore = score
	_, err := t.Db.Exec("update TaskExam SET RecievedScore = $1 where TaskID = $2 and exam_id = $3", t.RecievedScore, t.TaskId, t.ExamID)
    if err != nil {
        panic(err)
    }
}

// задание рубежного контроля
type TaskBC struct {
    TaskId []uint8
	BCID []uint8
	Description string
	MaxScore int
	MinScore int
	Deadline string // like that '2005-01-01'
	RecievedScore int
	Db *sql.DB
}


func (t TaskBC) Get_id() ([]uint8, []uint8) { return t.TaskId, t.BCID }
func (t TaskBC) Get_description() string { return t.Description }
func (t TaskBC) Get_max_score() int { return t.MaxScore }
func (t TaskBC) Get_min_score() int { return t.MinScore }
func (t TaskBC) Get_deadline() string { return t.Deadline }
func (t TaskBC) Get_score() int { return t.RecievedScore }


func (t *TaskBC) Set_description(text string) {
    t.Description = text
	_, err := t.Db.Exec("update TaskBC SET Description = $1 where TaskID = $2 and bc_id = $3", t.Description, t.TaskId, t.BCID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskBC) Set_max_score(score int) {
    t.MaxScore = score
	_, err := t.Db.Exec("update TaskBC SET MaxScore = $1 where TaskID = $2 and bc_id = $3", t.MaxScore, t.TaskId, t.BCID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskBC) Set_min_score(score int) {
    t.MinScore = score
	_, err := t.Db.Exec("update TaskBC SET MinScore = $1 where TaskID = $2 and bc_id = $3", t.MinScore, t.TaskId, t.BCID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskBC) Set_deadline(date string) {
    t.Deadline = date
	_, err := t.Db.Exec("update TaskBC SET Deadline = $1 where TaskID = $2 and bc_id = $3", t.Deadline, t.TaskId, t.BCID)
    if err != nil {
        panic(err)
    }
}

func (t *TaskBC) Set_score(score int) {
    t.RecievedScore = score
	_, err := t.Db.Exec("update TaskBC SET RecievedScore = $1 where TaskID = $2 and bc_id = $3", t.RecievedScore, t.TaskId, t.BCID)
    if err != nil {
        panic(err)
    }
}