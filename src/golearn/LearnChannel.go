package main

import "fmt"

type Message struct {
	content string
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func addstring(a string, c chan string) {
	c <- a
}

func addmsg(a Message, c chan Message) {
	c <- a
}

// this is not buffered channel
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// buffered channel
	c2 := make(chan int, 2)
	c2 <- 0
	c2 <- 1
	fmt.Println(<-c2, <-c2)

	strc := make(chan string)

	go addstring("first message", strc)
	go addstring("second message", strc)

	// strc <- "first message"
	// strc <- "second message"

	fmt.Println(<-strc, <-strc)

	msgc := make(chan Message)

	// go addstring("first message", strc)
	// go addstring("second message", strc)

	go addmsg(Message{"first msg"}, msgc)
	go addmsg(Message{"second msg"}, msgc)

	msg1, msg2 := <-msgc, <-msgc

	fmt.Println(msg1.content, msg2.content)
}
