package main

import "fmt"

// uma função em Go, dá suporte ao conceito de Closure;
// a função anônima declarada dentro da função chamada de 'closure', tem
// consciência de sua origem, e portanto, mantem uma 'memória' do que está
// no contexto da função 'closure';
func closure() func() {
	x := 10

	// função anônima, que está sendo retornada quando a função
	// 'closure' é invocada;
	// neste caso, a função anônima que está sendo retornada, está no
	// contexto da função 'closure';
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {
	x := 20

	fmt.Println(x)

	imprimeX := closure()

	// qual deveria ser o resultado esperado da invocação da função imprimeX?
	// será que ela deveria retornar o valor de 'x', no contexto da função 'main',
	// ou ela deveria retornar o valor de 'x', no contexto da função 'closure'?
	imprimeX()
}
