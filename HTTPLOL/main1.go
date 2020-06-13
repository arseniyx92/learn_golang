package main

import (
	"io"
	"net/http"
)

func cat(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "cat cat cat")
}

func dog(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "doggy dog")
}

func main(){
	//http.HandleFunc("/dog/", dog)
	//http.HandleFunc("/cat", cat)
	//or
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/cat", http.HandlerFunc(cat))
	http.ListenAndServe(":8080", nil)
}

