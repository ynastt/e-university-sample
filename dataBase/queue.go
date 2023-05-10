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
