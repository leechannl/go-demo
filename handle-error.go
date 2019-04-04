package main

import "github.com/pkg/errors"
import "fmt"

func A() error {
	return errors.New("NullPointerException")
}

func B() error {
	return A()
}

func main() {
	fmt.Printf("Error: %+v\n", B())
}
