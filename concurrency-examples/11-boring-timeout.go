package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Bob")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boring(msg string) <-chan string { // Return receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1.5e3)) * time.Millisecond)
		}
	}()
	return c
}
