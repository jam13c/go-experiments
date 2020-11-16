package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplicas := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplicas(i)
	}
	return <-c
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Search function
type Search func(query string) Result

//Result string
type Result string
