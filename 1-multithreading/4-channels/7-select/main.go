package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(2 * time.Second)
			c1 <- Message{i, "Hello from RabbitMQ"}
		}

	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			c2 <- Message{i, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("ID: %d, Msg: %s\n", msg1.id, msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("ID: %d, Msg: %s\n", msg2.id, msg2.Msg)
		case <-time.After(3 * time.Second):
			println("timeout")
			//default:
			//	println("default")
		}
	}
}
