package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("task1")
	go task("task2")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d : Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
}
