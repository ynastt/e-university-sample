package fill

import (
	"database/sql"
	"e-university-sample/configs"
    "encoding/json"
	dbase "e-university-sample/dataBase"
	"fmt"
)

var (
    db *sql.DB
    err error  
)

func ConnectAndFill() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
        "password=%s dbname=%s sslmode=disable",
        configs.PostgresConfig.Host,
		configs.PostgresConfig.Port,
		configs.PostgresConfig.User,
		configs.PostgresConfig.Password,
		configs.PostgresConfig.DBname,
	)
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Successfully connected!\n\n")

	// var query string 
	// query = "insert into StudentGroup(GroupID, GroupName, YearOfAdmission, Course, AmountOfStudents) Values (gen_random_uuid(),'ИУ9-61Б', 2020, 3, 28)"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
	res, err := db.Query("select * from StudentGroup")
    if err != nil {
        panic(err)
    }
    studentGroups := []dbase.StudentGroup{}
    for res.Next(){
        g := dbase.StudentGroup{}
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
        fmt.Println(g.Get_name(), g.Get_id())
    }   

	_, err = db.Query("DELETE FROM Users where Login = 'yarv';")
    if err != nil {
        panic(err)
    }

	_, err = db.Query("INSERT INTO Users(UserID, Login, Passw, UsersRights) Values (gen_random_uuid(),'yarv', '12', 2);")
    if err != nil {
        panic(err)
    }

	res, err = db.Query("select * from Users")
    if err != nil {
        panic(err)
    }

    users := []dbase.User{}
    for res.Next(){
        g := dbase.User{}
        err := res.Scan(&g.Id, &g.Login, &g.Passw, &g.UserRights)
        g.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        users = append(users, g)
    }
    fmt.Println("\nusers names:")
    for _, g := range users{
        fmt.Println(g.Get_login(), g.Get_passw(), g.Get_userRights(), g.Get_id())
    } 
    _, err = db.Query("DELETE FROM Student where Email = 'yarovnast@mail.ru';")
    if err != nil {
        panic(err)
    }
	query := "INSERT INTO Student(StudentID, StudentName, Surname, Patronymic, Email, Phone, YearOfAdmission, PassedCourses, NumInGroup, user_id, group_id) Values" + 
			 "(gen_random_uuid(), 'Анастасия', 'Яровикова', 'Сергеевна', 'yarovnast@mail.ru', 89569045346, 2020, 3, 29, (SELECT UserID from Users where Login = 'yarv'), (SELECT GroupID from StudentGroup where GroupName = 'ИУ9-61Б'));"
	_, err = db.Query(query)
    if err != nil {
        panic(err)
    }
	stud := dbase.Student{}
	row := db.QueryRow("select * from Student where Email = $1", "yarovnast@mail.ru")
	err = row.Scan(&stud.Id, &stud.Name, &stud.Surname, &stud.Patronymic, &stud.Email, &stud.Phone, &stud.Year, &stud.Courses, &stud.Number, &stud.Userid, &stud.Groupid)
	if err != nil{
		println(err.Error())
	}
	fmt.Println("student:", stud.Id)
    fmt.Println( stud.Name, stud.Surname, stud.Patronymic)
    fmt.Println(stud.Userid) 
    fmt.Println(stud.Groupid)

    // res, err = db.Query("DELETE * FROM Subject")
    // if err != nil {
    //     panic(err)
    // }

    // query := "INSERT INTO Subject(SubjectID, Description, SubjectProgram, NumberOfHours, NumberOfCredits) Values" + 
	// 		 "(gen_random_uuid(), 'БАЗЫ ДАННЫХ', '2 модуля 2 рк', 240, 100)," +
    //          "(gen_random_uuid(), 'РПиРП', '2 модуля 2 рк', 180, 100);"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }

    subjs := []dbase.Subject{}
    res, err = db.Query("select * from Subject")
    if err != nil {
        panic(err)
    }
    for res.Next(){
        subj := dbase.Subject{}
        var prog string
        err = res.Scan(&subj.Id, &subj.Description, &prog, &subj.Hours, &subj.Credits)
        subj.Program = json.RawMessage([]byte(prog))
        subj.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        subjs = append(subjs, subj)
    }
    fmt.Println("\nsubjects names:")
    for _, subj := range subjs{
        fmt.Println(subj.Get_id(), subj.Get_description())
    }
}