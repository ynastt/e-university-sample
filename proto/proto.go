package proto

import (
	"database/sql"
	"encoding/json"
)

// Request -- запрос клиента к серверу.
type Request struct {
	// Поле Command может принимать три значения:
	// * "quit" - прощание с сервером (после этого сервер рвёт соединение);
	// * "check" - передача данных авторизации и просьба проверить логин и пароль, существует ли такой пользователь.
	// * "student" - поиск студента в соответствии с переданным логином пользвоателя
	Command string `json:"command"`

	// Если Command == "check", в поле Data должна лежать структура с  логином и паролем пользователя.
	// Если Command == "student", в поле Data должна лежать структура с ФИО пользователя.
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// Response -- ответ сервера клиенту.
type Response struct {
	// Поле Status может принимать три значения:
	// * "quited" - успешное выполнение команды "quit";
	// * "failed" - в процессе выполнения команды произошла ошибка;
	// * "ok" - проверка данных авторизации прошла успешно.
	Status string `json:"status"`

	// Если Status == "failed", то в поле Data находится сообщение об ошибке.
	// Если Status == "ok", поле Data не пустое,
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// LoginInfo -- логин и пароль.
type LoginInfo struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Exists   bool   `json:"exists"`
}

// StudInfo -- ФИО студента.
type StudInfo struct {
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	Patronymic string   `json:"patronymic"`
	Email      string   `json:"studemail"`
	Group      string   `json:"group"`
	Courses    []string `json:"courses"`
}

type Lab struct {
	Num       int            `json:"labnum"`
	Date      string         `json:"labdate"`
	Deadline  string         `json:"labdeadline"`
	Name      string         `json:"labname"`
	Text      string         `json:"labtext"`
	Min       int            `json:"labmin"`
	Max       int            `json:"labmax"`
	Recieved  sql.NullInt64  `json:"labscore"`
	Instance  sql.NullInt64  `json:"labinstance"`
	Bonus     sql.NullInt64  `json:"labbonus"`
	Comment   sql.NullString `json:"labcomment"`
	Module_id []uint8        `json:"labmoduleid"`
}

type RK struct {
	Num       int            `json:"rknum"`
	Date      sql.NullString `json:"rkdate"`
	Min       int            `json:"rkmin"`
	Max       int            `json:"rkmax"`
	Variant   sql.NullInt64  `json:"rkvariant"`
	Instance  sql.NullInt64  `json:"rkinstance"`
	Recieved  sql.NullInt64  `json:"rkscore"`
	Comment   sql.NullString `json:"rkcomment"`
	Module_id []uint8        `json:"rkmoduleid"`
}

type Attend struct {
	Num        int           `json:"eventnumber"`
	Theme      string        `json:"eventtheme"`
	Date       string        `json:"eventdate"`
	Attendance bool          `json:"eventattendance"`
	Bonus      sql.NullInt64 `json:"eventbonus"`
	Module_id  []uint8       `json:"attmoduleid"`
}

type Module struct {
	ModNumber int      `json:"modnum"`
	Labs      []Lab    `json:"modlabs"`
	Rks       []RK     `json:"modrks"`
	Lects     []Attend `json:"modlects"`
	Sems      []Attend `json:"modsems"`
}

type Exam struct {
	Date string `json:"exdate"`
	Min  int    `json:"exmin"`
	Max  int    `json:"exmax"`
}

// SubjectInfo - информация по дисциплине(модули, Лр, РК, лекции, Семинары)
type SubjectInfo struct {
	Name string   `json:"subjectname"`
	Mods []Module `json:"subjectmodules"`
	Exam Exam     `json:"subjectexam"`
}

type CP struct {
	Subject        string  `json:"cpsubj"`
	Description    string  `json:"cpdescript"`
	StartDate      string  `json:"cpstartdate"`
	Deadline       string  `json:"cpdeadline"`
	Student_id     []uint8 `json:"cpstudid"`
	ProjAssignment string  `json:"cpassignment"`
	TitleOfProject string  `json:"cptitle"`
	RecievedScore  int     `json:"cpscore"`
	DateOfPassing  string  `json:"cppass"`
}
