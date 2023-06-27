package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "html/template"
    "encoding/json"
    "log"
    //"e-university-sample/dataBase"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    //password = "ub7u3nAntu"
    dbname   = "Euniversity"
)

func auth(w http.ResponseWriter, r *http.Request){
    fmt.Println("auth")
    // http.ServeFile(w, r, "templates/auth.html")
    t, err := template.ParseFiles("templates/auth.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "auth", nil)
}

func index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "index", nil)
}

func new_group(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/new_group.html", "templates/header.html", "templates/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_group", nil)
}

func new_teacher(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/new_teacher.html", "templates/header.html", "templates/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_teacher", nil)
}

type mmn struct {
	Students []stud
}

type stud struct {
	Name       string
	Surname    string
	Patronymic string
	Email      string
	Phone      string
	Courses    int
	Number     int
	Year       int
	UserLogin  string
}

func save_group(w http.ResponseWriter, r *http.Request){
    group_num := r.FormValue("group_num")
    course := r.FormValue("course")
    amount_of_students := r.FormValue("amount_of_students")
    year_of_admission := r.FormValue("year_of_admission")
    group := []byte(r.FormValue("group"))
    var app = mmn{}
	err1 := json.Unmarshal(group, &app)
	if err1 != nil {
		log.Fatal("Unmarshaling error")
	}



    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Successfully connected!\n\n")
    res, err := db.Query(
        fmt.Sprintf("INSERT INTO StudentGroup(GroupID, GroupName, YearOfAdmission, Course , AmountOfStudents) Values (gen_random_uuid(),'%s', '%s','%s',  '%s')",
        group_num, course, amount_of_students, year_of_admission))
        
    if err != nil {
        panic(err)
    }

    defer res.Close()

    for _, aa := range app.Students {
        sname := aa.Name
	    ssurname := aa.Surname
	    spatronymic := aa.Patronymic 
	    semail := aa.Email
	    sphone := aa.Phone
	    scourses := aa.Courses
	    snumber := aa.Number
	    syear := aa.Year
	    suserLogin := aa.UserLogin
        res1, err := db.Query(
            fmt.Sprintf("INSERT INTO Student(StudentID, StudentName, Surname, Patronymic , Email, Phone, YearOfAdmission, PassedCourses, NumInGroup,user_id, group_id ) Values (gen_random_uuid(),'%s', '%s','%s',  '%s',  '%s',  '%d',  '%d',  '%d',(SELECT UserID From Users Where Login = '%s'), (SELECT GroupID From StudentGroup Where GroupName = '%s'))",
            sname, ssurname, spatronymic, semail, sphone, scourses, snumber, syear, suserLogin, group_num))
        
        if err != nil {
            panic(err)
        }
        defer res1.Close()
    }


    
    http.Redirect(w,r,"/", http.StatusSeeOther)
}


func save_teacher(w http.ResponseWriter, r *http.Request){
    name := r.FormValue("name")
    surname := r.FormValue("surname")
    patronymic := r.FormValue("patronymic")
    email := r.FormValue("email")
    login := r.FormValue("login")

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Successfully connected!\n\n")

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Teacher(TeacherID, user_id, TeacherName, Surname,Patronymic, Email) Values (gen_random_uuid(),(SELECT UserID From Users Where Login = '%s') ,'%s', '%s','%s',  '%s')",
            login, name, surname, patronymic, email ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func handleFunc() {
    http.HandleFunc("/", auth)
    http.HandleFunc("/main", index)
    http.HandleFunc("/new_teacher", new_teacher)
    http.HandleFunc("/save_teacher", save_teacher)
    http.HandleFunc("/new_group", new_group)
    http.HandleFunc("/save_group", save_group)
    fmt.Println("Server is listening...")
    http.ListenAndServe("localhost:8080", nil)
}

func main() {
    handleFunc()


    /*studentGroups := []dataBase.StudentGroup{}
    for res.Next(){
        g := dataBase.StudentGroup{}
        err := res.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
        g.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        studentGroups = append(studentGroups, g)
    }
    fmt.Println("\nstudentGroups names:")
    for _, g := range studentGroups{
        fmt.Println(g.Get_name())
    }*/
}
