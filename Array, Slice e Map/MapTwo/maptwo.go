package main

import "fmt"

func main() {
	funcionariosESalarios := map[string]float64{
		"José":   5000.00,
		"Maria":  6000.00,
		"Júnior": 7000.00,
	}

	for nome, salario := range funcionariosESalarios {
		fmt.Println(nome, salario)

	}
}
