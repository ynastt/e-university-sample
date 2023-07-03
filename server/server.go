package main

import (
	"database/sql"
	"e-university-sample/configs"
	dbase "e-university-sample/dataBase"
	proto "e-university-sample/proto"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	fill "e-university-sample/fill"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	err   error
	conn  *net.TCPConn
	query string
)

func connectPostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.PostgresConfig.Host,
		configs.PostgresConfig.Port,
		configs.PostgresConfig.User,
		configs.PostgresConfig.Password,
		configs.PostgresConfig.DBname,
	)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("Successfully connected!\n\n")

}

type StudentClient struct {
	conn *net.TCPConn  // Объект TCP-соединения
	enc  *json.Encoder // Объект для кодирования и отправки сообщений
}

// NewClient - конструктор клиента, принимает в качестве параметра
// объект TCP-соединения.
func NewStudentClient(conn *net.TCPConn) *StudentClient {
	return &StudentClient{
		conn: conn,
		enc:  json.NewEncoder(conn),
	}
}

// serve - метод, в котором реализован цикл взаимодействия с клиентом.
// метод serve будет вызаваться в отдельной go-программе.
func (client *StudentClient) serve() {
	defer client.conn.Close()
	log.Println("new client serve")
	decoder := json.NewDecoder(client.conn)
	for {
		var req proto.Request
		if err := decoder.Decode(&req); err != nil {
			log.Println("client: cannot decode message", "reason ", err)
			break
		} else {
			log.Println("client received command", req.Command)
			client.handleRequest(&req)
		}
	}
}

