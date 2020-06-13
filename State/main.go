package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	v := req.FormValue("q")
	io.WriteString(w, "U r searching for " + v)
}

//localhost:8080/?q=dog
//it works by GET and POST
//only one thing is in URL, it has always 'q' by default
//in POST it's 'name' where something like 'q' contains