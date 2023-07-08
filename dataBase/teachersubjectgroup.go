package dataBase

import (
	"database/sql"
)

// рубежный контроль в модуле
type teachersubjectgroup struct {
	TeacherID []uint8
	SubjectID []uint8
	GroupID []uint8
	TeacherRole int
	Db *sql.DB
}


