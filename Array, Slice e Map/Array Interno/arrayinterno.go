package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3) //atribuindo ao slice s2 o slice s1 e mais 3 valores
	s1[0] = 7
	fmt.Println(s1, s2)
}
