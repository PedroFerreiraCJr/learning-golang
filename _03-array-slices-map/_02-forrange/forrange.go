package main

import "fmt"

func main() {
	// essa é outra forma de fazer a inicialização de um array;
	// nela, está sendo solicitado que o compilador faça a contagem da quantidade de elementos
	// contidos dentro do operador chaves, e a partir disto, ele defina o tamanho do array;
	// uma coisa importante sobre essa sintaxe é que, caso seja removido da seguinte declaração
	// o símbolo de reticências, obtermos um 'Slice', ao invés de um array;
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	// nesta nova forma de declarar um 'for', estamos utilizando um 'range' para que ele retorne
	// o valor das variáveis 'i', e 'numero';
	// o valor da primeira variável, conterá o valor do índice;
	// o valor da segunda variável, conterá o valor o valor sem si, do array que está sendo
	// percorrido;
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// essa é outra forma de fazer a mesma coisa, que é percorrer os valores do array;
	// mas, nessa abordagem, está sendo informado que o valor da primeira variável, que foi
	// definida com o símbolo de underscore, deve ser ignorado, ou seja, não deve considerar
	// isso como uma declaração de variável, e portanto, que não estaria sendo usado, o que
	// geraria um erro de compilação, porque a linguagem Go não permite que seja declarado uma
	// variável e ela não esteja sendo utilizada no programa;
	for _, num := range numeros {
		// utiliza somente a variável 'num';
		fmt.Println(num)
	}

	// essa notação de 'range' é semelhante ao foreach de outras linguagens;
}
