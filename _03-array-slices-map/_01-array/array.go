package main

import "fmt"

func main() {
	// arrays são estruturas de dados homogêneas (tem o mesmo tipo), e são estáticas (tem tamanho fixo);
	// se você define um array de 'int', ele só pode ter inteiros;

	// essa é uma declaração de um array contendo 3 elementos do tipo de dado 'int';
	// essa linha declara e já inicializa um array contendo 3 elementos do tipo de dado 'int', que
	// são, automaticamente, inicializados para o valor default, que para valores do tipo 'int', é '0';
	var notas [3]float64
	fmt.Println(notas)

	// essa linha, está fazendo a inicialização de cada uma das 3 posições do array, que foi declarado
	// logo acima;
	// o array são estruturas indexadas, portanto, podemos usar um índice para acessar cada um
	// dos valores do array;
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	// é importante mencionar que o compilador da linguagem Go, faz a verificação dos limites de um
	// array, em tempo de compilação; portanto, tentar acessar um índice inválido não compila, conforme
	// o seguinte exemplo;
	//notas[3] = 10

	total := 0.0
	// a função 'len' retorna o tamanho de seu argumento, que deve ser um array;
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	// calcula a média das notas
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f", media)
}
