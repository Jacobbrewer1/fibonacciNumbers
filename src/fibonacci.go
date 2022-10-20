package main

import (
	"log"
)

func init() {
	log.Println("initializing logging")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("logging initialized")
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func main() {
	const workers int = 8
	const limit int = 50

	var jobs = make(chan int, limit)
	var results = make(chan int, limit)

	for i := 0; i < workers; i++ {
		go func(x int) {
			log.Printf("starting worker %d\n", x+1)
			worker(jobs, results)
		}(i)
	}

	for i := 0; i < limit; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < limit; i++ {
		log.Println(<-results)
	}
	close(results)
}
