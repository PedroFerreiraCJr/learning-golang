package main

import "fmt"

func main() {
	// a seguinte função serve para criar um slice a partir de alguns argumentos;
	// o primeiro argumento é o tipo do slice;
	// o segundo argumento é o número de elementos do slice;
	// o funcionamento da função 'make', para os argumentos informados, é criar um array
	// de 10 posições, onde o slice retornado faz referência para o primeiro
	// índice (índice '0'), do array criado;
	s := make([]int, 10)

	// atribui o valor '42' à última posição do slice;
	s[9] = 42
	fmt.Println(s)

	// slice de inteiros, com 10 elementos, com um array interno de 20 posições;
	// os valores dos elementos do array criado, são inicializados para o valor padrão de acordo
	// com o tipo de dado informado; por isso que é possível ler qualquer uma das posições do
	// array interno através do slice retornado pela função 'make';
	s = make([]int, 10, 20)

	// a função 'len', aplicada a um slice, retorna o tamanho dele;
	// a função 'cap', aplicada a um slice, retorna a capacidade máxima do array que o slice referencia;
	fmt.Println(s, len(s), cap(s))

	// anexa ao slice 's', os valores informados como argumento para a função 'append';
	// a função 'append', retorna um novo slice, que aponta para o mesmo array;
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// neste exemplo, a capacidade máxima do array intero referênciado pelo slice 's', alcança
	// o tamanho máximo, e então é substituído por outro array; então, os dados do array anterior
	// são copiados para o array que acabou de ser instanciado;
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
