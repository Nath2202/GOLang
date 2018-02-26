package main

import (
	"fmt"
)

func main() {
	i := 1

	//não é possível fazer aritmética com ponteiros

	var p *int = nil
	p = &i
	*p++ //acessando o valor da variável
	i++

	fmt.Println(p, *p, i, &i, &p)
}
