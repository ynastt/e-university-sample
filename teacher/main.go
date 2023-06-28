package main

import (
    "database/sql"
    "fmt"
    "time"
    "net/http"
    "html/template"
    "encoding/json"
    "log"
    "e-university-sample/dataBase"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    //password = "ub7u3nAntu"
    dbname   = "Euniversity"
)



func auth(w http.ResponseWriter, r *http.Request){
    fmt.Println("auth")
    // http.ServeFile(w, r, "teacher/templs/auth.html")
    t, err := template.ParseFiles("teacher/templs/auth.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "auth", nil)
}

func index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/index.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "index", nil)
}

func new_group(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_group", nil)
}

func new_bc(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_bc", nil)
}

func new_ex(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_ex", nil)
}

func new_lab(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_lab", nil)
}

func new_lect(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_lect", nil)
}

func new_module(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_module", nil)
}

func new_sem(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_sem", nil)
}

func new_teacher(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_teacher", nil)
}

func new_subject(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_subject", nil)
}

func new_student(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/new_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "new_student", nil)
}

func delete_student(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_student", nil)
}

func delete_bc(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_bc", nil)
}

func delete_ex(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_ex", nil)
}
func delete_group(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_group", nil)
}
func delete_lab(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_lab", nil)
}
func delete_lect(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_lect", nil)
}
func delete_module(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_module", nil)
}
func delete_sem(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_sem", nil)
}
func delete_subject(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_subject", nil)
}
func delete_teacher(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/delete_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "delete_teacher", nil)
}

