package main

import "fmt"

func multiplica(a, b int) int {
	return a * b
}

func executa(funcao func(int, int) int, v1, v2 int) int {
	return funcao(v1, v2)
}

func main() {
	resultado := executa(multiplica, 3, 2)
	fmt.Println(resultado)
}
