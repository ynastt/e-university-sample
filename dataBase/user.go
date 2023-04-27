package dataBase 
    
import (
    "database/sql"
)
const (
	teacher = iota
	student
)

type User struct{
    Id []uint8
    Login string
    Passw string
    UserRights bool
    Db *sql.DB
}

//func (u user) init() {}

func (u User) Get_id() []uint8 { return u.Id }
func (u User) Get_login() string { return u.Login }
func (u User) Get_passw() string { return u.Passw }
func (u User) Get_userRights() bool{ return u.UserRights }

func (u *User) Set_login(log1 string) {
    u.Login = log1
    _, err := u.Db.Exec("update Users (Login) values ($1)", u.Login)
    if err != nil {
        panic(err)
    }
}
func (u *User) Set_passw(passw1 string) {
    u.Passw = passw1
    _, err := u.Db.Exec("update Users (Passw) values ($1)", u.Passw)
    if err != nil {
        panic(err)
    }
}
func (u *User) Set_userRights(userRights1 bool) {
     u.UserRights = userRights1
     _, err := u.Db.Exec("update Users (UsersRights) values ($1)", u.UserRights)
     if err != nil {
        panic(err)
    }
}
