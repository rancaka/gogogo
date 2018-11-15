package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	web     = New("web")
	ads     = New("ads")
	youtube = New("youtube")
)

func main() {
	start := time.Now()
	results := []Result{}
	results = append(results, web("mantul"))
	results = append(results, ads("mantul"))
	results = append(results, youtube("mantul"))
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
