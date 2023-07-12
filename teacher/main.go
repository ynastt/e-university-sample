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
    "github.com/google/uuid"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    //password = "ub7u3nAntu"
    dbname   = "Euniversity"
)

type Session struct {
	Token  uuid.UUID
	UserId []uint8
}

var sessionTable []Session

func index(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/index.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "index", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_group(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_group", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func logIn(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("teacher/templs/logIn.html", "teacher/templs/header.html", "teacher/templs/footer.html")

    if err != nil {
        fmt.Fprintf(w, err.Error())
    }

    t.ExecuteTemplate(w, "logIn", nil)
}

func new_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_coursePr.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_coursePr", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_bc", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_ex", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_lab", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_lect", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_module", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_sem", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_teacher", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_subject", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_student", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func new_connTGS(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/new_connTGS.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "new_connTGS", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func delete_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_student", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func delete_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_bc", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func delete_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_ex", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_group(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_group", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func delete_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_coursePr.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_coursePr", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_lab", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_lect", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_module", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_sem", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_subject", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func delete_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/delete_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "delete_teacher", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
                fmt.Sprintf("Delete FROM CourseProject Where Description = '%s'", Description))
            if err != nil {
                panic(err)
            }
            defer res.Close()
            http.Redirect(w,r,"/", http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_group(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func del_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}


func update_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_student.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_student", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func update_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_bc.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_bc", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_ex.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_ex", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_group(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_group.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_group", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_coursePr.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_coursePr", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_lab.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_lab", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_lect.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_lect", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_module.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_module", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_sem.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_sem", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_subject.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_subject", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func update_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            t, err := template.ParseFiles("teacher/templs/update_teacher.html", "teacher/templs/header.html", "teacher/templs/footer.html")

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            t.ExecuteTemplate(w, "update_teacher", nil)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func upd_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func upd_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_group(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func upd_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
                fmt.Sprintf("update CourseProject SET %s  = '%s' where Description = '%s'", Uchange, Uvalue, Description))
            if err != nil {
                panic(err)
            }
            defer res.Close()
            http.Redirect(w,r,"/", http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}
func upd_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func upd_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

type GrSb struct {
	Subj  dataBase.Subject
	Grs   []dataBase.StudentGroup
}

func show_list_groups_with_subj(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            rows, err := db.Query( fmt.Sprintf("Select * from Subject where SubjectId in (select Distinct subject_id from TeacherSubjectGroup Where teacher_id in (Select TeacherID From Teacher Where user_id = '%s'))", session.UserId))
            if err != nil {
                log.Println(err)
            }
            defer rows.Close()
            
            gs := []GrSb{}
            //subjs := []dataBase.Subject{}
     
            for rows.Next(){
                g := GrSb{}
                p := dataBase.Subject{}
                err := rows.Scan(&p.Id, &p.Description, &p.Program, &p.Hours, &p.Credits)
                if err != nil{
                    fmt.Println(err)
                    continue
                }
                g.Subj = p
                fmt.Println(p.Description)
                rowsg, errg := db.Query( fmt.Sprintf("select * from StudentGroup where GroupId in (select group_id from TeacherSubjectGroup Where subject_id = '%s' and teacher_id in (Select TeacherID From Teacher Where user_id = '%s'))", p.Id, session.UserId))
                if errg != nil {
                    log.Println(errg)
                }
                defer rowsg.Close()
                //subjs := []dataBase.Subject{}
        
                for rowsg.Next(){
                    gr := dataBase.StudentGroup{}
                    err := rowsg.Scan(&gr.Id, &gr.Name, &gr.YearOfAdm, &gr.Course, &gr.Amount)
                    if err != nil{
                        fmt.Println(err)
                        continue
                    }
                    fmt.Println(gr.Name)
                    g.Grs = append(g.Grs, gr)
                }
                gs = append(gs, g)
            }
        
            tmpl, _ := template.ParseFiles("teacher/templs/show_groups_and_subjs.html", "teacher/templs/header.html", "teacher/templs/footer.html")
            tmpl.ExecuteTemplate(w, "show_groups_and_subjs", gs)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}


func show_list_subjects(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func show_list_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            rows, err := db.Query("select * from CourseProject")
            if err != nil {
                log.Println(err)
            }
            defer rows.Close()
            subjs := []dataBase.CourseProject{}
     
            for rows.Next(){
                p := dataBase.CourseProject{}
                err := rows.Scan(&p.Id, &p.Subject, &p.Description, &p.Hours, &p.StartDate, &p.Deadline)
                p.StartDate = p.StartDate[:10]
                p.Deadline = p.Deadline[:10]
                fmt.Println(p.Description)
                if err != nil{
                    fmt.Println(err)
                    continue
                }
                subjs = append(subjs, p)
            }
        
            tmpl, _ := template.ParseFiles("teacher/templs/show_coursePrs.html", "teacher/templs/header.html", "teacher/templs/footer.html")
            tmpl.ExecuteTemplate(w, "show_coursePrs", subjs)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

type BCInstance2 struct {
    StudentId uuid.UUID
	BCID uuid.UUID
    Date string 
	NumOfInstance int
	Score int
    Variant int
    Remarks string
	Theme string
}
type LabInstance2 struct {
    StudentId uuid.UUID
	LabID uuid.UUID
    Date string 
	NumOfInstance int
	Score int
    Variant int
    Remarks string
    BonusScore int
	Name string
}

type LectureAttendance2 struct {
    StudentId uuid.UUID
	LectureID uuid.UUID
	Attendance bool
	BonusScore int
	Theme string
}

type SeminarAttendance2 struct {
    StudentId uuid.UUID
	SeminarID uuid.UUID
	Attendance bool
	BonusScore int
    Theme string
}


type Mods struct {
    Labs []dataBase.Lab
    LabInstances []LabInstance2
    Sems []dataBase.Seminar
    SeminarAttendances []SeminarAttendance2
    Lects []dataBase.Lecture
    LectureAttendances []LectureAttendance2
    Bc []dataBase.Bc
    BCInstances []BCInstance2
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



type Stud struct {
    Id uuid.UUID
    Name string
	Surname string
	Patronymic string
	Email string
    Phone string
    Courses int
    Number int
    Year int
    Userid []uint8
    Groupid []uint8
    Idstr string
}

type stHstr struct {
    Students []Stud
    Subjn string
    Groupn string
}

func show_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)

}

func show_students(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Grname := vars["grname"]
            Subjname := vars["subjname"]
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

            fmt.Printf("Successfully connected! students\n\n")
            
            rowsSt, errSt := db.Query(fmt.Sprintf("Select * from Student where group_id in (Select GroupId from StudentGroup where GroupName = '%s') order by NumInGroup", Grname))
            if errSt != nil {
                log.Println(errSt)
            }
            defer rowsSt.Close()
            stds := []Stud{}
     
            for rowsSt.Next(){
                s := Stud{}
                errSt := rowsSt.Scan(&s.Id, &s.Name, &s.Surname, &s.Patronymic, &s.Email, &s.Phone, &s.Year, &s.Courses, &s.Number, &s.Userid, &s.Groupid)
                s.Idstr = s.Id.String()
                fmt.Println(s.Name, s.Number)
                if errSt != nil{
                    fmt.Println(errSt)
                    continue
                }
                stds = append(stds, s)
            }
            shs := stHstr{}
            shs.Students = stds
            shs.Subjn = Subjname
            shs.Groupn = Grname
            tmpl, _ := template.ParseFiles("teacher/templs/show_students.html", "teacher/templs/header.html", "teacher/templs/footer.html")
            tmpl.ExecuteTemplate(w, "show_students", shs)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

type Idsubj struct {
    Examid uuid.UUID 
    Modules []Mods
    IsLab bool
    IsSem bool
    IsBc bool
    IsLect bool
}

type Sbj2 struct {
    Exams dataBase.Exam
    Examinst dataBase.ExamInstance
    EXid string
    Modules []ModInfo
    Description string
    IsLab bool
    IsSem bool
}

func show_marks(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Subjname := vars["subjname"]
            Studname := vars["studname"]
            var subj = Sbj2{}
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
            rows, err := db.Query(fmt.Sprintf("select * from Exam Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Subjname))
            if err != nil {
                log.Println(err)
            }
            defer rows.Close()
            var date time.Time;
            for rows.Next(){
                p := dataBase.Exam{}
                err := rows.Scan(&p.Id, &p.Questions, &p.MaxScore, &p.MinScore, &date, &p.SubjectID)
                p.Date = date.String()[:11]
                //fmt.Printf("%s %s %s %s\n", p.Questions, p.MaxScore, p.MinScore,p.Date)
                if err != nil{
                    fmt.Println(err)
                    continue
                }
                subj.Exams = p
            }
            subj.Description = Subjname

            
            
            rows1, err1 := db.Query(fmt.Sprintf("select * from Modules Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Subjname))
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
                lbis := []LabInstance2{}
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
                    rowslbi, errlbi := db.Query(fmt.Sprintf("select * from LabInstance Where student_id = '%s' and lab_id = '%s' ", Studname, l.Id))
                    if errlbi != nil {
                        log.Println(errlbi)
                    }
                    defer rowslbi.Close()
                    lbi := LabInstance2{}
                    for rowslbi.Next(){   
                        errlbi := rowslbi.Scan(&lbi.StudentId, &lbi.LabID, &date, &lbi.NumOfInstance, &lbi.Score, &lbi.Variant, &lbi.Remarks, &lbi.BonusScore)
                        lbi.Date = date.String()[:11]
                        lbi.Name = l.Name
                        fmt.Printf("%s %s %s\n", date, lbi.NumOfInstance, lbi.Score)
                        if errlbi != nil{
                            fmt.Println(errlbi)
                            continue
                        }
                    }
                    lbis = append(lbis, lbi)
                }

                

                rowssem, errsem := db.Query(fmt.Sprintf("select * from Seminar Where module_id = '%s'", p.Id))
                if errsem != nil {
                    log.Println(errsem)
                }
                defer rowssem.Close()
                //var mms Mods
                sems := []dataBase.Seminar{}
                semas := []SeminarAttendance2{}
                for rowssem.Next(){
                    s := dataBase.Seminar{}
                    errsem := rowssem.Scan(&s.Id, &s.Theme, &s.Text, &s.ModuleID)
                    fmt.Printf("%s\n", s.Theme)
                    if errsem != nil{
                        fmt.Println(errsem)
                        continue
                    }
                    sems = append(sems, s)

                    rowssema, errsema := db.Query(fmt.Sprintf("select * from SeminarAttendance Where student_id = '%s' and seminar_id = '%s' ", Studname, s.Id))
                    if errsema != nil {
                        log.Println(errsema)
                    }
                    defer rowssema.Close()
                    sema := SeminarAttendance2{}
                    for rowssema.Next(){   
                        errsema := rowssema.Scan(&sema.StudentId, &sema.SeminarID, &sema.Attendance, &sema.BonusScore)
                        sema.Theme = s.Theme
                        //fmt.Printf("%s %s %s\n", dlbi.NumOfInstance, lbi.Score)
                        if errsema != nil{
                            fmt.Println(errsema)
                            continue
                        }
                    }
                    semas = append(semas, sema)
                }

                rowslect, errlect := db.Query(fmt.Sprintf("select * from Lecture Where module_id = '%s'", p.Id))
                if errlect != nil {
                    log.Println(errlect)
                }
                defer rowslect.Close()
                //var mms Mods
                lects := []dataBase.Lecture{}
                lectas := []LectureAttendance2{}
                for rowslect.Next(){
                    lt := dataBase.Lecture{}
                    errlect := rowslect.Scan(&lt.Id, &lt.Theme, &lt.Text, &lt.ModuleID)
                    fmt.Printf("%s\n", lt.Theme)
                    if errlect != nil{
                        fmt.Println(errlect)
                        continue
                    }
                    lects = append(lects, lt)

                    rowslecta, errlecta := db.Query(fmt.Sprintf("select * from LectureAttendance Where student_id = '%s' and lecture_id = '%s' ", Studname, lt.Id))
                    if errlecta != nil {
                        log.Println(errlecta)
                    }
                    defer rowslecta.Close()
                    lecta := LectureAttendance2{}
                    for rowslecta.Next(){   
                        errlecta := rowslecta.Scan(&lecta.StudentId, &lecta.LectureID, &lecta.Attendance, &lecta.BonusScore)
                        lecta.Theme = lt.Theme
                        //lbi.Date = date.String()[:11]
                        //fmt.Printf("%s %s %s\n", date, lbi.NumOfInstance, lbi.Score)
                        if errlecta != nil{
                            fmt.Println(errlecta)
                            continue
                        }
                    }
                    lectas = append(lectas, lecta)
                }

                rowsbc, errbc := db.Query(fmt.Sprintf("select * from BC Where module_id = '%s'", p.Id))
                if errbc != nil {
                    log.Println(errbc)
                }
                defer rowsbc.Close()
                //var mms Mods
                bcs := []dataBase.Bc{}
                bcis := []BCInstance2{}
                for rowsbc.Next(){
                    b := dataBase.Bc{}
                    errbc := rowsbc.Scan(&b.Id, &b.Theme, &b.Questions, &b.MaxScore, &b.MinScore, &b.ModuleID)
                    fmt.Printf("%s %d %d\n", b.Theme, b.MaxScore, b.MinScore)
                    if errbc != nil{
                        fmt.Println(errbc)
                        continue
                    }
                    bcs = append(bcs, b)

                    rowsbci, errbci := db.Query(fmt.Sprintf("select * from BCInstance Where student_id = '%s' and bc_id = '%s' ", Studname, b.Id))
                    if errbci != nil {
                        log.Println(errbci)
                    }
                    defer rowsbci.Close()
                    bci := BCInstance2{}
                    for rowsbci.Next(){   
                        errbci := rowsbci.Scan(&bci.StudentId, &bci.BCID, &date, &bci.NumOfInstance, &bci.Score, &bci.Variant, &bci.Remarks)
                        bci.Date = date.String()[:11]
                        bci.Theme = b.Theme
                        fmt.Printf("%s %s %s\n", date, bci.NumOfInstance, bci.Score)
                        if errbci != nil{
                            fmt.Println(errbci)
                            continue
                        }
                    }
                    bcis = append(bcis, bci)
                }
            

                var mm ModInfo
                mm.Module = p
                mm.Ms.Labs = lbs
                mm.Ms.LabInstances = lbis
                if len(lbs) == 0 {
                    subj.IsLab = false
                    //fmt.Println("aaaaaaaaaaaaaaaaaaa")
                } else {
                    subj.IsLab =true
                }
                mm.Ms.Sems = sems
                mm.Ms.SeminarAttendances = semas
                if len(sems) == 0 {
                    subj.IsSem = false
                    //fmt.Println("jkdsxjdcfjkdsfjkdfjkkfjd") 
                } else {
                    subj.IsSem =true
                }
                mm.Ms.Lects = lects
                mm.Ms.LectureAttendances = lectas
                mm.Ms.Bc = bcs
                mm.Ms.BCInstances = bcis
                subj.Modules = append(subj.Modules, mm)
            }
            //вся инфа о предмете получена
            rowsei, errei := db.Query(fmt.Sprintf("select * from ExamInstance Where student_id = '%s' and exam_id = '%s' ", Studname, subj.Exams.Id))
            if errei != nil {
                log.Println(errei)
            }
            defer rowsei.Close()
            for rowsei.Next(){
                p := dataBase.ExamInstance{}
                errei := rowsei.Scan(&p.StudentId, &p.ExamID, &date, &p.NumOfInstance, &p.Score, &p.Ticket)
                p.Date = date.String()[:11]
                fmt.Printf("%s %s %s %s\n", date, p.NumOfInstance, p.Score,p.Ticket)
                if errei != nil{
                    fmt.Println(errei)
                    continue
                }
                subj.Examinst = p
            }
            tmpl, _ := template.ParseFiles("teacher/templs/show_marks.html", "teacher/templs/header.html", "teacher/templs/footer.html")
            tmpl.ExecuteTemplate(w, "show_marks", subj)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
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
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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

            fmt.Printf("Successfully connected to save group!\n\n")
            
            groupid := uuid.New()
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
                    fmt.Sprintf("INSERT INTO Student(StudentID, StudentName, Surname, Patronymic , Email, Phone, YearOfAdmission, PassedCourses, NumInGroup,user_id, group_id ) Values (gen_random_uuid(),'%s', '%s','%s',  '%s',  '%s',  '%d',  '%d',  '%d',(SELECT UserID From Users Where Login = '%s'), '%s')",
                    sname, ssurname, spatronymic, semail, sphone, syear, scourses, snumber, suserLogin, groupid))
                
                if err != nil {
                    panic(err)
                }
                defer res1.Close()
            }

            res, err := db.Query(
                fmt.Sprintf("INSERT INTO StudentGroup(GroupID, GroupName, YearOfAdmission, Course , AmountOfStudents) Values ('%s','%s', '%s','%s',  '%s')",
                groupid, group_num, course, amount_of_students, year_of_admission))
                
            if err != nil {
                panic(err)
            }

            defer res.Close()
            
            http.Redirect(w,r,"/", http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_coursePr(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            Subject := r.FormValue("Subject")
            Description := r.FormValue("Description")
            NumberOfHours := r.FormValue("NumberOfHours")
            StartDate := r.FormValue("StartDate")
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
                fmt.Sprintf("INSERT INTO CourseProject(ProjectID, Subject, Description, NumberOfHours , StartDate, Deadline) Values (gen_random_uuid(),'%s', '%s','%s',  '%s', '%s')",
                Subject, Description, NumberOfHours, StartDate,Deadline ))
                
            if err != nil {
                panic(err)
            }

            defer res.Close()
            http.Redirect(w,r,"/", http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func postLogIn(w http.ResponseWriter, r *http.Request) {
	Login := r.FormValue("Login")
    Password := r.FormValue("Password")
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

    user := dataBase.User{}
    res, err := db.Query(
        fmt.Sprintf("select * from Users where Login = '%s'",Login))
    if err != nil {
        panic(err)
    }
    defer res.Close()
    for res.Next(){
       
        err = res.Scan(&user.Id, &user.Login, &user.Passw, &user.UserRights)

        if err != nil{
            fmt.Println(err)
            continue
        }
    }
	
	
    fmt.Println(Login, Password, user.Login, user.Passw, user.UserRights)
    if user.UserRights == 2 {
		if Login == user.Login && Password == user.Passw {
			u := uuid.New()

			sessionTable = append(sessionTable, Session{u, user.Id})

			c := http.Cookie{
				Name:   "token",
				Value:  u.String(),
				MaxAge: 300,
			}

			http.SetCookie(w, &c)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
        http.Redirect(w,r,"/logIn", http.StatusSeeOther)
    }
    //defer row.Close()
    //
}


func save_teacher(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_connTGS(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            Description := r.FormValue("Description")
            Email := r.FormValue("Email")
            GroupName := r.FormValue("GroupName")
            TeacherRole := r.FormValue("TeacherRole")

            psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
            "password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)
            db, err := sql.Open("postgres", psqlInfo)
            
            if err != nil {
                panic(err)
            }
            defer db.Close()
            //db.SetConnMaxLifetime(0)
            err = db.Ping()
            if err != nil {
                panic(err)
            }

            fmt.Printf("Successfully connected!\n\n")

            var idsb = Idsubj{}

            rowsEx, errEx := db.Query(fmt.Sprintf("select ExamID, ExamDate  from Exam Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Description))
            if errEx != nil {
                log.Println(errEx)
            }
            defer rowsEx.Close()
            //var date time.Time;
            //fmt.Println("ksdjkfdjkfdjkfgjkfgjkfdjkf")
            var e uuid.UUID
            var date time.Time
            var edate string
            for rowsEx.Next(){
                errEx := rowsEx.Scan(&e, &date)
                edate = date.String()[:11]
                //fmt.Println("экзамен")
                //fmt.Printf("%s %s %s %s\n", p.Questions, p.MaxScore, p.MinScore,p.Date)
                if errEx != nil{
                    fmt.Println(errEx)
                    continue
                }
            }
            //subj.Description = Description
            rows1, err1 := db.Query(fmt.Sprintf("select ModuleID from Modules Where subject_id IN (SELECT SubjectID From Subject Where Description = '%s')", Description))
            if err1 != nil {
                log.Println(err1)
            }
            defer rows1.Close()
            //subj.Modules = []ModInfo{}
            for rows1.Next(){
                var mid []uint8
                err1 := rows1.Scan(&mid)
                //fmt.Printf("%s %d %d\n", p.Name, p.MaxScore, p.MinScore)
                if err1 != nil{
                    fmt.Println(err1)
                    continue
                }

                rowslab, errLab := db.Query(fmt.Sprintf("select LabID, Deadline from Lab Where module_id = '%s'", mid))
                if errLab != nil {
                    log.Println(errLab)
                }
                 
                //var mms Mods
                lbs := []dataBase.Lab{}
                var datelab2 time.Time;
                for rowslab.Next(){
                    l := dataBase.Lab{}
                    errlab := rowslab.Scan(&l.Id, &datelab2)
                    //fmt.Printf("%s %d %d\n", l.Name, l.MaxScore, l.MinScore)
                    l.Deadline = datelab2.String()[:11]
                    if errlab != nil{
                        fmt.Println(errlab)
                        continue
                    }
                    lbs = append(lbs, l)
                }

                rowssem, errsem := db.Query(fmt.Sprintf("select SeminarID from Seminar Where module_id = '%s'", mid))
                if errsem != nil {
                    log.Println(errsem)
                }
                
                //var mms Mods
                sems := []dataBase.Seminar{}
                for rowssem.Next(){
                    s := dataBase.Seminar{}
                    errsem := rowssem.Scan(&s.Id)
                    //fmt.Printf("%s\n", s.Theme)
                    if errsem != nil{
                        fmt.Println(errsem)
                        continue
                    }
                    sems = append(sems, s)
                }

                rowslect, errlect := db.Query(fmt.Sprintf("select LectureID from Lecture Where module_id = '%s'", mid))
                if errlect != nil {
                    log.Println(errlect)
                }
                
                //var mms Mods
                lects := []dataBase.Lecture{}
                for rowslect.Next(){
                    lt := dataBase.Lecture{}
                    errlect := rowslect.Scan(&lt.Id)
                    //fmt.Printf("%s\n", lt.Theme)
                    if errlect != nil{
                        fmt.Println(errlect)
                        continue
                    }
                    lects = append(lects, lt)
                }

                rowsbc, errbc := db.Query(fmt.Sprintf("select BCID from BC Where module_id = '%s'", mid))
                if errbc != nil {
                    log.Println(errbc)
                }
                
                //var mms Mods
                bcs := []dataBase.Bc{}
                for rowsbc.Next(){
                    b := dataBase.Bc{}
                    errbc := rowsbc.Scan(&b.Id)
                    //fmt.Printf("%s %d %d\n", b.Theme, b.MaxScore, b.MinScore)
                    if errbc != nil{
                        fmt.Println(errbc)
                        continue
                    }
                    bcs = append(bcs, b)
                }
            

                var mm Mods
                mm.Labs = lbs
                if len(lbs) == 0 {
                    idsb.IsLab = false
                    //fmt.Println("aaaaaaaaaaaaaaaaaaa")
                } else {
                    idsb.IsLab =true
                }
                mm.Sems = sems
                if len(sems) == 0 {
                    idsb.IsSem = false
                    //fmt.Println("jkdsxjdcfjkdsfjkdfjkkfjd") 
                } else {
                    idsb.IsSem =true
                }
                mm.Lects = lects
                mm.Bc = bcs
                idsb.Modules = append(idsb.Modules, mm)
                defer rowslab.Close()
                defer rowssem.Close()
                defer rowslect.Close()
                defer rowsbc.Close()
            }
            //получили всю информацию о предмете

            rowsSt, errSt := db.Query(fmt.Sprintf("select StudentID, StudentName from Student Where group_id IN (SELECT GroupID From StudentGroup Where GroupName = '%s')", GroupName))
            if errSt != nil {
                log.Println(errSt)
            }
            defer rowsSt.Close()
            st := dataBase.Student{}
            //var date time.Time;
            for rowsSt.Next(){
                st = dataBase.Student{}
                errSt := rowsSt.Scan(&st.Id, &st.Name)
                fmt.Println(st.Name,st.Id, "aboba")
                if errSt != nil{
                    fmt.Println(errSt)
                    continue
                }
                resExm, errExm := db.Query(
                    fmt.Sprintf("INSERT INTO ExamInstance(student_id, exam_id, DateOdPassing, NumOfInstance, RecievedScore, TicketNumber) Values ('%s', '%s' , '%s' ,  1, 0, 0)",
                    st.Id, e, edate))
                if errExm != nil {
                    panic(err)
                }
                defer resExm.Close()

                for _, mod := range idsb.Modules {
                    for _, lec := range mod.Lects {
                        //lecture attendance
                        fmt.Println(st.Name, "lecture")
                        reslc, errlc := db.Query(
                            fmt.Sprintf("INSERT INTO LectureAttendance(student_id, lecture_id, WasAttended, BonusScore) Values ('%s', '%s' ,  False, 0)",
                            st.Id, lec.Id ))
                        if errlc != nil {
                            panic(errlc)
                        }
                        reslc.Close()
                        
                    }
                    if (idsb.IsSem) {
                        for _, sem := range mod.Sems {
                            //seminar attendance
                            fmt.Println(st.Name, "seminar")
                            fmt.Println(st.Id, sem.Id)
                            ressm, errsm := db.Query(
                                fmt.Sprintf("INSERT INTO SeminarAttendance(student_id, seminar_id, WasAttended, BonusScore) Values ('%s', '%s' ,  False, 0)",
                                st.Id, sem.Id))
                            if errsm != nil {
                                panic(errsm)
                            }
                            ressm.Close()
                        }
                    }
                    if (idsb.IsLab) {
                        for _, lab := range mod.Labs {
                            //Lab instance
                            fmt.Println(st.Name, "lab")
                            reslb, errlb := db.Query(
                                fmt.Sprintf("INSERT INTO LabInstance(student_id, lab_id, DateOdPassing, NumOfInstance, RecievedScore, Variant, Remarks, BonusScore) Values ('%s', '%s' , '%s', 1, 0 , 0, '', 0)",
                                st.Id, lab.Id, lab.Deadline))
                            if errlb != nil {
                                panic(errlb)
                            }
                            reslb.Close()
                        }
                    }

                    for _, bc := range mod.Bc {
                        //BC instance
                        fmt.Println(st.Name, "bc")
                        resbc, errbc := db.Query(
                            fmt.Sprintf("INSERT INTO BCInstance(student_id, bc_id, DateOdPassing, NumOfInstance, RecievedScore, Variant, Remarks) Values ('%s', '%s' , '2000-01-01', 1, 0 , 0, '')",
                            st.Id, bc.Id))
                        if errbc != nil {
                            panic(errbc)
                        }
                        resbc.Close()
                    }
                }

            }
            res, err := db.Query(
                fmt.Sprintf("INSERT INTO TeacherSubjectGroup(teacher_id, subject_id, group_id, TeacherRole) Values ((SELECT TeacherID From Teacher Where Email = '%s'), (SELECT SubjectID From Subject Where Description = '%s') , (SELECT GroupID From StudentGroup Where GroupName = '%s'),  '%s')",
                Email, Description, GroupName, TeacherRole ))
            if err != nil {
                panic(err)
            }
            defer res.Close()
            http.Redirect(w,r,"/", http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}


func save_subject(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_student(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_bc(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_lab(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_ex(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_exinst(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Studid := vars["stid"]
            Exid := vars["exid"]
            Date := r.FormValue("date")
            Number := r.FormValue("number")
            Score := r.FormValue("score")
            Ticket := r.FormValue("ticket")

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
                fmt.Sprintf("update ExamInstance SET DateOdPassing  = '%s', NumOfInstance  = '%s', RecievedScore = '%s', TicketNumber = '%s' where student_id = '%s' AND exam_id = '%s'",
                Date, Number, Score, Ticket, Studid, Exid))
            if err != nil {
                panic(err)
            }
            defer res.Close()

            rowsgr, errgr := db.Query(fmt.Sprintf("select * from StudentGroup Where GroupID IN (Select group_id from Student where StudentID = '%s')", Studid))
            if errgr != nil {
                log.Println(errgr)
            }
            defer rowsgr.Close()
            g := dataBase.StudentGroup{}
            for rowsgr.Next(){
                errgr := rowsgr.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
                if errgr != nil{
                    fmt.Println(errgr)
                    continue
                }
            }

            rowssubj, errsubj := db.Query(fmt.Sprintf("select * from Subject Where SubjectID IN (Select subject_id from Exam where ExamID = '%s')", Exid))
            if errsubj != nil {
                log.Println(errsubj)
            }
            defer rowssubj.Close()
            sub := dataBase.Subject{}
            for rowssubj.Next(){
                errsubj := rowssubj.Scan(&sub.Id, &sub.Description, &sub.Program, &sub.Hours, &sub.Credits)
                if errsubj != nil{
                    fmt.Println(errsubj)
                    continue
                }
            }
            //{grname}+{subjname}/{subjname}+{studname}
            ss := "/show_groups_and_subjs/" + g.Name + "+" + sub.Description
            http.Redirect(w,r,ss, http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_lectatt(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Studid := vars["stid"]
            Ltid := vars["ltid"]
            Attendance := r.FormValue("Attendance")
            BonusScore := r.FormValue("BonusScore")

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
                fmt.Sprintf("update LectureAttendance SET WasAttended  = '%s', BonusScore  = '%s' where student_id = '%s' AND lecture_id = '%s'",
                Attendance, BonusScore, Studid, Ltid))
            if err != nil {
                panic(err)
            }
            defer res.Close()

            rowsgr, errgr := db.Query(fmt.Sprintf("select * from StudentGroup Where GroupID IN (Select group_id from Student where StudentID = '%s')", Studid))
            if errgr != nil {
                log.Println(errgr)
            }
            defer rowsgr.Close()
            g := dataBase.StudentGroup{}
            for rowsgr.Next(){
                errgr := rowsgr.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
                if errgr != nil{
                    fmt.Println(errgr)
                    continue
                }
            }

            rowssubj, errsubj := db.Query(fmt.Sprintf("select * from Subject Where SubjectID IN (Select subject_id from Modules where ModuleID in (Select module_id from Lecture where LectureID = '%s'))", Ltid))
            if errsubj != nil {
                log.Println(errsubj)
            }
            defer rowssubj.Close()
            sub := dataBase.Subject{}
            for rowssubj.Next(){
                errsubj := rowssubj.Scan(&sub.Id, &sub.Description, &sub.Program, &sub.Hours, &sub.Credits)
                if errsubj != nil{
                    fmt.Println(errsubj)
                    continue
                }
            }
            //{grname}+{subjname}/{subjname}+{studname}
            ss := "/show_groups_and_subjs/" + g.Name + "+" + sub.Description
            http.Redirect(w,r,ss, http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_sematt(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Studid := vars["stid"]
            Smid := vars["smid"]
            Attendance := r.FormValue("Attendance")
            BonusScore := r.FormValue("BonusScore")

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
                fmt.Sprintf("update SeminarAttendance SET WasAttended  = '%s', BonusScore  = '%s' where student_id = '%s' AND seminar_id = '%s'",
                Attendance, BonusScore, Studid, Smid))
            if err != nil {
                panic(err)
            }
            defer res.Close()

            rowsgr, errgr := db.Query(fmt.Sprintf("select * from StudentGroup Where GroupID IN (Select group_id from Student where StudentID = '%s')", Studid))
            if errgr != nil {
                log.Println(errgr)
            }
            defer rowsgr.Close()
            g := dataBase.StudentGroup{}
            for rowsgr.Next(){
                errgr := rowsgr.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
                if errgr != nil{
                    fmt.Println(errgr)
                    continue
                }
            }

            rowssubj, errsubj := db.Query(fmt.Sprintf("select * from Subject Where SubjectID IN (Select subject_id from Modules where ModuleID in (Select module_id from Seminar where SeminarID = '%s'))", Smid))
            if errsubj != nil {
                log.Println(errsubj)
            }
            defer rowssubj.Close()
            sub := dataBase.Subject{}
            for rowssubj.Next(){
                errsubj := rowssubj.Scan(&sub.Id, &sub.Description, &sub.Program, &sub.Hours, &sub.Credits)
                if errsubj != nil{
                    fmt.Println(errsubj)
                    continue
                }
            }
            //{grname}+{subjname}/{subjname}+{studname}
            ss := "/show_groups_and_subjs/" + g.Name + "+" + sub.Description
            http.Redirect(w,r,ss, http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_bcinst(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Studid := vars["stid"]
            Bcid := vars["bcid"]
            Date := r.FormValue("Date")
            NumOfInstance := r.FormValue("NumOfInstance")
            Score := r.FormValue("Score")
            Variant := r.FormValue("Variant")
            Remarks := r.FormValue("Remarks")

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
                fmt.Sprintf("update BCInstance SET DateOdPassing  = '%s', NumOfInstance  = '%s', RecievedScore  = '%s', Variant  = '%s', Remarks  = '%s' where student_id = '%s' AND bc_id = '%s'",
                Date, NumOfInstance, Score, Variant, Remarks, Studid, Bcid))
            if err != nil {
                panic(err)
            }
            defer res.Close()

            rowsgr, errgr := db.Query(fmt.Sprintf("select * from StudentGroup Where GroupID IN (Select group_id from Student where StudentID = '%s')", Studid))
            if errgr != nil {
                log.Println(errgr)
            }
            defer rowsgr.Close()
            g := dataBase.StudentGroup{}
            for rowsgr.Next(){
                errgr := rowsgr.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
                if errgr != nil{
                    fmt.Println(errgr)
                    continue
                }
            }

            rowssubj, errsubj := db.Query(fmt.Sprintf("select * from Subject Where SubjectID IN (Select subject_id from Modules where ModuleID in (Select module_id from BC where BCID = '%s'))", Bcid))
            if errsubj != nil {
                log.Println(errsubj)
            }
            defer rowssubj.Close()
            sub := dataBase.Subject{}
            for rowssubj.Next(){
                errsubj := rowssubj.Scan(&sub.Id, &sub.Description, &sub.Program, &sub.Hours, &sub.Credits)
                if errsubj != nil{
                    fmt.Println(errsubj)
                    continue
                }
            }
            //{grname}+{subjname}/{subjname}+{studname}
            ss := "/show_groups_and_subjs/" + g.Name + "+" + sub.Description
            http.Redirect(w,r,ss, http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_labi(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
            vars := mux.Vars(r)
            Studid := vars["stid"]
            Lbid := vars["lbid"]
            Date := r.FormValue("Date")
            NumOfInstance := r.FormValue("NumOfInstance")
            Score := r.FormValue("Score")
            Variant := r.FormValue("Variant")
            Remarks := r.FormValue("Remarks")
            BonusScore := r.FormValue("BonusScore")

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
                fmt.Sprintf("update LabInstance SET DateOdPassing  = '%s', NumOfInstance  = '%s', RecievedScore  = '%s', Variant  = '%s', Remarks  = '%s', BonusScore = '%s' where student_id = '%s' AND lab_id = '%s'",
                Date, NumOfInstance, Score, Variant, Remarks, BonusScore, Studid, Lbid))
            if err != nil {
                panic(err)
            }
            defer res.Close()

            rowsgr, errgr := db.Query(fmt.Sprintf("select * from StudentGroup Where GroupID IN (Select group_id from Student where StudentID = '%s')", Studid))
            if errgr != nil {
                log.Println(errgr)
            }
            defer rowsgr.Close()
            g := dataBase.StudentGroup{}
            for rowsgr.Next(){
                errgr := rowsgr.Scan(&g.Id, &g.Name, &g.YearOfAdm, &g.Course, &g.Amount)
                if errgr != nil{
                    fmt.Println(errgr)
                    continue
                }
            }

            rowssubj, errsubj := db.Query(fmt.Sprintf("select * from Subject Where SubjectID IN (Select subject_id from Modules where ModuleID in (Select module_id from Lab where LabID = '%s'))", Lbid))
            if errsubj != nil {
                log.Println(errsubj)
            }
            defer rowssubj.Close()
            sub := dataBase.Subject{}
            for rowssubj.Next(){
                errsubj := rowssubj.Scan(&sub.Id, &sub.Description, &sub.Program, &sub.Hours, &sub.Credits)
                if errsubj != nil{
                    fmt.Println(errsubj)
                    continue
                }
            }
            //{grname}+{subjname}/{subjname}+{studname}
            ss := "/show_groups_and_subjs/" + g.Name + "+" + sub.Description
            http.Redirect(w,r,ss, http.StatusSeeOther)
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}


func save_module(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_sem(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func save_lect(w http.ResponseWriter, r *http.Request){
    token, _ := r.Cookie("token")

	for _, session := range sessionTable {
		if session.Token.String() == token.Value {
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
            return
        }
    }
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func postLogOut(w http.ResponseWriter, r *http.Request) {
	tokenC, err := r.Cookie("token")
	if err != nil {
		panic(err)
	}
	fmt.Println(sessionTable)
	for i, s := range sessionTable {
		if s.Token.String() == tokenC.Value {
			sessionTable[i] = sessionTable[len(sessionTable)-1]
			sessionTable = sessionTable[:len(sessionTable)-1]
		}
		fmt.Println(sessionTable)
		
	}
    http.Redirect(w, r, "/logIn", http.StatusSeeOther)
}

func handleFunc() {
    router := mux.NewRouter()
    router.HandleFunc("/", show_list_subjects)
    router.HandleFunc("/logIn", logIn)
	router.HandleFunc("/postLogIn", postLogIn)
    router.HandleFunc("/main", index)
    router.HandleFunc("/new_teacher", new_teacher)
    router.HandleFunc("/save_teacher", save_teacher)
	router.HandleFunc("/new_coursePr", new_coursePr)
    router.HandleFunc("/save_coursePr", save_coursePr)
    router.HandleFunc("/new_group", new_group)
    router.HandleFunc("/save_group", save_group)
    router.HandleFunc("/new_subject", new_subject)
    router.HandleFunc("/save_subject", save_subject)
    router.HandleFunc("/new_student", new_student)
    router.HandleFunc("/save_student", save_student)
    router.HandleFunc("/del_student", del_student)
    router.HandleFunc("/subjects/", show_list_subjects)
	router.HandleFunc("/subjects/show_coursePrs", show_list_coursePr)
    router.HandleFunc("/subjects/{descr}", show_subject)
    router.HandleFunc("/show_groups_and_subjs/{grname}+{subjname}", show_students)
    router.HandleFunc("/show_groups_and_subjs/{grname}+{subjname}/{subjname}+{studname}", show_marks)
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
	router.HandleFunc("/del_coursePr", del_coursePr)
	router.HandleFunc("/delete_coursePr", delete_coursePr)
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
	router.HandleFunc("/update_coursePr", update_coursePr)
	router.HandleFunc("/upd_coursePr", upd_coursePr)
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
    router.HandleFunc("/new_connTGS", new_connTGS)
    router.HandleFunc("/save_connTGS", save_connTGS)
    router.HandleFunc("/show_groups_and_subjs", show_list_groups_with_subj)
    router.HandleFunc("/save_exinst/{stid}+{exid}",save_exinst)
    router.HandleFunc("/save_lectatt/{stid}+{ltid}",save_lectatt)
    router.HandleFunc("/save_sematt/{stid}+{smid}",save_sematt)
    router.HandleFunc("/save_bcinst/{stid}+{bcid}",save_bcinst)
    router.HandleFunc("/save_labi/{stid}+{lbid}",save_labi)
    router.HandleFunc("/postLogOut", postLogOut)
    http.Handle("/", router)
    fmt.Println("Server is listening...")
    http.ListenAndServe("localhost:8080", nil)
}

func main() {
    handleFunc()
}


