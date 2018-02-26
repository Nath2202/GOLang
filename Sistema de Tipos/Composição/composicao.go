package main

import "fmt"

type gato struct {
	qtdePatas   int
	qtdeOrelhas int
	cor         string
}

type partes struct {
	gato // campo anonimo
}

func main() {
	g := partes{}
	g.qtdePatas = 4
	g.qtdeOrelhas = 2
	g.cor = "Branco"

	fmt.Printf("O gato tem %d patas, %d orelhas e ele é %s\n", g.qtdePatas, g.qtdeOrelhas, g.cor)
}
