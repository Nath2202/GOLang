package main

import (
	"fmt"
)

//func <nome função> (<nome> tipo, <nome> tipo) tipo retorno
func soma(a int, b int) int {
return a + b
}

func imprime(c int){
fmt.Println(c)
}

func main(){
	resultado:= soma(3,4)
	imprime(resultado)
}

