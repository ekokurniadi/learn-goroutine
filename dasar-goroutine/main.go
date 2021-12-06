package main

import (
	"fmt"
	"time"
)

func main() {
	// menerapkan goroutine
	go helloGo()
	//beri jeda untuk menampilkan goroutine
	time.Sleep(1 * time.Second)
	//print func dari main
	fmt.Println("Ini dari Func main")
}

func helloGo() {
	fmt.Println("Hello Go Function")
}
