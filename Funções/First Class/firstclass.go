package main

import (
	"fmt"
)

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(4, 5))

	var sub = func(c,d int)int{
		return c-d
	}
	fmt.Println(sub(2,1))
}
