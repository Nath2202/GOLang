package main

import (
	"fmt"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não conheço"
	}
}

func main() {
	fmt.Println(tipo(4.3))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(2))
	fmt.Println(tipo(func() {}))
}
