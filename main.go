package main

import (
    "database/sql"
    "fmt"
    "e-university-sample/dataBase"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    //password = "ub7u3nAntu"
    password = "postgres"
    dbname   = "Euniversity"
)

func main() {
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


    res, err := db.Query("select * from Student")
    if err != nil {
        panic(err)
    }
    defer res.Close()
    students := []dataBase.Student{}
    //groups := []StudentGroup{}
    for res.Next(){
        g := dataBase.Student{}
        err := res.Scan(&g.Id, &g.Name, &g.Surname, &g.Patronymic, &g.Email, &g.Phone, &g.Courses, &g.Number, &g.Year, &g.Userid, &g.Groupid)
        g.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        students = append(students, g)
    }
    fmt.Println("\nStudents names:")
    for _, g := range students{
        fmt.Println(g.Get_name())
    }
}