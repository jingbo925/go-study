package main

import "fmt"

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("workder %d received %c\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

}
func main() {
	chanDemo()
}
