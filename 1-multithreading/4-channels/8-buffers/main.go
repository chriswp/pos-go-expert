package main

func main() {
	ch := make(chan string, 2)
	ch <- "Fera"
	ch <- "Férias"

	println(<-ch)
	println(<-ch)
}
