package main

import (
	"fmt"
)

type curso struct {
	nome string
}

func main() {

	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = "Nathalia"
	fmt.Println(coisa)

	coisa = curso{"Eng. da Computação"}
	fmt.Println(coisa)
}
