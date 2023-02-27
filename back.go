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
	Email    string
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

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func signup(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("signup.html"))
	if r.Method == "POST" {
		user := User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}
		_, err := db.Exec(`insert into info (username,email,password) value (?,?,?)`, user.Username, user.Email, user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "User created successfully!")
		return
	}

	tpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("login.html"))
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		var storedPassword string
		err := db.QueryRow(`SELECT password FROM info WHERE username = ?`, username).Scan(&storedPassword)
		if err != nil {
			fmt.Fprintf(w, "Login failed!")
			return
		}
		if password == storedPassword {
			fmt.Fprintf(w, "Login successful!")
			return
		}
		fmt.Fprintf(w, "Login failed!")
		return
	}

	tpl.Execute(w, nil)
}
