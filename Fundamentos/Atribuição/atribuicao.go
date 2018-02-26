package main

import (
	"fmt"
)

func main() {
	var b byte = 3
	fmt.Println(b)

	//tipos de atribuição
	i := 3
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3
    fmt.Println(i)

	//atribuição de multiplos valores
	x, y := 2, 3
    fmt.Println(x,y)
	//troca de valores entre x e y
	x, y = y, x
	fmt.Println(x,y)

}
