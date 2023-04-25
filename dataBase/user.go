package dataBase 

const (
	teacher = iota
	student
)

type User struct{
    id []uint8
    login string
    passw string
    userRights int
}