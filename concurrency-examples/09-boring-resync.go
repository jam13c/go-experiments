package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Bob"), boring("Ann"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring; I'm leaving!")
}

func boring(msg string) <-chan message { // Return receive-only channel of strings
	c := make(chan message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan message) <-chan message {
	c := make(chan message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

type message struct {
	str  string
	wait chan bool
}
