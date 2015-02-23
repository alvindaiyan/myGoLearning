package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //get url param, for POST method, get the body

	// notice: without注意:如果没有调用ParseForm方法，下面无法获取表单的数据

	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write back to the client
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get the http method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		//print out the form info
		fmt.Println("username:", r.Form["username"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //setup the visting url
	http.HandleFunc("/login", login)         //setup the visting url
	err := http.ListenAndServe(":9090", nil) //setup the listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
