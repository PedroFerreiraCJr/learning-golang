package main

import "fmt"

func main() {
	// declara uma variável, de maneira explícita, utilizando a palavra-chave 'var';
	// essa variável tem o tipo explícito definido como 'byte';
	var b byte = 3
	fmt.Println(b) // utiliza a variável 'b', para que não tenha erro de compilação;

	i := 3         // utiliza a inferência de tipo para declarar e inicializar a variável;
	i += 3         // i = i + 3
	i -= 3         // i = i - 3
	i /= 2         // i = i / 2
	i *= 2         // i = i * 2
	i %= 2         // i = i % 2
	fmt.Println(i) // essa linha não é necessária para compilar, porque a variável 'i' está sendo utilizada em outras operações;

	// assim como em outras linguagens, como por exemplo, a linguagem C, é possível
	// declarar mais de uma variável na mesma linha, apenas separando elas pelo
	// operador ',' (virgula);
	// nessa linha, está sendo feita a declaração e inicialização de duas variáveis;
	// a variável 'x', recebe o valor '1';
	// a variável 'y', recebe o valor '2';
	// vale lembrar que a linguagem Go, avalia se a variável está sendo utilizada
	// dentro do código; caso a variável nã esteja sendo utilizada ele apresenta um
	// erro de compilação; basta remover o fmt.Println(x, y) para ter esse
	// comportamento;
	x, y := 1, 2
	fmt.Println(x, y)

	// utiliza a atribuição "simples", sem o operador ':=', troca o valor das variáveis;
	x, y = y, x
	fmt.Println(x, y)
}
