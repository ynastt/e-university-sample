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

	_ "github.com/lib/pq"
	fill "e-university-sample/fill"
)

var (
    db *sql.DB
    err error  
	conn *net.TCPConn  
)

func connectPostgres() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
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
	conn   *net.TCPConn  // Объект TCP-соединения
	enc    *json.Encoder // Объект для кодирования и отправки сообщений
}

// NewClient - конструктор клиента, принимает в качестве параметра
// объект TCP-соединения.
func NewStudentClient(conn *net.TCPConn) *StudentClient {
	return &StudentClient{
		conn:   conn,
		enc:    json.NewEncoder(conn),
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
				if err != nil{
					errorMsg = err.Error()
				}
				fmt.Println(user.Id, user.Login, user.Passw, user.UserRights)
  
				if user.UserRights != 2 {
					errorMsg = "it is not student"
				} else {
					if info.Username == user.Login && info.Password == user.Passw {
						info.Exists = true
					}
					log.Println("client: user exists")
					client.respond("ok", &proto.LoginInfo{
						Username: user.Get_login(), 
						Password: user.Get_passw(),
						Exists: true,
					})
				}
			}
		}
		if errorMsg != ""  {
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
				if err != nil{
					errorMsg = err.Error()
				}
				fmt.Println("user id:", user.Id)
				stud := dbase.Student{}
				g := dbase.StudentGroup{}
				row = db.QueryRow("select * from Student where user_id = $1", string(user.Id))
				err = row.Scan(&stud.Id, &stud.Name, &stud.Surname, &stud.Patronymic, &stud.Email, &stud.Phone, &stud.Year, &stud.Courses, &stud.Number, &stud.Userid, &stud.Groupid)
				if err != nil{
					errorMsg = err.Error()
				}
				println("student: ", stud.Name, "errr:", errorMsg)
				println("group id: ", stud.Groupid, "errr:", errorMsg)
				row = db.QueryRow("select GroupName from StudentGroup where GroupID = $1", stud.Groupid)
				err = row.Scan(&g.Name)
				if err != nil{
					errorMsg = err.Error()
				}
				log.Println("client: student is found")
				log.Println(stud.Name, stud.Surname, stud.Patronymic, g.Name)
				client.respond("ok", &proto.StudInfo{
					Name: stud.Get_name(), 
					Surname: stud.Get_surname(),
					Patronymic: stud.Get_patronymic(),
					Group: g.Name,
				})
			}
		}
		
		if errorMsg != ""  {
			log.Println("client: finding student failed", "reason", errorMsg)
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
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
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
