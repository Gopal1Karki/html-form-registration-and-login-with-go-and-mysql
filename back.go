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
	tpl = template.Must(template.ParseFiles("signup.html"))

}

func main() {
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		tpl.Execute(w, nil)
		return
	}

	user := User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
	}
	display(user.Username, user.Email, user.Password)
	_, err := db.Exec(`insert into info (username,email,password) value (?,?,?)`,
		user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "User created successfully!")
}

func display(a string, b string, c string) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
