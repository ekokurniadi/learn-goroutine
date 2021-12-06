package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Worker Pool")

	start := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// go consumer(1, jobs, results)
	// go consumer(2, jobs, results)
	for i := 1; i <= 4; i++ {
		go consumer(i, jobs, results)
	}
	go producer(jobs, 10)

	// looping sejumlah job yang dikirim
	for i := 1; i <= 10; i++ {
		res := <-results

		fmt.Println("Hasil ke ", res)
	}
	fmt.Println("=======================")
	elapsed := time.Since(start)
	fmt.Println("waktu ", elapsed)

}

func FakeHTTPRequest(x int) int {
	return x * 10
}

// untuk menghasilkan jobs dalam bentuk channel
func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Consumer ke", id, "Mulai")
		time.Sleep(time.Second * 2)
		results <- FakeHTTPRequest(job)
	}
}
