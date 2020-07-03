package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var db *sql.DB

//User is ...
type User struct {
	NameErr  int
	Err      int
	Name     string
	Email    string
	Password []byte
}

func init() {
	db, err := sql.Open("mysql", "arseniyx92:123@tcp(34.65.166.197)/users?charset=utf8")
	check(err)
	err = db.Ping()
	check(err)
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets", http.FileServer(http.Dir("./stylesheets"))))
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "signup", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
		return
	}
	un := req.FormValue("name")
	email := req.FormValue("email")
	pas := req.FormValue("password")
	pas1 := req.FormValue("password1")
	pas2, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost)
	check(err)

	CurUser := User{
		NameErr:  2,
		Err:      0,
		Name:     un,
		Email:    email,
		Password: pas2,
	}

	// rows, err := db.Query(`SELECT * FROM users WHERE UserName=` + CurUser.Name)
	// rows, err := db.Query(`SELECT UserName FROM users`)
	// check(err)
	// var name string
	// cnt := 0
	// for rows.Next() {
	// 		err = rows.Scan(&name)
	// 		check(err)
	// 		cnt++
	// }
	// fmt.Println(cnt)

	if pas != pas1 {
		//TODO
		CurUser.Err = 1
		// fmt.Println(CurUser.Email)
		tpl.ExecuteTemplate(w, "signup.gohtml", CurUser)
		return
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
