package main

import "fmt"

// essa é uma declaração de variável, que recebe como valor uma função;
// essa forma de declarar uma função é conhecida por Função Anônima;
var somar = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(somar(6, 14))

	// essa é uma declaração de variável, com a notação sucinta, e que está
	// recebendo na sua inicialização uma função;
	// essa forma de declarar uma função é conhecida por Função Anônima;
	subtrair := func(a, b int) int {
		return a - b
	}
	fmt.Println(subtrair(10, 5))
}
