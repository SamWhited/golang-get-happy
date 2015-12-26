package main

import (
	"errors"
	"fmt"
)

func test() (a, b int, err error) {
	a = 1
	b = 1
	err = errors.New("Test")
	return
}

func main() {
	fmt.Println(test())
}
