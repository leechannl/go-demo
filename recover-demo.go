package main

import "fmt"
import "github.com/pkg/errors"

func R() {
	defer fmt.Println("R")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
	S()
}

func S() {
	defer fmt.Println("S")
	T()
}

func T() {
	defer fmt.Println("T")
	Break()
}

func Break() {
	defer fmt.Println("U")
	panic(errors.New("the show must go on"))
}

func main() {
	R()
}
