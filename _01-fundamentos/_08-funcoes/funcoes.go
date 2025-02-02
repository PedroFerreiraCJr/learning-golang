package main

import "fmt"

// declara uma função chamada 'somar', que recebe dois parâmetros do tipo 'int';
// essa função declarou explicitamente que deve retornar um valor do tipo 'int',
// como resultado da sua execução com sucesso;
func somar(a int, b int) int {
	return a + b
}

// declara uma função chamada 'imprimir', que recebe um parâmetro do tipo 'int';
// essa função não declara um tipo de retorno, portanto, neste caso, a função
// não precisa utilizar a palavra-chave 'return' para devolver um valor para
// quem a invocar;
func imprimir(valor int) {
	fmt.Println(valor)
}
