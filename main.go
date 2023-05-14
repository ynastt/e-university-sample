package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "html/template"
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

func index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "index", nil)
}

func new_teacher(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("templates/new_teacher.html", "templates/header.html", "templates/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_teacher", nil)
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

    res, err := db.Query(fmt.Sprintf("INSERT INTO Teacher(TeacherID, user_id, TeacherName, Surname,Patronymic, Email) Values (gen_random_uuid(),(SELECT UserID From Users Where Login = '%s') ,'%s', '%s','%s',  '%s')",login, name, surname, patronymic, email ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func handleFunc() {
    http.HandleFunc("/", index)
    http.HandleFunc("/new_teacher", new_teacher)
    http.HandleFunc("/save_teacher", save_teacher)
    http.ListenAndServe(":8080", nil)
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
