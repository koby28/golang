package main

import (
	"net/http"
	"io"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></html></body>`

func SimpleServer(w http.ResponseWriter, request *http.Request)  {
	io.WriteString(w,"<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request)  {
	w.Header().Set("Content-Type","text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w,form)
	case "POST":
		request.ParseForm()
	    io.WriteString(w,request.Form["in"][1])
	    io.WriteString(w,"\n")
	    io.WriteString(w,request.FormValue("in"))
	}
}
func main() {
    http.HandleFunc("/test1",SimpleServer)
    http.HandleFunc("/test2",FormServer)
    if err := http.ListenAndServe(":8089",nil);err != nil{}
}
