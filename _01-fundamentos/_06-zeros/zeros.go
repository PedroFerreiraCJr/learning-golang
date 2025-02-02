package main

import "fmt"

func main() {
	// o valor padrão de uma variável do tipo 'int' é zero (0);
	var a int

	// o valor padrão de um variável do tipo 'float64' é zero ponto zero (0.0);
	var b float64

	// o valor padrão de uma variável do tipo 'bool' é falso (false);
	var c bool

	// o valor padrão de uma variável do tipo 'string' é string vazia ("");
	var d string

	// o valor padrão de uma variável do tipo '*int' (ponteiro para int), é o valor 'nil'
	var e *int

	fmt.Printf("%v, %v, %v, %q, %v", a, b, c, d, e)
}
