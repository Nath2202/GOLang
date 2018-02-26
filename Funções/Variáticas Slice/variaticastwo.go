package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Printf("Lista de Aprovados\n")
	for i, aprovado := range aprovados {
		fmt.Printf("%d - %s\n", i+1, aprovado)

	}
}

func main() {
	aprovados := []string{"João", "Maria", "Eduardo", "André"}
	imprimirAprovados(aprovados...)
}
