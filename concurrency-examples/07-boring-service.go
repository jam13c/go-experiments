package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	bob := boring("Bob")
	ann := boring("Ann")
	for i := 0; i < 10; i++ {
		fmt.Println(<-bob)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving!")
}

func boring(msg string) <-chan string { // Return receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c
}
