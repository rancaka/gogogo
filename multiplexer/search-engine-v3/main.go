package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	web      = New("web replica 1")
	web2     = New("web replica 2")
	web3     = New("web replica 3")
	ads      = New("ads replica 1")
	ads2     = New("ads replica 2")
	ads3     = New("ads replica 3")
	youtube  = New("youtube replica 1")
	youtube2 = New("youtube replica 2")
	youtube3 = New("youtube replica 3")
)

func main() {
	start := time.Now()

	result := make(chan Result)
	go func() {
		result <- ReplicateSearch("mantul", web, web2, web3)
	}()
	go func() {
		result <- ReplicateSearch("mantul", ads, ads2, ads3)
	}()
	go func() {
		result <- ReplicateSearch("mantul", youtube, youtube2, youtube3)
	}()

	results := []Result{}
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case res := <-result:
			results = append(results, res)
			continue
		case <-timeout:
		}
		break
	}
	totalDuration := time.Since(start)

	for i, result := range results {
		fmt.Printf("%v. %v\n", i+1, result)
	}
	fmt.Printf("total duration: %v\n", totalDuration)
}

func New(source string) Search {
	return func(query string) Result {
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		result := Result(fmt.Sprintf("%v result of \"%v\" took %v", source, query, time.Since(start)))
		return result
	}
}

func ReplicateSearch(query string, search ...Search) Result {
	result := make(chan Result)
	searchReplicate := func(i int) { result <- search[i](query) }
	for i := range search {
		go searchReplicate(i)
	}
	return <-result
}
