package main

import "fmt"

type nota float64

func (n nota) between(begin, end float64) bool {
	return float64(n) >= begin && float64(n) <= end
}

func notaParaConceito(n nota) string {
	switch {
	case n.between(9.0, 10.0):
		return "A"
	case n.between(7.0, 8.99):
		return "B"
	case n.between(5.0, 6.99):
		return "C"
	case n.between(3.00, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.0))
}
