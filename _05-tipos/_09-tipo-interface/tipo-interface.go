package main

import "fmt"

type Curso struct {
	Nome string
}

func main() {
	// um exemplo sem nexo, kkkkk;
	// apenas para demonstrar que é possível registrar uma função que deve ser
	// invocada quando esta função estiver para terminar;
	// neste caso, está sendo passado uma função anônima, que é automaticamente
	// invocada quando definida; pois a palavra-chave 'defer' precisa disto;
	defer (func() {
		fmt.Println("O programa terminou!")
	})()

	// declara uma variável que é do tipo, que podemos chamda de "interface vazia";
	// essa sintaxe parece um pouco estranha, mas está correta;
	// é preciso colocar o símbolo de chaves '{}', após a palavra-chave 'interface';
	// a variável se chama 'coisa';
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = 14.6
	fmt.Println(coisa)

	type Pessoa struct {
		Nome      string
		Sobrenome string
	}

	// armazenamento do endereço de uma instância do tipo 'Pessoa', dentro da
	// variável 'coisa';
	coisa = &Pessoa{Nome: "Pedro", Sobrenome: "Ferreira"}
	fmt.Println(coisa)

	// armazenamento de uma instância do tipo 'Pessoa', dentro da variável
	// 'coisa';
	coisa = Pessoa{Nome: "Pedro", Sobrenome: "Ferreira"}
	fmt.Println(coisa)

	// os exemplos acima, demonstram que é possível armazenar qualquer valor
	// dentro de uma variável do tipo "interface vazia"; ou tipo genérico;
	// isso também se assemelha ao tipo 'Object' da linguagem Java, onde ele
	// é um supertipo de todos os tipos objeto;

	// declara um tipo personalizado, onde o alias para o novo tipo se chama
	// 'Dinamico', e o tipo base, é uma "interface vazia";
	// isso é como se fosse um ponteiro genérico nas linguagens de programação
	// C, e C++; nestas linguagens, após a memória ter sido alocado para armazenar
	// um valor, é possível atribuir este endereço para um ponteiro genérico (void*)
	// de forma que, podemos, por exemplo, passar qualquer valor para dentro de uma
	// função que declare um parâmetro do tipo ponteiro genérico; ou até mesmo
	// retornar qualquer valor de uma função, utilizando para isso, um ponteiro
	// genérico (void*);
	// em Golang, o tipo 'interface{}' define um tipo de dado, seja para parâmetro de
	// função, ou para um tipo de retorno, ou até mesmo em uma declaração de variável
	// onde, podemos armazenar qualquer valor dentro dela;
	type Dinamico interface{}

	var coisa2 Dinamico
	coisa2 = "Hello World in Golang!"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = Curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
