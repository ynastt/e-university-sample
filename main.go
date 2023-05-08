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
    password = "ub7u3nAntu"
    // password = "postgres"
    //password = "ub7u3nAntu"
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


    res, err := db.Query("select * from StudentGroup")
    if err != nil {
        panic(err)
    }
    defer res.Close()
    studentGroups := []dataBase.StudentGroup{}
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
    }
}
