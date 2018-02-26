package main

import (
	"fmt"
)

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido01 := pedido{
		userID: 1,
		itens: []item{
			item{1, 12, 10.00},
			item{2, 4, 6.00},
			item{3, 8, 9.00},
		},
	}
	fmt.Printf("Valor total do pedido é %.2f\n", pedido01.valorTotal())
}
