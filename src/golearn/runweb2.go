package main

import (
    "fmt"
    "net/http"
)


type String string
func (h String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, h)
}

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (h Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, h.Greeting, h.Punct, h.Who)
}



func main() {
    // your http.Handle calls here
    http.Handle("/string", String("<b>I'm a frayed knot.</b>"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}
