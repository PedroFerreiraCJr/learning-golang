package main

import "fmt"

func main() {
	// essa declaração de variável se torna um 'array', e não um 'slice', porque
	// a declaração de array tem um tamanho fixo informado no momento da declaração
	// da variável;
	array1 := [3]int{1, 2, 3}

	// slice de inteiros;
	var slice1 []int

	// esse código não compila porque não é possível appendar um novo valor dentro
	// de um array de tamanho fixo, porque eventualmente, seria necessário aumentar
	// o tamanho do array, e isso não é feito através da função 'append';
	//array1 = append(array1, 4, 5, 6)
	// a função 'append', adiciona elementos dentro de um slice;
	// se o slice, que é o primeiro argumento, já está com o tamanho máximo
	// do array interno, a função 'append' cria um outro arrau, com o dobro
	// do tamanho do array interno original, e copia para o array que foi criado
	// os valores no array original;
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	// a função 'copy', copia valores de determinado slice para outro slice;
	// neste caso, o slice de destino é a variável 'slice2'; enquanto que, a
	// variável 'slice1', é de onte os valores serão copiados;
	// neste caso, como o slice de destino, que foi criado através da função
	// 'make', tem um tamanho de 2, conforme o argumento informado para a
	// função 'make';
	// a função 'copy', não aumenta o tamanho do slice de destino;
	// outra aspecto importante da função 'copy', é que ela precisa receber
	// como argumentos, slices e não arrays;
	copy(slice2, slice1)
	fmt.Println(slice2)
}
