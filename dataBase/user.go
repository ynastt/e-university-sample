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
    _, err := u.Db.Exec("update Users set Login = $1 where UserID = $2", u.Login, u.Id)
    if err != nil {
        panic(err)
    }
}
func (u *User) Set_passw(passw1 string) {
    u.Passw = passw1
    _, err := u.Db.Exec("update Users set Passw = $1 where UserID = $2", u.Passw, u.Id)
    if err != nil {
        panic(err)
    }
}
func (u *User) Set_userRights(userRights1 bool) {
     u.UserRights = userRights1
     _, err := u.Db.Exec("update Users set UsersRights = $1 where UserID = $2", u.UserRights, u.Id)
     if err != nil {
        panic(err)
    }
}
