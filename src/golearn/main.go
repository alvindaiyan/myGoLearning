package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func print_binary(s []byte) {
	fmt.Println("Received b")

	for n := 0; n < len(s); n++ {
		fmt.Printf("%d", s[n])
	}
	fmt.Println("\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		print_binary(p)

		err = conn.WriteMessage(messageType, []byte("this is the message from server"))
		if err != nil {
			return
		}
	}
}

type Message struct {
	Id      int
	Content string
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	for {
		var m Message
		fmt.Println("start reading...")
		err := conn.ReadJSON(&m)
		// messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("errorhere:", err)
			return
		}

		fmt.Printf("text: %v\n", m)
		// print_binary(p)
		// fmt.Println(string(p))
		// fmt.Println(messageType)

		mr := Message{
			Id:      1,
			Content: "this is a return message",
		}

		err = conn.WriteJSON(mr)
		if err != nil {
			return
		}
	}
}

func main() {
	jsonstr := []byte(`{"id":1,"content":"test"}`)
	var m Message
	err1 := json.Unmarshal(jsonstr, &m)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	fmt.Printf("%v", m)

	m1 := Message{Id: 1, Content: "text"}
	b, err2 := json.Marshal(m1)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	fmt.Println("text: ", string(b))

	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/json", jsonHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
