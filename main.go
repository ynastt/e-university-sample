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


    res, err := db.Query("select * from Users")
    if err != nil {
        panic(err)
    }
    defer res.Close()
    users := []dataBase.User{}
    //groups := []StudentGroup{}
    for res.Next(){
        g := dataBase.User{}
        err := res.Scan(&g.Id, &g.Login, &g.Passw, &g.UserRights)
        g.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        users = append(users, g)
    }
    fmt.Println("\nUsers logins:")
    for _, g := range users{
        fmt.Println(g.Get_login())
    }
}