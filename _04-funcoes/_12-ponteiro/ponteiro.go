package main

import "fmt"

// essa função tem um parâmetro do tipo 'int';
// portanto, quando essa função for invocada, ela precisa que seja
// informado um valor do tipo 'int';
// na linguagem Go, a passagem de parâmetros para função é feita, por
// padrão, através de passagem de parâmetro para função por valor;
// desta forma, o argumento de for fornecido no momento da invocação da
// função abaixo, será copiado para o parâmetro da função, que neste caso,
// é um 'int', chamado de 'n';
func inc1(n int) {
	n++
}

// essa função tem um parâmetro do tipo '*int';
// portanto, quando essa função for invocada, ela precisa que seja
// informado um endereço de um valor do tipo 'int';
// como esse parâmetro declarado é um ponteiro, ele irá receber o endereço
// de uma outra variável, e desta forma, será capaz de alterar o valor que
// está contido na variável que foi passada para esta função como argumento;
func inc2(n *int) {
	// como o parâmetro desta função foi declarado como sendo um ponteiro, para
	// que seja possível alterar o valor que está armazenado no endereço que
	// o ponteiro de inteiro 'n' aponta, é preciso desreferenciar o ponteiro 'n';
	// para que o compilador saiba que estamos querendo alterar o valor que está
	// contido dentro do endereço de memória que foi armazenado no ponteiro 'n';
	*n++
}

func main() {
	numero := 6

	// essa é uma passagem por valor; recebe uma cópia do valor de 'n';
	inc1(numero)
	fmt.Println(numero)

	// essa é uma passagem por referência; recebe o endereço da
	// variável 'numero';
	inc2(&numero)
	fmt.Println(numero)
}
