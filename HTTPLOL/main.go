package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil{
		log.Panicln(err)
	}
	w.Header().Set("Content-type", "text-html; charset=utf-8")
	tmp.ExecuteTemplate(w, "tmp.gohtml", req.Form)
}

func init(){
	tmp = template.Must(template.ParseGlob("templates/*.gphtml"))
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080", d)
}
//package main
//
//import (
//	"html/template"
//	"log"
//	"net/http"
//)
//
//type hotdog int
//
//var tpl *template.Template
//
//func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
//	err := req.ParseForm()
//	//changing header
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	if err != nil{
//		log.Fatalln(err)
//	}
//	tpl.ExecuteTemplate(w, "tpl.gohtml", req.Form)
//}
//
//func init(){
//	tpl = template.Must(template.ParseGlob("templates/*"))
//}
//
//func main(){
//	var d hotdog
//	http.ListenAndServe(":8080", d)
//}



//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type hotdog int
//
//func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintln(w, "Lolik I'm, or NOT?!")
//}
//
//func main(){
//	var d hotdog
//	http.ListenAndServe(":8080", d)
//}