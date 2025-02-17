package main

import "fmt"

// declara uma função que define em sua assinatura que ela deve receber
// 2 valores do tipo 'int', e que ela irá retornar a multiplicação de
// 'a' por 'b';
func multiplicacao(a, b int) int {
	return a * b
}

// declara uma função que tem, como primeiro parâmetro, uma declaração de
// assinatura de função; utilizando a palavra-chave 'func', é possível notar que
// esta função solicita que seja informado como argumento uma outra função;
// neste caso, além da função esperada como primeiro argumento, esta função também
// espera receber 2 outros valores do tipo 'int';
// essa função também declara que irá retornar um valor do tipo 'int';
// é preciso notar que, o primeiro parâmetro declarado na função 'exec', é
// uma outra função; e ela tem uma assinatura em específico;
// neste caso, a função que 'exec' espera receber deve ter como parâmetro 2
// valores do tipo 'int', e além disso, essa função esperada deve ter como
// resultado um valor do tipo 'int';
func exec(fn func(int, int) int, param1, param2 int) int {
	// para invocar a função que foi declarada como parâmetro desta funão,
	// basta utilizar o operador parênteses '()';
	// desta forma, assim como qualquer outra declaração de função, é preciso
	// informar os argumentos que a função espera receber;
	// neste caso, a função 'fn' espera receber como argumento 2 valores do
	// tipo 'int';
	return fn(param1, param2)
}

func main() {
	// declara uma variável que deve receber como valor inicial o que for
	// retornado pela invocação da função 'exec';
	// a função 'exec', por sua vez, deve receber uma func, com determinada
	// assinatura;
	// além disto, é preciso informar mais dois valores do tipo 'int';
	// portanto, neste exemplo, está sendo feita a invocação da função 'exec',
	// recebendo como argumento os valores 'multiplicacao', '3', '4';

	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
