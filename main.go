package main

import (
    "database/sql"
    "fmt"
  
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "ub7u3nAntu"
    dbname   = "Euniversity"
)

type StudentGroup struct{
    groupId []uint8
    groupName string
    yearOfAdm int
    course int
    amount int
}

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
    groups := []StudentGroup{}
    for res.Next(){
        g := StudentGroup{}
        err := res.Scan(&g.groupId, &g.groupName, &g.yearOfAdm, &g.course, &g.amount)
        if err != nil{
            fmt.Println(err)
            continue
        }
        groups = append(groups, g)
    }
    fmt.Println("\nGroup Names:")
    for _, g := range groups{
        fmt.Println(g.groupName)
    }
}