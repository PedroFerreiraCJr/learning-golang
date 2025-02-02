package main

import "fmt"

func main() {
	// declara e inicializa uma variável do tipo 'int'; o compilador considera um
	// literal numérico, sem casa decimal, como sendo do tipo de dado 'int';
	i := 1

	// IMPORTANTE: a linguagem Go não tem aritmética de ponteiros;
	// declara uma variável que é um ponteiro;
	// a variável 'ponteiro' foi inicializada com o valor literal 'nil';
	// o seguinte ponteiro aponta para "lugar nenhum";
	// para declarar um ponteiro, é preciso informar o tipo do ponteiro, que neste
	// caso, é um ponteiro para 'int'; uma coisa importante, é que é preciso usar
	// o operador '*' (asterísco), para informar ao compilador que esta variável é
	// um ponteiro; e depois o tipo que ele pode armazenar, que neste exemplo, é
	// o tipo 'int'; desta forma, temos uma variável que é um ponteiro para 'int';
	var ponteiro *int = nil

	// para obter o endereço de uma variável, devemos utilizar o operador '&' (E comercial); nesse exemplo, estamos obtendo o endereço da variável 'i', e
	// atribuindo a uma variável que é um ponteiro, chamada neste caso, 'ponteiro';
	ponteiro = &i // pegando o endereço da variável 'i';

	// desta forma, se torna possível alterar o valor da variável apontada pelo
	// ponteiro, onde neste caso, o ponteiro se chama 'ponteiro', indiretamente,
	// somente utilizando a variável ponteiro;

	// para obter o valor da variável que está sendo referenciada pela variável
	// 'ponteiro', basta utilizar o operador '*' (asterísco), novamente;
	fmt.Println("O valor da variável 'i', acessada diretamente é:", i)
	fmt.Println("O valor da variável 'i', acessada através do ponteiro é:", *ponteiro)

	// colocando somente o nome da variável que é um ponteiro, neste caso, a variável
	// chamada 'ponteiro', é possível obter o endereço referenciado pelo ponteiro;

	// desreferencia o ponteiro e incrementa o valor contido no endereço de memória
	// que está contido dentro da variável 'ponteiro';
	*ponteiro++ // desreferenciado (pegando o valor);

	fmt.Println("O valor da variável 'i', após o incremento, acessada diretamente é:", i)
	fmt.Println("O valor da variável 'i', após o incremento, acessada através do ponteiro é:", *ponteiro)
	fmt.Println(ponteiro, *ponteiro, i, &i)
}
