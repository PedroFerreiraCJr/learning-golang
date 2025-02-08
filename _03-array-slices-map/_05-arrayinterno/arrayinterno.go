package main

import "fmt"

// essa aula é uma tentativa de provar que o slice aponta para um array
// interno;
func main() {
	// declara um slice chamado 's1', que é inicializado através da função
	// 'make', que retorna um slice; como argumentos ela recebe o tipo do array,
	// que neste caso, é um array de inteiros; o segundo argumento é o tamanho do
	// slice, que aponta para o array interno; esse array interno tem capacidade
	// para 20 elementos, que é exatamente o valor informado no terceiro argumento;
	s1 := make([]int, 10, 20)

	// a função 'append', retorna um novo slice, que aponta para o array interno
	// informado através do slice 's1'; a função 'append', ainda recebe argumentos
	// adicionais, que são os valores que devem ser adicionados ao final do array
	// interno apontado pelo argumento 's1';
	s2 := append(s1, 1, 2, 3)

	// esse print, prova que os valores foram adicionados no final do array interno
	// apontado pelo slice 's1';
	fmt.Println(s1, s2)

	// tentativa de provar que os dois slices (s1, s2), apontam para o mesmo
	// array interno; porque, neste caso, como ambos os slices apontam para o
	// mesmo array interno, então, alterando o valor de um slice, deve fazer
	// com que o outro slice, também seja afetado pela alteração;
	s1[0] = 7
}
