package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", me)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func me(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/me.jpg">`)
}
