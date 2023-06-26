package main

import (
	"e-university-sample/configs"
	proto "e-university-sample/proto"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	"github.com/skratchdot/open-golang/open"
)

type UserDetails struct {
	Username string
	Password string
	Succes   bool
}

type Stud struct {
	Name       string
	Surname    string
	Patronymic string
	Group      string
}

type Course struct {
	CourseLink string
	CourseName string
	CourseID   []uint8
}

type StructForPage struct {
	StudentInfo Stud
	CourseInfo  []Course
}

type Lab struct {
	Num      int
	Date     string
	Deadline string
	Name     string
	Text     string
	Min      int
	Max      int
	Instance int
	Recieved int
	Bonus    int
	Comment  string
}

type RK struct {
	Num      int
	Date     string
	Min      int
	Max      int
	Variant  int
	Recieved int
	Comment  string
}

type Attend struct {
	Num        int
	Theme      string
	Date       string
	Attendance bool
	Bonus      int
}

type Module struct {
	ModNumber int
	Labs      []Lab
	Rks       []RK
	Lects     []Attend
	Sems      []Attend
}

type Exam struct {
	Date string
	Min  int
	Max  int
}

type Subject struct {
	Name string
	Mods []Module
	Exam Exam
}

var err error
var conn *net.TCPConn
var data UserDetails
var stud Stud

// send_request - вспомогательная функция для передачи запроса с указанной командой
// и данными. Данные могут быть пустыми (nil).
func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{Command: command, Data: &raw})
}

func auth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Authentification")
	t, err := template.ParseFiles("templs/auth.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "auth", nil)
}

func check_student(w http.ResponseWriter, r *http.Request) {
	err = r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if r.FormValue("username") != "" && r.FormValue("userpassword") != "" {
		data = UserDetails{
			Username: r.FormValue("username"),
			Password: r.FormValue("userpassword"),
			Succes:   false,
		}
	}
	// запрос на сервер
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	// данные для передачи на сервер
	var info proto.LoginInfo
	info.Username = data.Username
	info.Password = data.Password
	println(conn)
	println(encoder)
	send_request(encoder, "check", &info)
	// Получение ответа с сервера
	var resp proto.Response
	if err := decoder.Decode(&resp); err != nil {
		log.Printf("error: %v\n", err)
	}
	switch resp.Status {
	case "failed":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var errorMsg string
			if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("failed: %s\n", errorMsg)
			}
		}
	case "ok":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var info proto.LoginInfo
			if err := json.Unmarshal(*resp.Data, &info); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("result: %v\n", info.Exists)
				data.Succes = info.Exists
			}
		}
	default:
		log.Printf("error: server reports unknown status %q\n", resp.Status)
	}
	// print(data.Username, data.Password)
	// print(data.Succes)
	if data.Succes {
		log.Printf("Successful authentification")
		http.Redirect(w, r, "/student", 301)
	} else {
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		http.Redirect(w, r, "/logout", 301)
	}
}

