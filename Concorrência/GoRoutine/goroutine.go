package main

import (
	"fmt"
	"time"
)

func falar(pessoa, texto string, num int) {
	for i := 0; i < num; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// falar("Maria", "Porque não fala comigo?", 3)
	// falar("Leonardo", "Eu te amo!", 5)

	go falar("Nathalia", "Ta pegando fogo bicho", 10)
	time.Sleep(time.Second*6)
}
