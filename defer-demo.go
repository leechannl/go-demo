package main

import "fmt"

// A deferred function's arguments are evaluated when the defer statement is evaluate
func a() {
	i := 0
	defer fmt.Println("first defer: ", i)
	i++
	defer fmt.Println("second defer: ", i)
	return
}

// Deferred function calls are executed in LIFO order after the surrounding function returns
func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Deferred functions may read and assign to the returning function's named return values
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	b()
	fmt.Println("c() return: ", c())
}
