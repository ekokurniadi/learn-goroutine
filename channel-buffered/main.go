package main

import "fmt"

/*
Note :
Un-buffered channel :
1. done <-true(sender)  -> sender akan menunggu / akan jalan ketika sudah ada penerima yang siap
2. <-done(receiver) -> penerima
*/

func main() {
	// membuat buffered channel
	hello := make(chan string, 2)
	hello <- "Hello"
	hello <- "Golang"
	close(hello)
	// fmt.Println(<-hello)
	// fmt.Println(<-hello)

	for value := range hello {
		fmt.Println(value)
	}
}
