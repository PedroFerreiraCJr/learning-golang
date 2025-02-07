package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array;
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// é possível ter um array, e vários slices apontando para o mesmo array;
	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array;
	// slice define um pedaço de um array;
	// define um slice, a partir do array 'a2', onde o primeiro índice começa em '1';
	// já o final do slice seguinte termina no índice 2, porque o valor informado foi '3';
	// desta forma, os índices contido no intervalo do slice, começa em '1', e vai até o índice
	// '3-1', portanto, o índice '2', está incluído, mas não o índice '3';
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// o slice não gera um array diferente;
	// o slice tem um ponteiro que aponta para o primeiro elemento, e um tamanho do slice;

	// a partir do array 'a2', define um slice, contendo os elementos do índice: 0, 1; e não
	// inclui o valor do índice 2;
	// novo slice, mas aponta para o mesmo array;
	s3 := a2[:2]
	fmt.Println(a2, s3)

	// você pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array;

	// também é possível ter um slice de um slice;
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
