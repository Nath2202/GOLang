package main

import "fmt"

func main() {
	var aprovados map[int]string
	//maps devem ser inicializados
	aprovados = make(map[int]string)
	aprovados[45554321852] = "Nathalia"
	aprovados[44585609807] = "Leonardo"
	//fmt.Println(aprovados)

	for id, value := range aprovados {
		fmt.Printf("CPF: %d - %s\n", id, value)

	}
	delete(aprovados, 45554321852)
}
