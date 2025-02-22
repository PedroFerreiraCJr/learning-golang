package main

import "fmt"

// internamente dentro da função, o parâmetro variático é acessado como
// um array;
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// isso é uma declaração de um slice, porque não foi definido o tamanho
	// desta estrutura, conforme é obrigatório no array;
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}

	// para diferenciar um array de um slice, basta utilizar o operador
	// '...', conforme o seguinte exemplo;
	//aprovados := [...]string{"Maria", "Pedro", "Rafael", "Guilherme"}

	// essa é a forma de passar como argumento de uma função variática, um
	// argumento do tipo 'slice', que neste caso é a variável 'aprovados';
	// neste caso, o compilador fará o espalhamento dos valores contidos no
	// argumento que é um slice, para os argumentos da função
	// 'imprimirAprovados';
	imprimirAprovados(aprovados...)

	// no caso acima, não seria possível passar como argumento da função
	// 'imprimirAprovados', um argumento do tipo 'array', independente da forma
	// como ele fosse declarado, como por exemplo, como uma declaração de variável
	// com um número específico de valores definidos no array, ou deixando para o
	// compilador contar, automaticamente, a quantidade de valores do array, e desta
	// forma, o tamanho do array fosse definidos de forma automática;
	sample1 := [2]string{"a", "b"}
	sample2 := [...]string{"a", "b", "c"}

	fmt.Println(sample1)
	fmt.Println(sample2)

	// essas duas linhas não compilam;
	//imprimirAprovados(sample1...)
	//imprimirAprovados(sample2...)
}
