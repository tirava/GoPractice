package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 2)
	for i := 0; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock
}
