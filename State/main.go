package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//to fivicon not loaded, coz it's not allowed to load on the page - obvious
	//https://github.com/GoesToEleven/golang-web-dev/blob/master/027_passing-data/05_form-file/01_read/main.go
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	v := req.FormValue("q")
	// it's (f, h, err := req.FormFile) i we wanna file
	io.WriteString(w, "U r searching for " + v)
}

//localhost:8080/?q=dog
//it works by GET and POST
//only one thing is in URL, it has always 'q' by default
//in POST it's 'name' where something like 'q' contains

//https://github.com/GoesToEleven/golang-web-dev/tree/master/027_passing-data