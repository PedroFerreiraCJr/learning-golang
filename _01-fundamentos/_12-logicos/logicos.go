package main

import "fmt"

// essa função, estabelece na sua definição, que é como se fosse um "contrato",
// que ela deve devolver 3 valores (isso mesmo, 3 valores); as funções em Golang
// podem retornar mais de um valor quando executadas;
// outra coisa importante nessa declaração de função, é que, ela define os parâmetros
// de entrada que ela deve receber para executar sua funcionalidade, seu conjunto de
// passos; essa função, nos seus parâmetros declara que recebe dos valores; os
// parâmetros foram declarados com uma sintaxe diferente; apenas especificamos dois
// nomes de variáveis (parâmetros são considerados variáveis), que tem o mesmo tipo;
// o tipo definido para os dois parâmetros é 'bool'; que define um tipo booleano;
// outra coisa interessante sobre essa função, é que ela declara que tem como
// resultado, 3 valores do tipo 'bool'; ou seja, essa função quando invocada, e
// quando executando suas operações, instruções, corretamente, ela retorna 3 valores
// do tipo 'bool';
func comprar(trabalho1, trabalho2 bool) (bool, bool, bool) {
	// caso ambas as variáveis envolvidas nesta operação tiverem o valor 'true',
	// a expressão terá o resultado 'true';
	comprarTv50 := trabalho1 && trabalho2

	// apenas se uma das variáveis envolvidas nesta operação tiver o valor 'true',
	// a expressão terá o resultado 'true';
	comprarTv32 := trabalho1 != trabalho2 // ou exclusivo;
	// se qualquer uma das variáveis envolvidas nesta operação tiver o valor 'true',
	// a expressão terá o resultado 'true';
	comprarSorvete := trabalho1 || trabalho2

	// nesta linha, a função 'comprar' está cumprindo o "contrato" estabelecido, e
	// retornando os 3 valores prometidos;
	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	// declara e inicializa 3 variáveis com uma única invocação de função;
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n",
		tv50, tv32, sorvete, !sorvete)

	// apenas atribui valores as variáveis, porque elas já foram declaradas;
	tv50, tv32, sorvete = comprar(false, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n",
		tv50, tv32, sorvete, !sorvete)

	// apenas atribui valores as variáveis, porque elas já foram declaradas;
	tv50, tv32, sorvete = comprar(true, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n",
		tv50, tv32, sorvete, !sorvete)

	// apenas atribui valores as variáveis, porque elas já foram declaradas;
	tv50, tv32, sorvete = comprar(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t\n",
		tv50, tv32, sorvete, !sorvete)
}
