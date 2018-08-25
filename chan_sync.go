package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done //if this line is removed, the worker will never even run
}
