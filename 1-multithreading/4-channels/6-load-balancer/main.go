package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Println("worker", workerId, "received data", x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 1000000

	//inicializa os workers
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
