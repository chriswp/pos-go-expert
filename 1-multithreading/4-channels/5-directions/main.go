package main

import "fmt"

func receiver(name string, out chan<- string) {
	out <- name
}

func reader(out <-chan string) {
	fmt.Println(<-out)
}

func main() {
	ch := make(chan string)
	go receiver("Fera", ch)
	reader(ch)
}
