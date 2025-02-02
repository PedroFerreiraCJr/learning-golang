package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix - Golang só tem a forma posfixada de operadores unários;
	// operador de incremento;
	x++
	fmt.Println(x)

	// operador de decremento;
	y--
	fmt.Println(y)

	// essa linha não compila, porque vai contra o conceito de Go para tornar a linguagem
	// de fácil entendimento;
	//fmt.Println(x == y--)
}
