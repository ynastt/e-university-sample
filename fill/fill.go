package fill

import (
	"database/sql"
	"e-university-sample/configs"
	dbase "e-university-sample/dataBase"
	"fmt"
)

var (
    db *sql.DB
    err error  
    query string
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
// student groups
	// var query string 
	// query = "insert into StudentGroup(GroupID, GroupName, YearOfAdmission, Course, AmountOfStudents) Values (gen_random_uuid(),'ИУ9-61Б', 2020, 3, 28)"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
	res, err := db.Query("select * from StudentGroup;")
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
// users
	// _, err = db.Query("DELETE FROM Users where Login = 'yarv';")
    // if err != nil {
    //     panic(err)
    // }

	// _, err = db.Query("INSERT INTO Users(UserID, Login, Passw, UsersRights) Values (gen_random_uuid(),'yarv', '12', 2);")
    // if err != nil {
    //     panic(err)
    // }

	res, err = db.Query("select * from Users;")
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
// students    
    // _, err = db.Query("DELETE FROM Student where Email = 'yarovnast@mail.ru';")
    // if err != nil {
    //     panic(err)
    // }
	// query := "INSERT INTO Student(StudentID, StudentName, Surname, Patronymic, Email, Phone, YearOfAdmission, PassedCourses, NumInGroup, user_id, group_id) Values" + 
	// 		 "(gen_random_uuid(), 'Анастасия', 'Яровикова', 'Сергеевна', 'yarovnast@mail.ru', 89569045346, 2020, 3, 29, (SELECT UserID from Users where Login = 'yarv'), (SELECT GroupID from StudentGroup where GroupName = 'ИУ9-61Б'));"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
	stud := dbase.Student{}
	row := db.QueryRow("select * from Student where Email = $1;", "yarovnast@mail.ru")
	err = row.Scan(&stud.Id, &stud.Name, &stud.Surname, &stud.Patronymic, &stud.Email, &stud.Phone, &stud.Year, &stud.Courses, &stud.Number, &stud.Userid, &stud.Groupid)
	if err != nil{
		println(err.Error())
	}
	fmt.Println("student:", stud.Id)
    fmt.Println( stud.Name, stud.Surname, stud.Patronymic)
    fmt.Println(stud.Userid) 
    fmt.Println(stud.Groupid)

// subjects
    // query = "INSERT INTO Subject(SubjectID, Description, SubjectProgram, NumberOfHours, NumberOfCredits) Values" + 
	// 		 "(gen_random_uuid(), 'БАЗЫ ДАННЫХ', '2 модуля 2 рк', 240, 100)," +
    //          "(gen_random_uuid(), 'РПиРП', '2 модуля 2 рк', 180, 100);"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }

    subjs := []dbase.Subject{}
    res, err = db.Query("select * from Subject;")
    if err != nil {
        panic(err)
    }
    for res.Next(){
        subj := dbase.Subject{}
        err = res.Scan(&subj.Id, &subj.Description, &subj.Program, &subj.Hours, &subj.Credits)
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
// exam
    // query = "delete from exam where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ');"
    // _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
    // query = "INSERT INTO Exam(ExamID, Questions, MaxScore, MinScore, ExamDate, subject_id) Values" + 
	// 		 "(gen_random_uuid(), 'Список вопросов будет опубликован на зачетной неделе', 30, 18, '2023-01-15', (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ'));"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
    ex := dbase.Exam{}
	row = db.QueryRow("select * from Exam where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ')")
	err = row.Scan(&ex.Id, &ex.Questions, &ex.MaxScore, &ex.MinScore, &ex.Date, &ex.SubjectID)
	if err != nil{
		println(err.Error())
	}
    fmt.Println("\nexam:")
    fmt.Println(ex.Get_id())
    fmt.Println(ex.Get_date(), ex.Get_max_score(), ex.Get_min_score())

// modules
    // res, err = db.Query("DELETE FROM Modules where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ')")
    // if err != nil {
    //     panic(err)
    // }
    // query = "INSERT INTO Modules(ModuleID, subject_id, ModuleName, MaxScore, MinScore) Values" + 
	// 		 "(gen_random_uuid(), (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ'), 'Проектирование баз данных', 25, 20)," +
    //          "(gen_random_uuid(), (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ'), 'Реализация баз данных', 40, 25);"
	// _, err = db.Query(query)
    // if err != nil {
    //     panic(err)
    // }
    mods := []dbase.Module{}
    res, err = db.Query("select * from Modules where subject_id = (SELECT SubjectID from Subject where Description = 'БАЗЫ ДАННЫХ')")
    if err != nil {
        panic(err)
    }
    mod_ids := make([][]uint8, 0)
    for res.Next(){
        mod := dbase.Module{}
        err = res.Scan(&mod.Id, &mod.SubjectID, &mod.Name, &mod.MaxScore, &mod.MinScore)
        mod.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        mod_ids = append(mod_ids, mod.Id)
        mods = append(mods, mod)
    }
    fmt.Println("\nmodules for databases subject:")
    for _, mod := range mods {
        fmt.Println(mod.Get_id())
        fmt.Println(mod.Get_name(), mod.Get_max_score(), mod.Get_min_score())
    }
// labs
    // res, err = db.Query("delete from Lab;")
    // if err != nil {
    //     panic(err)
    // }
    // for pgAdmin:
//     delete from Lab;

// INSERT INTO Lab(LabID, LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id) Values 
// 	 (gen_random_uuid(), 1, 'ER', '\nЛабораторная работа №1\nМоделирование данных с использованием модели сущность связь\n\t1. Выбрать простейшую предметную область, соответствующую 4-5 сущностям.\n\t2. Сформировать требования к предметной области.\n\t3. Создать модель\"сущность-связь\" для предметной области с обоснованием выбора кардинальных чисел связей.\n', 
//       8, 5, '2022-09-09', '2022-09-16', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
//       (gen_random_uuid(), 2, 'SOM', '\nЛабораторная работа №2\nМоделирование данных с использованием модели семантических объектов\n\t1.Создать модель семантических объектов для предметной области, выбранной в лабораторной работе №1\n\t2.Обосновать выбор кардинальных чисел атрибутов и типов объектов\n', 
//       5, 4, '2022-09-16', '2022-09-23', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
//       (gen_random_uuid(), 3, 'ER->R', '\nЛабораторная работа №3\nПреобразование модели \"сущность-связь\" в реляционную модель\n\t1.Преобразовать модель \"сущность-связь\", созданную в лабораторной работе №1, в реляционную модель согласно процедуре преобразования\n\t2.Обосновать выбор типов данных, ключей, правил обеспечения ограничей минимальной кардинальности.\n', 
//       7, 5, '2022-09-23', '2022-09-30', '84d9bb22-939a-469c-b27e-01ee3a43e63b'),
//       (gen_random_uuid(), 4, 'SOM->R', '\nЛабораторная работа №4\nПреобразование модели семантических объектов в реляционную модель\n',  
//       5, 4, '2022-09-30', '2022-10-07', '84d9bb22-939a-469c-b27e-01ee3a43e63b');
// INSERT INTO Lab(LabID, LabNumber, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id) Values 	  
// 	(gen_random_uuid(), 5, 'DB files', '\nЛабораторная работа №5\nОперации с базой данных, файлами, схемами\n', 3, 2, '2022-10-07', '2022-10-14', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 6, 'Constraints', '\nЛабораторная работа №6\nКлючи, ограничения, значения по умолчанию\n', 3, 2, '2022-10-14', '2022-10-21', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 7, 'Views', '\nЛабораторная работа №7\nПредставления и индексы\n\t1. Создать представление на основе одной из таблиц задания 6.\n\t2. Создать представлени на основе полей обеих связанныха таблиц задания 6.\n\t3. Создать индекс для одной из таблиц задания 6, включив в него дополнительные неключевые поля\n\t4. Создать индексированное представление.\n',3, 2, '2022-10-21', '2022-10-28', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 8, 'SP/Cursor/F', '\nЛабораторная работа №8\nПреобразование модели семантических объектов в реляционную модель\n',4, 3, '2022-10-28', '2022-11-11', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 9, 'Trigger', '\nЛабораторная работа №9\nПреобразование модели семантических объектов в реляционную модель\n', 4, 3, '2022-10-28', '2022-11-11', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 10, 'Transactions', 'Лабораторная работа №10\nПреобразование модели семантических объектов в реляционную модель\n', 4, 2, '2022-11-11', '2022-11-18', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 11, 'DB App', '\nЛабораторная работа №11\nПреобразование модели семантических объектов в реляционную модель\n',  7, 5, '2022-11-18', '2022-11-25', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 12, 'ADO.NET', '\nЛабораторная работа №12\nПреобразование модели семантических объектов в реляционную модель\n',4, 2, '2022-11-25', '2022-12-02', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 13, 'horizontal', '\nЛабораторная работа №13\nПреобразование модели семантических объектов в реляционную модель\n', 3, 2, '2022-12-02', '2022-12-09', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 14, 'vertical', '\nЛабораторная работа №14\nПреобразование модели семантических объектов в реляционную модель\n',3, 2, '2022-12-09', '2022-12-16', '04b03c25-59be-42f7-bb1f-4c452a36aa6e'),
//     (gen_random_uuid(), 15, 'linked', '\nЛабораторная работа №15\nПреобразование модели семантических объектов в реляционную модель\n',3, 2, '2022-12-16', '2022-12-23', '04b03c25-59be-42f7-bb1f-4c452a36aa6e');
// select * from Lab;
    lab_ids := make([][]uint8, 0)
    labs := []dbase.Lab{}
    res, err = db.Query("select * from Lab where module_id = $1", mod_ids[0])
    if err != nil {
        panic(err)
    }
    for res.Next(){
        l := dbase.Lab{}
        err = res.Scan(&l.Id, &l.Number, &l.Name, &l.Text, &l.MaxScore, &l.MinScore, &l.Date, &l.Deadline, &l.ModuleID)
        l.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        lab_ids = append(lab_ids, l.Id)
        labs = append(labs, l)
    }
    fmt.Println("\nlabs for module 1:")
    for _, l := range labs {
        // fmt.Println(l.Get_id())
        fmt.Println(l.Get_number(), l.Get_name(), l.Get_max_score(), l.Get_min_score()/*, l.Get_date(), l.Get_dealine(), l.Get_text()*/)
    }
    labs = []dbase.Lab{}
    res, err = db.Query("select * from Lab where module_id = $1", mod_ids[1])
    if err != nil {
        panic(err)
    }
    for res.Next(){
        l := dbase.Lab{}
        err = res.Scan(&l.Id, &l.Number, &l.Name, &l.Text, &l.MaxScore, &l.MinScore, &l.Date, &l.Deadline, &l.ModuleID)
        l.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        lab_ids = append(lab_ids, l.Id)
        labs = append(labs, l)
    }
    fmt.Println("\nlabs for module 2:")
    for _, l := range labs {
        // fmt.Println(l.Get_id())
        fmt.Println(l.Get_number(), l.Get_name(), l.Get_max_score(), l.Get_min_score()/*, l.Get_date(), l.Get_dealine(), l.Get_text()*/)
    }
// lab instance
    res, err = db.Query("DELETE FROM LabInstance;")
    if err != nil {
        panic(err)
    }
    query = "INSERT INTO LabInstance(student_id, lab_id, DateOdPassing, NumOfInstance, RecievedScore, Variant, Remarks, BonusScore) Values" + 
			 "((select STudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT LabID from Lab where Labname = 'ER'), '2022-09-16', 1, 8, NULL, 'БД отлично спроектирована!', 1)," +
             "((select STudentID from Student where Email = 'yarovnast@mail.ru'), (SELECT LabID from Lab where Labname = 'Views'), '2022-09-23', 1, 1, NULL, 'ЛР сдана позже дедлайна', 0);"
	_, err = db.Query(query)
    if err != nil {
        panic(err)
    }
    labinsts := []dbase.LabInstance{}
    res, err = db.Query("select * from LabInstance where student_id = (select STudentID from Student where Email = 'yarovnast@mail.ru')")
    if err != nil {
        panic(err)
    }
    for res.Next(){
        l := dbase.LabInstance{}
        err = res.Scan(&l.StudentId, &l.LabID, &l.Date, &l.NumOfInstance, &l.Score, &l.Variant, &l.Remarks, &l.BonusScore)
        l.Db = db
        if err != nil{
            fmt.Println(err)
            continue
        }
        labinsts = append(labinsts, l)
    }
    fmt.Println("\nlabs of student Яровикова А.С:")
    for _, l := range labinsts {
        fmt.Println(l.Get_id())
        fmt.Println(l.Get_date(), l.Get_score(), l.Get_bonus_score(), l.Get_remarks(), l.Get_variant())
    }

// bc
    // query = "INSERT INTO Lab(LabID, LabName, LabText, MaxScore, MinScore, LabDate, Deadline, module_id) Values" + 
	// 		 "(gen_random_uuid(), 'ER', '\nЛабораторная работа №1\nМоделирование данных с использованием модели сущность связь\n\t1. Выбрать простейшую предметную область, соответствующую 4-5 сущностям." + 
    //          "\n\t2. Сформировать требования к предметной области.\n\t3. Создать модель\"сущность-связь\" для предметной области с обоснованием выбора кардинальных чисел связей.\n'," + 
    //          "8, 5, '2022-09-09', '2022-09-16', (SELECT ModuleID from Modules where ModuleID = $1))," +
    //          "(gen_random_uuid(), 'linked', '\nЛабораторная работа №15\nПреобразование модели семантических объектов в реляционную модель\n'," +  
    //          "3, 2, '2022-12-16', '2022-12-23', (SELECT ModuleID from Modules where ModuleID = $2));"
	// _, err = db.Query(query, mod_ids[0], mod_ids[1])
    // if err != nil {
    //     panic(err)
    // }
    // labs := []dbase.Lab{}
    // res, err = db.Query("select * from Lab where module_id = $1", mod_ids[0])
    // if err != nil {
    //     panic(err)
    // }
    // for res.Next(){
    //     l := dbase.Lab{}
    //     err = res.Scan(&l.Id, &l.Name, &l.Text, &l.MaxScore, &l.MinScore, &l.Date, &l.Deadline, &l.ModuleID)
    //     l.Db = db
    //     if err != nil{
    //         fmt.Println(err)
    //         continue
    //     }
    //     labs = append(labs, l)
    // }
    // fmt.Println("\nlabs for module 1:")
    // for _, l := range labs {
    //     // fmt.Println(l.Get_id())
    //     fmt.Println(l.Get_name(), l.Get_max_score(), l.Get_min_score()/*, l.Get_date(), l.Get_dealine(), l.Get_text()*/)
    // }
    // labs = []dbase.Lab{}
    // res, err = db.Query("select * from Lab where module_id = $1", mod_ids[1])
    // if err != nil {
    //     panic(err)
    // }
    // for res.Next(){
    //     l := dbase.Lab{}
    //     err = res.Scan(&l.Id, &l.Name, &l.Text, &l.MaxScore, &l.MinScore, &l.Date, &l.Deadline, &l.ModuleID)
    //     l.Db = db
    //     if err != nil{
    //         fmt.Println(err)
    //         continue
    //     }
    //     labs = append(labs, l)
    // }
    // fmt.Println("\nrkss for module 2:")
    // for _, l := range labs {
    //     // fmt.Println(l.Get_id())
    //     fmt.Println(l.Get_name(), l.Get_max_score(), l.Get_min_score()/*, l.Get_date(), l.Get_dealine(), l.Get_text()*/)
    // }
// lect

// sems
}