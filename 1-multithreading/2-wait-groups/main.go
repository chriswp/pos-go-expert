package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("task1", &waitGroup)
	go task("task2", &waitGroup)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d : Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()

}
