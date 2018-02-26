package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto01 produto
	produto01 = produto{
		nome:     "Lapis",
		preco:    2.00,
		desconto: 0.05,
	}
	fmt.Println(produto01)
	fmt.Println(produto01.precoComDesconto())
}
