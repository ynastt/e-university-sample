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
	// "github.com/gorilla/mux"
    "github.com/google/uuid"
)

type UserDetails struct {
	Username string
	Password string
	ID 		 []uint8
	Succes   bool
}

type Stud struct {
	Name       string
	Surname    string
	Patronymic string
	Email      string
	Group      string
	Courses    []string
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
	Instance int
	Recieved int
	Comment  string
}

type Attend struct {
	Num        int
	Theme      string
	Date       string
	Attendance string
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
	Name     string
	Mods     []Module
	Exam     Exam
	Sems_ok  bool
	Lects_ok bool
	Labs_ok  bool
	Rks_ok   bool
}

type CP struct {
	Subject        string
	Description    string
	StartDate      string
	Deadline       string
	student_id     []uint8
	ProjAssignment string
	TitleOfProject string
	RecievedScore  int
	DateOfPassing  string
}

var (
	err    error
	conn   *net.TCPConn
	data   UserDetails
	stud   Stud
	isAuth bool
)

type Session struct {
	Token  uuid.UUID
	UserId []uint8
}

var sessionTable []Session

// send_request - вспомогательная функция для передачи запроса с указанной командой
// и данными. Данные могут быть пустыми (nil).
func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{Command: command, Data: &raw})
}

func auth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Authentification")
	isAuth = false
	conn = connectToServer()
	defer conn.Close()
	t, err := template.ParseFiles("templs/auth.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "auth", nil)
}

func auth_again(w http.ResponseWriter, r *http.Request) {
	log.Printf("Again Authentification")
	isAuth = false
	conn = connectToServer()
	defer conn.Close()
	t, err := template.ParseFiles("templs/badauth.html", "templs/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "badauth", nil)
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
			ID:       nil,
			Succes:   false,
		}
	}
	// запрос на сервер
	conn = connectToServer()
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	// данные для передачи на сервер
	var info proto.LoginInfo
	info.Username = data.Username
	info.Password = data.Password
	// println(conn)
	// println(encoder)
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
				data.ID = info.ID
			}
		}
	default:
		log.Printf("error: server reports unknown status %q\n", resp.Status)
	}
	// print(data.Username, data.Password)
	// print(data.Succes)

	if data.Succes {
		log.Printf("Successful authentification")
		isAuth = true
		u := uuid.New()
		sessionTable = append(sessionTable, Session{u, data.ID})

		c := http.Cookie{
			Name:   "token",
			Value:  u.String(),
			MaxAge: 300,
		}

		http.SetCookie(w, &c)
		http.Redirect(w, r, "/student", 301)
	} else {
		isAuth = false
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		tokenC, err := r.Cookie("token")
		if err != nil {
			panic(err)
		}
		fmt.Println(sessionTable)
		for i, s := range sessionTable {
			if s.Token.String() == tokenC.Value {
				sessionTable[i] = sessionTable[len(sessionTable)-1]
				sessionTable = sessionTable[:len(sessionTable)-1]
			}
			fmt.Println(sessionTable)
		}
		http.Redirect(w, r, "/badlogout", 301)
	}
}

func student_main(w http.ResponseWriter, r *http.Request) {
	log.Printf("Student index page")
	if isAuth {
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
					stud.Email = info.Email
					stud.Group = info.Group
					stud.Courses = info.Courses
					println("student is:", stud.Surname)
				}
			}
		default:
			log.Printf("error: server reports unknown status %q\n", resp.Status)
		}
		var crs []Course
		var strct StructForPage
		m := configs.SubjectsLinks.Subjects
		for _, c := range stud.Courses {
			crs = append(crs, Course{CourseLink: m[c], CourseName: c})
		}
		strct.StudentInfo = stud
		strct.CourseInfo = crs
		err = t.ExecuteTemplate(w, "student_main", strct)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		isAuth = false
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		tokenC, err := r.Cookie("token")
		if err != nil {
			panic(err)
		}
		fmt.Println(sessionTable)
		for i, s := range sessionTable {
			if s.Token.String() == tokenC.Value {
				sessionTable[i] = sessionTable[len(sessionTable)-1]
				sessionTable = sessionTable[:len(sessionTable)-1]
			}
			fmt.Println(sessionTable)
		}
		http.Redirect(w, r, "/badlogout", 301)
	}

}

func courseproject(w http.ResponseWriter, r *http.Request) {
	log.Printf("courseproject page")
	if isAuth {
		t, err := template.ParseFiles("templs/coursework.html", "templs/header.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		// запрос на сервер
		conn = connectToServer()
		defer conn.Close()
		encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
		var strct CP
		// данные для передачи на сервер
		var info proto.CP
		info.Subject = stud.Email

		send_request(encoder, "project", &info)
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
				var info proto.CP
				if err := json.Unmarshal(*resp.Data, &info); err != nil {
					log.Printf("error: malformed data field in response\n")
				} else {
					log.Printf("result: {%s, %s, %d}\n", info.Subject, info.TitleOfProject, info.RecievedScore)
					conn.Close()
					strct.Subject = info.Subject
					strct.Description = info.Description
					strct.StartDate = info.StartDate[0:10]
					strct.Deadline = info.Deadline[0:10]
					strct.student_id = info.Student_id
					strct.ProjAssignment = info.ProjAssignment
					strct.TitleOfProject = info.TitleOfProject
					strct.RecievedScore = info.RecievedScore
					strct.DateOfPassing = info.DateOfPassing[0:10]

				}
			}
		default:
			log.Printf("error: server reports unknown status %q\n", resp.Status)
		}

		err = t.ExecuteTemplate(w, "coursework", strct)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		isAuth = false
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		tokenC, err := r.Cookie("token")
		if err != nil {
			panic(err)
		}
		fmt.Println(sessionTable)
		for i, s := range sessionTable {
			if s.Token.String() == tokenC.Value {
				sessionTable[i] = sessionTable[len(sessionTable)-1]
				sessionTable = sessionTable[:len(sessionTable)-1]
			}
			fmt.Println(sessionTable)
		}		
		http.Redirect(w, r, "/badlogout", 301)
	}
}

