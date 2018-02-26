package main

import (
	f "fmt"
	m "math"
)

func main(){

	const PI float64 = 3.1415
	var raio = 3.2
	
	//forma reduzida de declarar var

	area := PI * m.Pow(raio, 2)

	f.Println("Area da circuferência é", area)

	const(
		a = 1
		b = 2
	)

	var(
		c = 3
		d = 4
	)

	f.Println(a,b,c,d)

	//e, f := false, true
var e, h bool = true, false
	f.Println(e,h)

	g, j, k := 10, "dez", false

	f.Println(g,j,k)

}	