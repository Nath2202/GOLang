package main

import (
	"fmt"
)

func fatora(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatoraAnterior := fatora(n - 1)
		return n * fatoraAnterior
	}
}

func executa(funcao func(uint) uint, n1 uint) uint {
	return funcao(n1)
}

func main() {
	resultado := executa(fatora, 3)
	fmt.Println(resultado)
}
