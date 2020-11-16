package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := boring("Bob", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Bob says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // Return receive-only channel of strings
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("All clean!")
}
