package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func MonthDayYear(t time.Time) string{
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": MonthDayYear,
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil{
		log.Fatal(err)
	}
}