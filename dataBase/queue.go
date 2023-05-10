package dataBase 

import (
    "database/sql"
)


type Queue struct {
    Id []uint8
	StartDate string
	Db *sql.DB
}

func (q Queue) Get_id() []uint8 { return q.Id }
func (q Queue) Get_start_date() string { return q.StartDate }

func (q *Queue) Set_start_date(date string) {
    q.StartDate = date
	_, err := q.Db.Exec("update Queue SET StartDate = $1 where QueueID = $2", q.StartDate, q.Id)
    if err != nil {
        panic(err)
    }
}

var tasks_counter = 2
var tasks = map[string]int {
    "Лабораторная работа": 1,
	"Рубежный контроль": 2,
}

type StudentInQueue struct {
    StudentId []uint8
	QueueId []uint8
	Number int
	Task int
	Db *sql.DB
}

func (s StudentInQueue) Get_id() ([]uint8, []uint8) { return s.StudentId, s.QueueId }
func (s StudentInQueue) Get_order_number() int { return s.Number }
func (s StudentInQueue) Get_task() string {
	t, ok := mapkey(tasks, s.Task)
	if !ok {
  		panic("there is no such task in queue in database")
	}
	return t
}

func (s *StudentInQueue) Set_order_number(num int) {
    s.Number = num
	_, err := s.Db.Exec("update StudentInQueue SET NumInQueue = $1 where student_id = $2 and queue_id = $3", s.Number, s.StudentId, s.QueueId)
    if err != nil {
        panic(err)
    }
}

func (s *StudentInQueue) Set_task(name string) {
    if _, ok := tasks[name]; !ok {
		tasks_counter += 1
		tasks[name] = tasks_counter
	}
	s.Task = tasks[name]
	_, err := s.Db.Exec("update StudentInQueue SET Task = $1 where student_id = $2 and queue_id = $3", s.Task, s.StudentId, s.QueueId)
    if err != nil {
        panic(err)
    }
}