func del_student(w http.ResponseWriter, r *http.Request){
    DEmail := r.FormValue("DEmail")
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

    res, err := db.Query(
        fmt.Sprintf("DELETE FROM Student Where Email = '%s'", DEmail))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_bc(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM BC Where Theme = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_ex(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Exam Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_group(w http.ResponseWriter, r *http.Request){
    GroupName := r.FormValue("GroupName")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM StudentGroup Where GroupName = '%s'", GroupName))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_lab(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    LabName := r.FormValue("LabName")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Lab Where LabName = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", LabName, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_lect(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Lecture Where Theme = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_module(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Modules Where ModuleID IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_sem(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Seminar Where Theme = '%s' AND module_id = (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id = (SELECT SubjectID From Subject Where Description = '%s'))", Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_subject(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Subject Where Description = '%s'", Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func del_teacher(w http.ResponseWriter, r *http.Request){
    Email := r.FormValue("Email")
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

    res, err := db.Query(
        fmt.Sprintf("Delete FROM Teacher Where Email = '%s'" , Email))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}


func update_student(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_student", nil)
}

func update_bc(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_bc", nil)
}
func update_ex(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_ex", nil)
}
func update_group(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_group", nil)
}
func update_lab(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_lab", nil)
}
func update_lect(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_lect", nil)
}
func update_module(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_module", nil)
}
func update_sem(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_sem", nil)
}
func update_subject(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_subject", nil)
}
func update_teacher(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/update_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "update_teacher", nil)
}

func upd_student(w http.ResponseWriter, r *http.Request){
    UEmail := r.FormValue("UEmail")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Student SET %s  = '%s' where Email = '%s'", Uchange, Uvalue, UEmail))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func upd_bc(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update BC SET %s  = '%s' where Theme = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Uchange, Uvalue,Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_ex(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Exam SET %s  = '%s' Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Uchange, Uvalue, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_group(w http.ResponseWriter, r *http.Request){
    GroupName := r.FormValue("GroupName")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update StudentGroup SET %s  = '%s' where GroupName = '%s'", Uchange, Uvalue, GroupName))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_lab(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    LabName := r.FormValue("LabName")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Lab SET %s  = '%s' where LabName = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Uchange, Uvalue, LabName, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_lect(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Lecture SET %s  = '%s' where Theme = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Uchange, Uvalue, Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_module(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Modules SET %s  = '%s' Where ModuleID IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Uchange, Uvalue, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_sem(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    ModuleName := r.FormValue("ModuleName")
    Theme := r.FormValue("Theme")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Seminar SET %s  = '%s' where Theme = '%s' AND module_id IN (SELECT ModuleID From Modules Where ModuleName = '%s' AND subject_id IN (SELECT SubjectID From Subject Where Description = '%s'))", Uchange, Uvalue, Theme, ModuleName, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}
func upd_subject(w http.ResponseWriter, r *http.Request){
    Description := r.FormValue("Description")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Subject SET %s  = '%s' Where SubjectID IN (SELECT SubjectID From Subject Where Description = '%s')", Uchange, Uvalue, Description))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func upd_teacher(w http.ResponseWriter, r *http.Request){
    UEmail := r.FormValue("Email")
    Uchange := r.FormValue("change")
    Uvalue := r.FormValue("value")
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

    res, err := db.Query(
        fmt.Sprintf("update Teacher SET %s  = '%s' where Email = '%s'", Uchange, Uvalue, UEmail))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}


func show_list_subjects(w http.ResponseWriter, r *http.Request){
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
    rows, err := db.Query("select * from Subject")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    subjs := []dataBase.Subject{}
     
    for rows.Next(){
        p := dataBase.Subject{}
        err := rows.Scan(&p.Id, &p.Description, &p.Program, &p.Hours, &p.Credits)
        //fmt.Println(p.Program)
        if err != nil{
            fmt.Println(err)
            continue
        }
        subjs = append(subjs, p)
    }
 
    tmpl, _ := template.ParseFiles("teacher/templs/subjects.html", "teacher/templs/header.html", "teacher/templs/footer.html")
    tmpl.ExecuteTemplate(w, "subjects", subjs)
}

type Mods struct {
    Labs []dataBase.Lab
    Sems []dataBase.Seminar
    Lects []dataBase.Lecture
    Bc []dataBase.Bc
}

type ModInfo struct {
    Ms Mods
    Module dataBase.Module
}

type Sbj struct {
    Exams dataBase.Exam
    Modules []ModInfo
    Description string
    IsLab bool
    IsSem bool
}

func show_subject(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    Descr := vars["descr"]
    var subj = Sbj{}
    t, err := template.ParseFiles("teacher/templs/show_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }
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
    rows, err := db.Query(fmt.Sprintf("select * from Exam Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Descr))
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    var date time.Time;
    for rows.Next(){
        p := dataBase.Exam{}
        err := rows.Scan(&p.Id, &p.Questions, &p.MaxScore, &p.MinScore, &date, &p.SubjectID)
        p.Date = date.String()[:11]
        fmt.Printf("%s %s %s %s\n", p.Questions, p.MaxScore, p.MinScore,p.Date)
        if err != nil{
            fmt.Println(err)
            continue
        }
        subj.Exams = p
    }
    subj.Description = Descr

    rows1, err1 := db.Query(fmt.Sprintf("select * from Modules Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Descr))
    if err1 != nil {
        log.Println(err1)
    }
    defer rows1.Close()
    subj.Modules = []ModInfo{}
    for rows1.Next(){
        p := dataBase.Module{}
        err1 := rows1.Scan(&p.Id, &p.SubjectID, &p.Name, &p.MaxScore, &p.MinScore)
        fmt.Printf("%s %d %d\n", p.Name, p.MaxScore, p.MinScore)
        if err1 != nil{
            fmt.Println(err1)
            continue
        }

        rowslab, errLab := db.Query(fmt.Sprintf("select * from Lab Where module_id = '%s'", p.Id))
        if errLab != nil {
            log.Println(errLab)
        }
        defer rowslab.Close()
        //var mms Mods
        lbs := []dataBase.Lab{}
        var datelab1 time.Time;
        var datelab2 time.Time;
        for rowslab.Next(){
            l := dataBase.Lab{}
            errlab := rowslab.Scan(&l.Id, &l.Name, &l.Text, &l.MaxScore, &l.MinScore, &datelab1, &datelab2, &l.ModuleID)
            fmt.Printf("%s %d %d\n", l.Name, l.MaxScore, l.MinScore)
            l.Date = datelab1.String()[:11]
            l.Deadline = datelab2.String()[:11]
            if errlab != nil{
                fmt.Println(errlab)
                continue
            }
            lbs = append(lbs, l)
        }

        rowssem, errsem := db.Query(fmt.Sprintf("select * from Seminar Where module_id = '%s'", p.Id))
        if errsem != nil {
            log.Println(errsem)
        }
        defer rowssem.Close()
        //var mms Mods
        sems := []dataBase.Seminar{}
        for rowssem.Next(){
            s := dataBase.Seminar{}
            errsem := rowssem.Scan(&s.Id, &s.Theme, &s.Text, &s.ModuleID)
            fmt.Printf("%s\n", s.Theme)
            if errsem != nil{
                fmt.Println(errsem)
                continue
            }
            sems = append(sems, s)
        }

        rowslect, errlect := db.Query(fmt.Sprintf("select * from Lecture Where module_id = '%s'", p.Id))
        if errlect != nil {
            log.Println(errlect)
        }
        defer rowslect.Close()
        //var mms Mods
        lects := []dataBase.Lecture{}
        for rowslect.Next(){
            lt := dataBase.Lecture{}
            errlect := rowslect.Scan(&lt.Id, &lt.Theme, &lt.Text, &lt.ModuleID)
            fmt.Printf("%s\n", lt.Theme)
            if errlect != nil{
                fmt.Println(errlect)
                continue
            }
            lects = append(lects, lt)
        }

        rowsbc, errbc := db.Query(fmt.Sprintf("select * from BC Where module_id = '%s'", p.Id))
        if errbc != nil {
            log.Println(errbc)
        }
        defer rowsbc.Close()
        //var mms Mods
        bcs := []dataBase.Bc{}
        for rowsbc.Next(){
            b := dataBase.Bc{}
            errbc := rowsbc.Scan(&b.Id, &b.Theme, &b.Questions, &b.MaxScore, &b.MinScore, &b.ModuleID)
            fmt.Printf("%s %d %d\n", b.Theme, b.MaxScore, b.MinScore)
            if errbc != nil{
                fmt.Println(errbc)
                continue
            }
            bcs = append(bcs, b)
        }
       

        var mm ModInfo
        mm.Module = p
        mm.Ms.Labs = lbs
        if len(lbs) == 0 {
            subj.IsLab = false
            //fmt.Println("aaaaaaaaaaaaaaaaaaa")
        } else {
            subj.IsLab =true
        }
        mm.Ms.Sems = sems
        if len(sems) == 0 {
            subj.IsSem = false
            //fmt.Println("jkdsxjdcfjkdsfjkdfjkkfjd") 
        } else {
            subj.IsSem =true
        }
        mm.Ms.Lects = lects
        mm.Ms.Bc = bcs
        subj.Modules = append(subj.Modules, mm)
    }
    
    t.ExecuteTemplate(w, "show_subject", subj)

}

type mmn struct {
	Students []stud
}

type stud struct {
	Name       string
	Surname    string
	Patronymic string
	Email      string
	Phone      string
	Courses    int
	Number     int
	Year       int
	UserLogin  string
}

func save_group(w http.ResponseWriter, r *http.Request){
    group_num := r.FormValue("group_num")
    course := r.FormValue("course")
    amount_of_students := r.FormValue("amount_of_students")
    year_of_admission := r.FormValue("year_of_admission")
    group := []byte(r.FormValue("group"))
    var app = mmn{}
	err1 := json.Unmarshal(group, &app)
	if err1 != nil {
		log.Fatal("Unmarshaling error")
	}



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
    res, err := db.Query(
        fmt.Sprintf("INSERT INTO StudentGroup(GroupID, GroupName, YearOfAdmission, Course , AmountOfStudents) Values (gen_random_uuid(),'%s', '%s','%s',  '%s')",
        group_num, course, amount_of_students, year_of_admission))
        
    if err != nil {
        panic(err)
    }

    defer res.Close()

    for _, aa := range app.Students {
        sname := aa.Name
	    ssurname := aa.Surname
	    spatronymic := aa.Patronymic 
	    semail := aa.Email
	    sphone := aa.Phone
	    scourses := aa.Courses
	    snumber := aa.Number
	    syear := aa.Year
	    suserLogin := aa.UserLogin
        res1, err := db.Query(
            fmt.Sprintf("INSERT INTO Student(StudentID, StudentName, Surname, Patronymic , Email, Phone, YearOfAdmission, PassedCourses, NumInGroup,user_id, group_id ) Values (gen_random_uuid(),'%s', '%s','%s',  '%s',  '%s',  '%d',  '%d',  '%d',(SELECT UserID From Users Where Login = '%s'), (SELECT GroupID From StudentGroup Where GroupName = '%s'))",
            sname, ssurname, spatronymic, semail, sphone, scourses, snumber, syear, suserLogin, group_num))
        
        if err != nil {
            panic(err)
        }
        defer res1.Close()
    }


    
    http.Redirect(w,r,"/", http.StatusSeeOther)
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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Teacher(TeacherID, user_id, TeacherName, Surname,Patronymic, Email) Values (gen_random_uuid(),(SELECT UserID From Users Where Login = '%s') ,'%s', '%s','%s',  '%s')",
            login, name, surname, patronymic, email ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_subject(w http.ResponseWriter, r *http.Request){
    description := r.FormValue("description")
    subjectProgram := r.FormValue("subjectProgram")
    numberOfHours := r.FormValue("numberOfHours")
    numberOfCredits := r.FormValue("numberOfCredits")
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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Subject(SubjectID, Description, SubjectProgram, NumberOfHours,NumberOfCredits) Values (gen_random_uuid(),'%s', '%s','%s',  '%s')",
        description, subjectProgram, numberOfHours, numberOfCredits))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_student(w http.ResponseWriter, r *http.Request){
    Studname := r.FormValue("Sname")
    Studsurname := r.FormValue("Ssurname")
    Studpatronymic := r.FormValue("Spatronymic")
    Studemail := r.FormValue("Semail")
    Studphone := r.FormValue("Sphone")
    Studseryear := r.FormValue("Syear")
    StudpassedCourses := r.FormValue("SpassedCourses")
    StudnumInGroup := r.FormValue("SnumInGroup")
    Studuser_login := r.FormValue("Suser_login")
    Studgroup_name := r.FormValue("Sgroup_name")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Student(StudentID, StudentName, Surname, Patronymic , Email, Phone, YearOfAdmission, PassedCourses, NumInGroup,user_id, group_id ) Values (gen_random_uuid(),'%s', '%s','%s',  '%s',  '%s',  '%s',  '%s',  '%s',(SELECT UserID From Users Where Login = '%s'), (SELECT GroupID From StudentGroup Where GroupName = '%s'))",
        Studname, Studsurname, Studpatronymic, Studemail, Studphone, Studseryear, StudpassedCourses, StudnumInGroup, Studuser_login, Studgroup_name))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_bc(w http.ResponseWriter, r *http.Request){
    module_name := r.FormValue("module_name")
    Theme := r.FormValue("Theme")
    Questions := r.FormValue("Questions")
    MaxScore := r.FormValue("MaxScore")
    MinScore := r.FormValue("MinScore")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO BC(BCID, Theme, Questions, MaxScore, MinScore, module_id) Values (gen_random_uuid(), '%s', '%s','%s',  '%s',(SELECT ModuleID From Modules Where ModuleName = '%s'))",
        Theme, Questions, MaxScore, MinScore, module_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_lab(w http.ResponseWriter, r *http.Request){
    module_name := r.FormValue("module_name")
    LabName := r.FormValue("LabName")
    LabText := r.FormValue("LabText")
    MaxScore := r.FormValue("MaxScore")
    MinScore := r.FormValue("MinScore")
    LabDate := r.FormValue("LabDate")
    Deadline := r.FormValue("Deadline")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Lab(LabID, LabName, LabText, MaxScore, MinScore,LabDate,Deadline, module_id) Values (gen_random_uuid(), '%s', '%s','%s','%s','%s',  '%s',(SELECT ModuleID From Modules Where ModuleName = '%s'))",
        LabName, LabText, MaxScore, MinScore, LabDate,Deadline,module_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_ex(w http.ResponseWriter, r *http.Request){
    subject_name := r.FormValue("subject_name")
    Questions := r.FormValue("Questions")
    MaxScore := r.FormValue("MaxScore")
    MinScore := r.FormValue("MinScore")
    ExamDate := r.FormValue("ExamDate")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Exam(ExamID, Questions, MaxScore, MinScore,ExamDate, subject_id) Values (gen_random_uuid(), '%s', '%s' , '%s' , '%s' ,(SELECT SubjectID From Subject Where Description = '%s'))",
         Questions, MaxScore, MinScore, ExamDate,subject_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_module(w http.ResponseWriter, r *http.Request){
    subject_name := r.FormValue("subject_name")
    ModuleName := r.FormValue("ModuleName")
    MaxScore := r.FormValue("MaxScore")
    MinScore := r.FormValue("MinScore")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Modules(ModuleID, ModuleName, MaxScore, MinScore, subject_id) Values (gen_random_uuid(), '%s', '%s','%s',(SELECT SubjectID From Subject Where Description = '%s'))",
        ModuleName, MaxScore, MinScore,subject_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_sem(w http.ResponseWriter, r *http.Request){
    module_name := r.FormValue("module_name")
    Theme := r.FormValue("Theme")
    SeminarText := r.FormValue("SeminarText")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Seminar(SeminarID, Theme, SeminarText,module_id) Values (gen_random_uuid(), '%s', '%s',(SELECT ModuleID From Modules Where ModuleName = '%s'))",
        Theme, SeminarText, module_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func save_lect(w http.ResponseWriter, r *http.Request){
    module_name := r.FormValue("module_name")
    Theme := r.FormValue("Theme")
    LectureText := r.FormValue("LectureText")

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

    res, err := db.Query(
        fmt.Sprintf("INSERT INTO Lecture(LectureID, Theme, LectureText,module_id) Values (gen_random_uuid(), '%s', '%s',(SELECT ModuleID From Modules Where ModuleName = '%s'))",
        Theme, LectureText, module_name ))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    http.Redirect(w,r,"/", http.StatusSeeOther)
}

func handleFunc() {
    router := mux.NewRouter()
    router.HandleFunc("/", show_list_subjects)
    router.HandleFunc("/main", index)
    router.HandleFunc("/new_teacher", new_teacher)
    router.HandleFunc("/save_teacher", save_teacher)
    router.HandleFunc("/new_group", new_group)
    router.HandleFunc("/save_group", save_group)
    router.HandleFunc("/new_subject", new_subject)
    router.HandleFunc("/save_subject", save_subject)
    router.HandleFunc("/new_student", new_student)
    router.HandleFunc("/save_student", save_student)
    router.HandleFunc("/del_student", del_student)
    router.HandleFunc("/subjects/", show_list_subjects)
    router.HandleFunc("/subjects/{descr}", show_subject)
    router.HandleFunc("/new_bc", new_bc)
    router.HandleFunc("/save_bc", save_bc)
    router.HandleFunc("/new_lab", new_lab)
    router.HandleFunc("/save_lab", save_lab)
    router.HandleFunc("/new_ex", new_ex)
    router.HandleFunc("/save_ex", save_ex)
    router.HandleFunc("/new_module", new_module)
    router.HandleFunc("/save_module", save_module)
    router.HandleFunc("/new_sem", new_sem)
    router.HandleFunc("/save_sem", save_sem)
    router.HandleFunc("/new_lect", new_lect)
    router.HandleFunc("/save_lect", save_lect)
    router.HandleFunc("/delete_student", delete_student)
    router.HandleFunc("/delete_bc", delete_bc)
    router.HandleFunc("/delete_ex", delete_ex)
    router.HandleFunc("/delete_group", delete_group)
    router.HandleFunc("/delete_lab", delete_lab)
    router.HandleFunc("/delete_lect", delete_lect)
    router.HandleFunc("/delete_module", delete_module)
    router.HandleFunc("/delete_sem", delete_sem)
    router.HandleFunc("/delete_subject", delete_subject)
    router.HandleFunc("/delete_teacher", delete_teacher)
    router.HandleFunc("/del_bc", del_bc)
    router.HandleFunc("/del_ex", del_ex)
    router.HandleFunc("/del_group", del_group)
    router.HandleFunc("/del_lab", del_lab)
    router.HandleFunc("/del_lect", del_lect)
    router.HandleFunc("/del_module", del_module)
    router.HandleFunc("/del_sem", del_sem)
    router.HandleFunc("/del_subject", del_subject)
    router.HandleFunc("/del_teacher", del_teacher)
	router.HandleFunc("/update_student", update_student)
	router.HandleFunc("/update_bc", update_bc)
    router.HandleFunc("/update_ex", update_ex)
    router.HandleFunc("/update_group", update_group)
    router.HandleFunc("/update_lab", update_lab)
    router.HandleFunc("/update_lect", update_lect)
    router.HandleFunc("/update_module", update_module)
    router.HandleFunc("/update_sem", update_sem)
    router.HandleFunc("/update_subject", update_subject)
    router.HandleFunc("/update_teacher", update_teacher)
    router.HandleFunc("/upd_student", upd_student)
	router.HandleFunc("/upd_bc", upd_bc)
    router.HandleFunc("/upd_ex", upd_ex)
    router.HandleFunc("/upd_group", upd_group)
    router.HandleFunc("/upd_lab", upd_lab)
    router.HandleFunc("/upd_lect", upd_lect)
    router.HandleFunc("/upd_module", upd_module)
    router.HandleFunc("/upd_sem", upd_sem)
    router.HandleFunc("/upd_subject", upd_subject)
    router.HandleFunc("/upd_teacher", upd_teacher)
    http.Handle("/", router)
    fmt.Println("Server is listening...")
    http.ListenAndServe("localhost:8080", nil)
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
