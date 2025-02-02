package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516
	// a linha seguinte não compila; não dá para realizar a concanetação, usando o operador
	// mais (+), entre um valor literal string e um valor numérico;
	// a linha abaixo está comentada pois gera um erro de compilação, dado a explicação anterior;
	//fmt.Println("O valor de x é" + x)

	// a função 'Sprint', converte um valor numérico para string;
	xs := fmt.Sprint(x)

	// a concatenação entre valores que são strings, mesmo envolvendo variáveis é correta;
	fmt.Println("O valor de x é" + xs)

	// outra alternativa
	fmt.Println("O valor de x é", x)

	// imprime uma variável do tipo 'float'; podendo formatar a quantidade de casas decimais;
	fmt.Printf("O valor de x é: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
