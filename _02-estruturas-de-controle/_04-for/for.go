package main

import (
	"fmt"
	"time"
)

func main() {
	// declara e inicializa a variável 'i', para que ela sirva de contador
	// na instrução do laço 'for', definido logo abaixo;
	i := 1

	// exemplo 1: 'for', semelhante ao laço 'while';
	// a linguagem Go não possui 'while';
	// o laço 'for' definido logo abaixo, estabelece a condição de parada
	// do laço, na seguinte expressão lógica (expressão que retorna verdadeiro
	// ou falso);
	for i <= 10 {
		fmt.Println(i)

		// incrementa o valor da variável 'i', para que a condição do laço
		// 'for', possa ter, eventualmente, o valor 'falso' retornado, e então
		// o laço possa terminar de ser executado;
		// a linguagem Go não possui operador de incremento/decremento
		// prefixado; a linguagem somente possui estes operadores em
		// sua notação posfixada;
		i++
	}

	// exemplo 2: 'for', semelhante ao laço 'for' tradicional da linguagem C;
	// a seguinte declaração de laço 'for', é uma outra possibilidade da
	// linguagem; ele estabelece 3 partes consituintes; a primeira parte
	// é a declaração e inicialização da variável 'i', que foi inicializada
	// com o valor '0';
	// depois, na segunda parte, temos a expressão condicional que informa até
	// quando esse laço deve executar, que deve ser um valor booleano;
	// logo em seguida, temos a parte de atualização da variável envolvida na
	// expressão condicional, para que, eventualmente, se possa ter o valor
	// falso retornado pela expressão, e portanto, o laço seja interrompido;
	for i := 0; i <= 20; i += 2 {
		// imprime o valor da variável 'i';
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// exemplo 3: misturando 'for' com 'if';
	for i := 1; i <= 10; i++ {
		// essa expressão verifica se o resto da divisão inteira do valor da
		// variável 'i' pelo literal inteiro '2' retorna um valor igual a '0';
		if i%2 == 0 {
			// caso o resultado da expressão seja 'true', essa impressão será
			// realizada;
			fmt.Println("Par ")
		} else {
			// caso o resultado da expressão seja 'false', essa impressão será
			// realizada;
			fmt.Println("Ímpar ")
		}
	}

	// exemplo 4: laço 'for' infinito;
	for {
		// imprime uma menssagem no console;
		fmt.Println("Para sempre...")

		// "dorme" um segundo, antes de voltar para o início do
		// bloco deste 'for';
		time.Sleep(time.Second)
	}

	// veremos o foreach no capítulo sobre array;
}
