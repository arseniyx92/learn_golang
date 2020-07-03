package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var dp *sql.DB
var err error

func main(){
	db, err := sql.Open("mysql", "arseniyx92:123@tcp(34.65.59.199:3306)/users?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request){
	_, err := io.WriteString(w, "Successfully connected.")
	check(err)
}

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}