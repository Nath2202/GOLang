package main

import (
	"time"
	"fmt"
)

func main(){

	fmt.Println("Strings: ", "Banana" == "Banana")
	fmt.Println("!=", 3!=2)
	fmt.Println("<", 3<2)
	fmt.Println(">", 3>2)
	fmt.Println("<=", 3<=2)
	fmt.Println(">=", 3>=2)

	//tipo data
	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)

	fmt.Println("Datas: ", d1 == d2)

	//struct
	type Pessoa struct{
		Nome string
	}

	p1 := Pessoa{"Nathalia"}
    fmt.Println(p1)
}