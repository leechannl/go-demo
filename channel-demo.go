package main

import "fmt"
import "time"

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	go multiplyByTwo(in, out)
	in <- n
	fmt.Println(<-out)
}

func multiplyByTwo(int <-chan int, out chan<- int) {
	result := <-int * 2
	time.Sleep(time.Second * 3)
	out <- result
}