func dbases(w http.ResponseWriter, r *http.Request) {
	log.Printf("subject db page")
	if isAuth {
		t, err := template.ParseFiles("templs/subject.html", "templs/header.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		crsname := "Базы данных"
		// запрос на сервер
		conn = connectToServer()
		defer conn.Close()
		encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
		var strct Subject
		var modules []Module
		var ex Exam
		// данные для передачи на сервер
		var info proto.SubjectInfo
		m := configs.SubjectsLinks.Subjects
		_, ok := m[strings.ToUpper(crsname)]
		if !ok {
			log.Fatal("error: unknown subject")
		} else {
			info.Name = strings.ToUpper(crsname)
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

					strct.Sems_ok = false
					strct.Rks_ok = false
					strct.Labs_ok = false
					strct.Lects_ok = false
					for _, m := range info.Mods {
						var mym Module
						mym.ModNumber = m.ModNumber
						mym.Labs = make([]Lab, 0)
						mym.Lects = make([]Attend, 0)
						mym.Sems = make([]Attend, 0)
						mym.Rks = make([]RK, 0)
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
							if l.Comment.Valid {
								myl.Comment = l.Comment.String
							} else {
								myl.Comment = "-"
							}
							myl.Date = l.Date[0:10]
							myl.Max, myl.Min, myl.Name = l.Max, l.Min, l.Name
							myl.Num = l.Num
							myl.Deadline, myl.Text = l.Deadline[0:10], l.Text
							mym.Labs = append(mym.Labs, myl)
							if len(mym.Labs) > 0 {
								strct.Labs_ok = true
							}
						}
						println("labs info")
						for _, r := range m.Rks {
							var myr RK
							if r.Recieved.Valid {
								myr.Recieved = int(r.Recieved.Int64)
							} else {
								myr.Recieved = 0
							}
							if r.Instance.Valid {
								myr.Instance = int(r.Instance.Int64)
							} else {
								myr.Instance = 0
							}
							if r.Variant.Valid {
								myr.Variant = int(r.Variant.Int64)
							} else {
								myr.Variant = 0
							}
							if r.Comment.Valid {
								myr.Comment = r.Comment.String
							} else {
								myr.Comment = "-"
							}
							if r.Date.Valid {
								myr.Date = r.Date.String[0:10]
							} else {
								myr.Date = "-"
							}
							println("date:", myr.Date)
							myr.Max, myr.Min = r.Max, r.Min
							myr.Num = r.Num
							mym.Rks = append(mym.Rks, myr)
							if len(mym.Rks) > 0 {
								strct.Rks_ok = true
							}
						}
						println("rk info")
						// println(strct.Rks_ok)
						for _, s := range m.Sems {
							var mys Attend
							if s.Bonus.Valid {
								mys.Bonus = int(s.Bonus.Int64)
							} else {
								mys.Bonus = 0
							}
							mys.Date = s.Date[0:10]
							if s.Attendance {
								mys.Attendance = "+"
							} else {
								mys.Attendance = "-"
							}
							mys.Num, mys.Theme = s.Num, s.Theme
							mym.Sems = append(mym.Sems, mys)
							if len(mym.Sems) > 0 {
								strct.Sems_ok = true
							}
						}
						println("sems info")
						for _, l := range m.Lects {
							var myl Attend
							if l.Bonus.Valid {
								myl.Bonus = int(l.Bonus.Int64)
							} else {
								myl.Bonus = 0
							}
							myl.Date = l.Date[0:10]
							if l.Attendance {
								myl.Attendance = "+"
							} else {
								myl.Attendance = "-"
							}
							myl.Num, myl.Theme = l.Num, l.Theme
							mym.Lects = append(mym.Lects, myl)
							if len(mym.Lects) > 0 {
								strct.Lects_ok = true
							}
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

		strct.Name = crsname
		strct.Mods = modules
		strct.Exam = ex
		err = t.ExecuteTemplate(w, "subject", strct)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		isAuth = false
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		tokenC, err := r.Cookie("token")
		if err != nil {
			panic(err)
		}
		fmt.Println(sessionTable)
		for i, s := range sessionTable {
			if s.Token.String() == tokenC.Value {
				sessionTable[i] = sessionTable[len(sessionTable)-1]
				sessionTable = sessionTable[:len(sessionTable)-1]
			}
			fmt.Println(sessionTable)
		}
		http.Redirect(w, r, "/badlogout", 301)
	}
}

func handleFunc() {
	http.HandleFunc("/logout", auth)
	http.HandleFunc("/badlogout", auth_again)
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
