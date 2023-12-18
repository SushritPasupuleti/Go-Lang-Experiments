package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doSomethingToChan(input int, wg *sync.WaitGroup, outputChan chan int) {
	iter := input
	randNum := rand.Intn(100)

	input = randNum * (input + 1)

	fmt.Println("Iter: ", iter, "Input:", input)

	// input = factorial(input)

	outputChan <- input

	wg.Done()
}

func sum(input []int) int {
	total := 0
	for _, val := range input {
		total += val
	}
	return total
}

func main() {
	var start time.Time
	outputChan := make(chan int)

	start = time.Now()
	fmt.Println("Start:", start)

	runs := 3

	wg := sync.WaitGroup{}
	wg.Add(runs)

	output := []int{}

	for i := 0; i < runs; i++ {
		go doSomethingToChan(i, &wg, outputChan)
		output = append(output, <-outputChan)
	}

	wg.Wait()

	fmt.Println(output)

	fmt.Println("Sum:", sum(output))

	end := time.Now()
	fmt.Println("End:", end)
}
