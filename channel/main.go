package main

import "fmt"

func main() {
	//harapan dari program ini adalah saling sync antara goroutine

	// buat channel dengan nama done
	done := make(chan bool)

	go helloGo(done)
	// ambil value dari channel done
	<-done

	fmt.Println("Ini Fungsi main")
}

//buat fungsi yang memiliki paramater channel boolean
func helloGo(done chan bool) {
	fmt.Println("Hello GO")
	//isi value done dengan nilai true
	done <- true
}
