package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// neste caso, é preciso fazer a conversão explícita para tipos iguais;
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado...
	// neste caso, para que seja possível realizar a concatenação, é preciso converter
	// o valor literal numérico '123' para uma string;
	// mas neste caso, o que está acontecendo é que o valor informado está sendo convertido
	// para uma string utilizando a tabela ASCI;
	fmt.Println("Teste " + string(97))

	// int para string;
	// neste caso, está sendo feita a conversão do valor numérico '123' para uma string;
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int;
	// na linguagem Go é possível retornar mais de um valor como resultado da execução
	// de uma função; o exemplo abaixo mostra que a função strconv está retornando mais de
	// um valor; a variável 'num' recebe o valor primário do resultado da invocação da
	// função strconv, mas caso dê algum erro na conversão do argumento da função para o
	// valor inteiro, a segunda variável conterá o erro que ocorreu;
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("O valor convertiodo é true")
	}
}
