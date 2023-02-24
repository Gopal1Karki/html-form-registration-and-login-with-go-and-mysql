package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
}

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/login?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	tpl = template.Must(template.ParseFiles("login.html"))
}

func main() {
	//http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

//unc signup(w http.ResponseWriter, r *http.Request) {
// code for creating a new user account
//}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		tpl.Execute(w, nil)
		return
	}
	type User struct {
		Username string
		Password string
	}
	var s User
	Username := r.FormValue("username")
	Password := r.FormValue("password")
	err := db.QueryRow(`select password from info where username = ?`, Username).Scan(&s.Password)
	if err != nil {
		log.Fatal(err)
	}
	if Password == s.Password {
		fmt.Fprintf(w, "Login Sucessfull!")
	}
}
