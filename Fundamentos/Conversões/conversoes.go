package main

import (
	"fmt"
	"strconv"
)

func main() {
	//inteiro para float64
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	//utilizamos esses conversores explicitos para que o compilador possa realizar a conta

	//float64 para int
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	//Obs.: essa função retorna dois valores: o valor convertido e um erro, caso o valor passado como parâmetro não seja uma string
	value, _ := strconv.Atoi("123")
	fmt.Println(value - 122)

	//string para booleano
	//Obs.: ele aceita como verdadeiro <1,t,T,true,TRUE,True> e falso <0,f,F,false,FALSE,False>
	//fora esses valores acima, ele retorna um erro
	aux, _ := strconv.ParseBool("f")

	fmt.Println(aux)

}
