package main

import "fmt"

type imprimivel interface{
	toString() string
}

type pessoa struct{
	nome string
	sobrenome string
}

type produto struct{
	nome string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome+ " "+p.sobrenome

}

func (pr produto) toString() string{
	return fmt.Sprintf("%s - R$ %.2f", pr.nome, pr.preco)
}

func imprimir(x imprimivel){
 fmt.Println(x.toString())
}

func main(){
	var testePessoa imprimivel = pessoa{"Nathalia","Lourenço"}
	fmt.Println(testePessoa.toString())
	imprimir(testePessoa)

	var testeProduto imprimivel = produto{"Banana", 2.00}
	fmt.Println(testeProduto.toString())
	imprimir(testeProduto)
}