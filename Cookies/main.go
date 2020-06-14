//small app that allow us to check how many times our site has been visited
//by particular person (browser)
package main

import (
	"io"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/", SetAndRead)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func SetAndRead(w http.ResponseWriter, req *http.Request){

	cookie, err := req.Cookie("how_many_times_visited")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "how_many_times_visited",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	if count == 1 {
		io.WriteString(w, "You've visited our site " + cookie.Value + " time")
		return
	}
	io.WriteString(w, "You've visited our site " + cookie.Value + " times")
}