package main 

import (
	"encoding/json"
    "fmt"
	"flag"
    "log"
    "net/http"
	"net"
    "html/template"
    _ "github.com/lib/pq"
	proto "e-university-sample/proto"
)

type UserDetails struct {
    Username string
    Password string
    Succes bool
}

var err error
var conn *net.TCPConn

func auth(w http.ResponseWriter, r *http.Request){
	log.Printf("Authentification")
    t, err := template.ParseFiles("templs/auth.html", "templs/header.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "auth", nil)
}

func student_main(w http.ResponseWriter, r *http.Request){ 
	log.Printf("Student index page")
    t, err := template.ParseFiles("templs/student_main.html", "templs/header.html")
    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
    t.ExecuteTemplate(w, "student_main", nil)
}

func check_student(w http.ResponseWriter, r *http.Request) {
	err = r.ParseForm()
    if err != nil {
        log.Println(err)
    }
	var data UserDetails
	if r.FormValue("username") != "" && r.FormValue("userpassword") != "" {
		data = UserDetails{
        	Username:   r.FormValue("username"),
        	Password:   r.FormValue("userpassword"),
			Succes: 	false, 
    	}
	}
	// запрос на сервер
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	// данные для передачи на сервер
	var info proto.LoginInfo
	info.Username = data.Username
	info.Password = data.Password
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

    print(data.Username, data.Password)
	print(data.Succes)
	
	// data.Succes = true
	if data.Succes {
		log.Printf("Successful authentification")
		http.Redirect(w, r, "/student", 301)
	} else {
		log.Printf("Authentification error: %s is missing from the database", data.Username)
		http.Redirect(w, r, "/logout", 301)
	}
}

func handleFunc() {
    http.HandleFunc("/logout", auth)
	http.HandleFunc("/login", check_student)
    http.HandleFunc("/student", student_main)
	fmt.Println("Server is listening...")
    log.Fatal(http.ListenAndServe(":80", nil))
}

// send_request - вспомогательная функция для передачи запроса с указанной командой
// и данными. Данные могут быть пустыми (nil).
func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{Command: command, Data: &raw})
}

func connectToServer() *net.TCPConn {
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:6000", "specify ip address and port")
	flag.Parse()

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
		handleFunc()
	}
}