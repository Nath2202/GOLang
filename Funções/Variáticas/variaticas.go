package main

import (
	"fmt"
)

func media(notas ...float64) float64 {
	total := 0.0

	for _, num := range notas {
		total += num
	}
	return total / float64(len(notas))
}

func main() {
	fmt.Printf("Média = %.2f\n", media(2.2, 3.4, 5.9, 10.0))
}