// handleRequest - метод обработки запроса от клиента студента
func (client *StudentClient) handleRequest(req *proto.Request) {
	switch req.Command {
	// case "quit":
	// 	client.respond("ok", nil)
	// 	return true
	case "check":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var info proto.LoginInfo
			if err := json.Unmarshal(*req.Data, &info); err != nil {
				errorMsg = "malformed data field"
			} else {
				user := dbase.User{}
				row := db.QueryRow("select * from Users where Login = $1", info.Username)
				err = row.Scan(&user.Id, &user.Login, &user.Passw, &user.UserRights)
				if err != nil {
					errorMsg = err.Error()
				}
				fmt.Println(user.Id, user.Login, user.Passw, user.UserRights)

				if user.UserRights != 2 && user.UserRights != 1 {
					errorMsg = "it is neither student nor teacher"
				} else {
					if info.Username == user.Login && info.Password == user.Passw {
						info.Exists = true
					}
					log.Println("client: user exists")
					client.respond("ok", &proto.LoginInfo{
						Username: user.Get_login(),
						Password: user.Get_passw(),
						Exists:   true,
					})
				}
			}
		}
		if errorMsg != "" {
			log.Println("client: checking failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	case "student":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var info proto.LoginInfo
			if err := json.Unmarshal(*req.Data, &info); err != nil {
				errorMsg = "malformed data field"
			} else {
				user := dbase.User{}
				row := db.QueryRow("select * from Users where Login = $1", info.Username)
				err = row.Scan(&user.Id, &user.Login, &user.Passw, &user.UserRights)
				if err != nil {
					errorMsg = err.Error()
				}
				fmt.Println("user id:", user.Id)
				stud := dbase.Student{}
				g := dbase.StudentGroup{}
				row = db.QueryRow("select * from Student where user_id = $1", string(user.Id))
				err = row.Scan(&stud.Id, &stud.Name, &stud.Surname, &stud.Patronymic, &stud.Email, &stud.Phone, &stud.Year, &stud.Courses, &stud.Number, &stud.Userid, &stud.Groupid)
				if err != nil {
					errorMsg = err.Error()
				}
				println("student: ", stud.Name, "errr:", errorMsg)
				println("group id: ", stud.Groupid, "errr:", errorMsg)
				row = db.QueryRow("select GroupName from StudentGroup where GroupID = $1", stud.Groupid)
				err = row.Scan(&g.Name)
				if err != nil {
					errorMsg = err.Error()
				}
				crs := make([]string, 0)
				res, err := db.Query("select subject_id from TeacherSubjectGroup where group_id = (SELECT GroupID from StudentGroup where GroupName= $1)", g.Name)
				if err != nil {
					panic(err)
				}
				for res.Next() {
					r := dbase.Subject{}
					err = res.Scan(&r.Id)
					if err != nil {
						fmt.Println(err)
						continue
					}
					row := db.QueryRow("select Description from Subject where SubjectID = $1", r.Id)
					err = row.Scan(&r.Description)
					if err != nil {
						errorMsg = err.Error()
					}
					crs = append(crs, r.Description)
					fmt.Println(crs)
				}
				log.Println("client: student is found")
				log.Println(stud.Name, stud.Surname, stud.Patronymic, g.Name)
				log.Println(stud.Courses)
				client.respond("ok", &proto.StudInfo{
					Name:       stud.Get_name(),
					Surname:    stud.Get_surname(),
					Patronymic: stud.Get_patronymic(),
					Email:      stud.Get_email(),
					Group:      g.Name,
					Courses:    crs,
				})
			}
		}

		if errorMsg != "" {
			log.Println("client: finding student failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	case "databases":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var info proto.SubjectInfo
			if err := json.Unmarshal(*req.Data, &info); err != nil {
				errorMsg = "malformed data field"
			} else {
				var subject_name = info.Name
				var realmodules []proto.Module
				//exam
				ex := dbase.Exam{}
				row := db.QueryRow("select * from Exam where subject_id = (SELECT SubjectID from Subject where Description = $1)", subject_name)
				err = row.Scan(&ex.Id, &ex.Questions, &ex.MaxScore, &ex.MinScore, &ex.Date, &ex.SubjectID)
				if err != nil {
					println(err.Error())
				}
				fmt.Println("exam date:", ex.Date)
				//modules
				mods := []dbase.Module{}
				res, err := db.Query("select * from Modules where subject_id = (SELECT SubjectID from Subject where Description = $1);", subject_name)
				if err != nil {
					panic(err)
				}
				k := 0
				for res.Next() {
					mod := dbase.Module{}
					err = res.Scan(&mod.Id, &mod.SubjectID, &mod.Name, &mod.MaxScore, &mod.MinScore)
					mod.Db = db
					if err != nil {
						fmt.Println(err)
						continue
					}
					k += 1
					mods = append(mods, mod)
					fmt.Println("module: ", mod.Name, " module num: ", k)
					//labs
					labs := []proto.Lab{}
					res2, err := db.Query("SELECT * FROM labview where module_id = $1", mod.Id)
					if err != nil {
						panic(err)
					}
					for res2.Next() {
						l := proto.Lab{}
						err = res2.Scan(&l.Num, &l.Name, &l.Text, &l.Max, &l.Min, &l.Date, &l.Deadline, &l.Module_id, &l.Instance, &l.Recieved, &l.Comment, &l.Bonus)
						if err != nil {
							fmt.Println(err)
							continue
						}
						labs = append(labs, l)
						fmt.Println(labs)
					}
					fmt.Println("\tstudent labs for module :")
					for _, l := range labs {
						// fmt.Println(l.Get_id())
						fmt.Println("\t", l.Num, l.Name, l.Recieved)
					}
					//rk s
					rks := []proto.RK{}
					res3, err := db.Query("SELECT * FROM rkview where module_id = $1", mod.Id)
					if err != nil {
						panic(err)
					}
					for res3.Next() {
						r := proto.RK{}
						err = res3.Scan(&r.Num, &r.Date, &r.Max, &r.Min, &r.Module_id, &r.Instance, &r.Variant, &r.Recieved, &r.Comment)
						if err != nil {
							fmt.Println(err)
							continue
						}
						rks = append(rks, r)
						fmt.Println(rks)
					}
					fmt.Println("\tstudent rks for module :")
					for _, r := range rks {
						// fmt.Println(l.Get_id())
						fmt.Println("\t", r.Num, r.Date, r.Recieved)
					}
					//lects
					lects := []proto.Attend{}
					res4, err := db.Query("SELECT * FROM lectview where module_id = $1", mod.Id)
					if err != nil {
						panic(err)
					}
					for res4.Next() {
						l := proto.Attend{}
						err = res4.Scan(&l.Num, &l.Theme, &l.Date, &l.Module_id, &l.Attendance, &l.Bonus)
						if err != nil {
							fmt.Println(err)
							continue
						}
						lects = append(lects, l)
						fmt.Println(lects)
					}
					fmt.Println("\tstudent lects for module :")
					for _, l := range lects {
						// fmt.Println(l.Get_id())
						fmt.Println("\t", l.Num, l.Date, l.Theme)
					}
					//sems
					sems := []proto.Attend{}
					res5, err := db.Query("SELECT * FROM semview where module_id = $1", mod.Id)
					if err != nil {
						panic(err)
					}
					for res5.Next() {
						l := proto.Attend{}
						err = res5.Scan(&l.Num, &l.Theme, &l.Date, &l.Module_id, &l.Attendance, &l.Bonus)
						if err != nil {
							fmt.Println(err)
							continue
						}
						sems = append(sems, l)
						fmt.Println(sems)
					}
					fmt.Println("\tstudent sems for module :")
					for _, l := range sems {
						// fmt.Println(l.Get_id())
						fmt.Println("\t", l.Num, l.Date, l.Theme)
					}
					realmodules = append(realmodules, proto.Module{
						ModNumber: k,
						Labs:      labs,
						Rks:       rks,
						Lects:     lects,
						Sems:      sems,
					})
				}
				fmt.Println("\nmodules for databases subject:")
				for _, mod := range mods {
					fmt.Println(mod.Get_id())
					fmt.Println(mod.Get_name(), mod.Get_max_score(), mod.Get_min_score())
				}
				client.respond("ok", &proto.SubjectInfo{
					Name: info.Name,
					Mods: realmodules,
					Exam: proto.Exam{Date: ex.Date, Min: ex.MinScore, Max: ex.MaxScore},
				})
			}
		}
		if errorMsg != "" {
			log.Println("client: fetching databases subject info failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	case "project":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var info proto.CP
			if err := json.Unmarshal(*req.Data, &info); err != nil {
				errorMsg = "malformed data field"
			} else {
				var studentmail = info.Subject
				println("email: ", studentmail)
				l := proto.CP{}
				row := db.QueryRow("SELECT * FROM cpview where student_id = (select StudentID from Student where Email = $1)", studentmail)
				err = row.Scan(&l.Subject, &l.Description, &l.StartDate, &l.Deadline, &l.Student_id,
					&l.ProjAssignment, &l.TitleOfProject, &l.RecievedScore, &l.DateOfPassing)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(l.Subject)
				fmt.Println("\tstudent project", l.DateOfPassing)
				client.respond("ok", &l)
			}
		}
		if errorMsg != "" {
			log.Println("client: fetching studentproject info failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}
	default:
		log.Println("client: unknown command")
		client.respond("failed", "unknown command")
	}
}

// respond - вспомогательный метод для передачи ответа с указанным статусом
// и данными. Данные могут быть пустыми (data == nil).
func (client *StudentClient) respond(status string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	client.enc.Encode(&proto.Response{Status: status, Data: &raw})
}

func openConnection() {
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:6000", "specify ip address and port")
	flag.Parse()
	// Разбор адреса, строковое представление которого находится в переменной addrStr.
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		log.Fatal("address resolution failed", "address", addrStr)
	} else {
		log.Println("resolved TCP address", "address", addr.String())

		// Инициация слушания сети на заданном адресе.
		if listener, err := net.ListenTCP("tcp", addr); err != nil {
			log.Fatal("listening failed", "reason", err)
		} else {
			// Цикл приёма входящих соединений.
			for {
				if conn, err := listener.AcceptTCP(); err != nil {
					log.Fatal("cannot accept connection", "reason", err)
					break
				} else {
					log.Println("accepted connection", "address", conn.RemoteAddr().String())

					// Запуск go-программы для обслуживания клиентов.
					go NewStudentClient(conn).serve()
					log.Println("here")
				}
			}
		}
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.PostgresConfig.Host,
		configs.PostgresConfig.Port,
		configs.PostgresConfig.User,
		configs.PostgresConfig.Password,
		configs.PostgresConfig.DBname,
	)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected!\n\n")
	fill.ConnectAndFill()
	openConnection()
}
