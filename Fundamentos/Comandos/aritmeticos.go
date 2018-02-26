package main

import (
	"math"
	"fmt"
)

func main(){
	a:= 3
	b:= 3
	
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b)

	//bitwise: faz o calculo binário dos valores e retorna o valor inteiro deles
	//logica AND
	fmt.Println("E =>", a&b)
	//logica OR
	fmt.Println("OU =>", a|b)
	//logica XOR
	fmt.Println("XOR =>", a^b)
	
	c:= 3.0
	d:= 2.0

	//operações usando math
	fmt.Println("Maior valor: ", math.Max(c,d))
	fmt.Println("Menor valor: ", math.Min(c,d))
    fmt.Println("Exponenciação: ", math.Pow(c,d))
}