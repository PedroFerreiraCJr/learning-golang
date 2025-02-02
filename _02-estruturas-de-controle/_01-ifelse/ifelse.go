package main

import "fmt"

// essa linha declara uma função que requer um parâmetro de entrada
// do tipo 'float64';
func imprimirResultado(nota float64) {
	// essa é a sintaxe solicitada pela linguagem Go;
	// não devemos colocar a expressão booleana dentro de parênteses;
	// e mesmo que se tenha somente uma única instrução para ser executada,
	// é preciso que seja fornecido um par de chaves;
	// um aspecto sobre a utilização de parênteses é que, eles são opcionais, e
	// algumas IDE's, como o VSCode, tem uma configuração de automaticamente
	// remover os parênteses, caso ele seja desnecessário;
	// outra coisa importante, é que isso também vale para o operador
	// ';' (ponto-e-vírgula); eles são removidos automaticamente pela
	// própria IDE;
	// a branch, ramificação, definida pelo bloco 'else', é opcional;
	// abaixo, um exemplo simples, que demonstra que os parênteses podem
	// ser utilizados; mas que, podem ser utilizado, geralmente, quando
	// for preciso alterar a ordem de precedência de avaliação da
	// expressão, então ele pode ser usado, normalmente;
	if nota >= 7 && (true || false) {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	// invoca a função 'imprimirResultado', passando como argumento o valor
	// literal decimal, '7.3';
	imprimirResultado(7.3)

	// invoca a função 'imprimirResultado', passando como argumento o valor
	// literal decimal, '5.1';
	imprimirResultado(5.1)
}