func student_main(w http.ResponseWriter, r *http.Request) {
	log.Printf("Student index page")
	t, err := template.ParseFiles("templs/student_main.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// запрос на сервер
	conn = connectToServer()
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	// данные для передачи на сервер
	var info proto.LoginInfo
	info.Username = data.Username
	send_request(encoder, "student", &info)
	// Получение ответа с сервера
	var resp proto.Response
	if err := decoder.Decode(&resp); err != nil {
		log.Printf("error: %v\n", err)
		println("here")
	}
	switch resp.Status {
	case "failed":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var errorMsg string
			if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("failed: %s\n", errorMsg)
			}
		}
	case "ok":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var info proto.StudInfo
			if err := json.Unmarshal(*resp.Data, &info); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("result: {%s, %s, %s, %s}\n", info.Name, info.Surname, info.Patronymic, info.Group)
				stud.Name = info.Name
				stud.Surname = info.Surname
				stud.Patronymic = info.Patronymic
				stud.Group = info.Group
				println("student is:", stud.Surname)
			}
		}
	default:
		log.Printf("error: server reports unknown status %q\n", resp.Status)
	}
	var crs []Course
	var strct StructForPage
	m := configs.SubjectsForStudents.Subjects
	for _, course := range m[stud.Group] {
		crs = append(crs, Course{CourseLink: course.Link, CourseName: course.Name})
	}
	strct.StudentInfo = stud
	strct.CourseInfo = crs
	err = t.ExecuteTemplate(w, "student_main", strct)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func courseproject(w http.ResponseWriter, r *http.Request) {
	log.Printf("courseproject page")
	t, err := template.ParseFiles("templs/coursework.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	err = t.ExecuteTemplate(w, "coursework", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func dbases(w http.ResponseWriter, r *http.Request) {
	log.Printf("subject db page")
	t, err := template.ParseFiles("templs/dbases_subject.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	crsname := "Базы данных"
	// запрос на сервер
	conn = connectToServer()
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	var modules []Module
	var ex Exam
	// данные для передачи на сервер
	var info proto.SubjectInfo
	m := configs.SubjectsForStudents.Subjects
	for i, course := range m[stud.Group] {
		if strings.ToUpper(crsname) == course.Name {
			info.Name = course.Name
			break
		}
		if i == len(m[stud.Group])-1 && strings.ToUpper(crsname) != course.Name {
			log.Fatal("error: unknown subject")
		}
	}
	info.Mods = []proto.Module{}
	info.Exam = proto.Exam{}
	send_request(encoder, "databases", &info)
	// Получение ответа с сервера
	var resp proto.Response
	if err := decoder.Decode(&resp); err != nil {
		println("here")
		log.Printf("error: %v\n", err)
	}
	switch resp.Status {
	case "failed":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var errorMsg string
			if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("failed: %s\n", errorMsg)
			}
		}
	case "ok":
		if resp.Data == nil {
			log.Printf("error: data field is absent in response\n")
		} else {
			var info proto.SubjectInfo
			if err := json.Unmarshal(*resp.Data, &info); err != nil {
				log.Printf("error: malformed data field in response\n")
			} else {
				log.Printf("result: {%s, %s, %d}\n", info.Name, info.Exam.Date, info.Mods[0].ModNumber)
				conn.Close()
				ex.Date = info.Exam.Date[0:10] // срез для нормального вида даты
				ex.Max = info.Exam.Max
				ex.Min = info.Exam.Min

				for _, m := range info.Mods {
					var mym Module
					mym.ModNumber = m.ModNumber
					for _, l := range m.Labs {
						var myl Lab
						if l.Bonus.Valid {
							myl.Bonus = int(l.Bonus.Int64)
						} else {
							myl.Bonus = 0
						}
						if l.Recieved.Valid {
							myl.Recieved = int(l.Recieved.Int64)
						} else {
							myl.Recieved = 0
						}
						if l.Instance.Valid {
							myl.Instance = int(l.Instance.Int64)
						} else {
							myl.Instance = 0
						}
						myl.Comment, myl.Date =  l.Comment, l.Date[0:10]
						myl.Max, myl.Min, myl.Name = l.Max, l.Min, l.Name
						myl.Num = l.Num
						myl.Deadline, myl.Text = l.Deadline[0:10], l.Text
						mym.Labs = append(mym.Labs, myl)
					}
					println("labs info")
					for _, r := range m.Rks {
						var myr RK
						myr.Comment, myr.Date = r.Comment, r.Date[0:10]
						myr.Max, myr.Min = r.Max, r.Min
						myr.Num, myr.Recieved, myr.Variant = r.Num, r.Recieved, r.Variant
						mym.Rks = append(mym.Rks, myr)
					}
					println("rk info")
					for _, s := range m.Sems {
						var mys Attend
						mys.Bonus, mys.Date = s.Bonus, s.Date[0:10]
						mys.Attendance = s.Attendance
						mys.Num, mys.Theme = s.Num, s.Theme
						mym.Sems = append(mym.Sems, mys)
					}
					println("sems info")
					for _, l := range m.Lects {
						var myl Attend
						myl.Bonus, myl.Date = l.Bonus, l.Date[0:10]
						myl.Attendance = l.Attendance
						myl.Num, myl.Theme = l.Num, l.Theme
						mym.Lects = append(mym.Lects, myl)
					}
					println("lect info")
					modules = append(modules, mym)
					println("modules info")
				}
				println("subject info")
			}
		}
	default:
		log.Printf("error: server reports unknown status %q\n", resp.Status)
	}
	var strct Subject
	strct.Name = crsname
	strct.Mods = modules
	strct.Exam = ex
	err = t.ExecuteTemplate(w, "dbases_subject", strct)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func handleFunc() {
	http.HandleFunc("/logout", auth)
	http.HandleFunc("/login", check_student)
	http.HandleFunc("/student", student_main)
	http.HandleFunc("/subject/databases", dbases)
	http.HandleFunc("/coursework", courseproject)
	// http.HandleFunc("/subject/rprp", rprp__page)
	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func connectToServer() *net.TCPConn {
	var addrStr string
	addrStr = "127.0.0.1:6000"
	// Разбор адреса, установка соединения с сервером и
	// возвращение связи, для дальнейшего запуска цикла взаимодействия с сервером.
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else if conn, err := net.DialTCP("tcp", nil, addr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		return conn
	}
	return nil
}

func main() {
	conn = connectToServer()
	if conn != nil {
		defer conn.Close()
		err := open.Start("http://localhost/logout")
		if err != nil {
			log.Println(err)
		}
		handleFunc()
	}
}
