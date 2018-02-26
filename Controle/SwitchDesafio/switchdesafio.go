package main

import (
	"fmt"
)

func notaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n <= 9:
		return "B"
	case n >= 5 && n <= 8:
		return "C"
	case n >= 3 && n <= 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaConceito(9.8))
	fmt.Println(notaConceito(5.8))
	fmt.Println(notaConceito(2.8))
	fmt.Println(notaConceito(0.0))
}
