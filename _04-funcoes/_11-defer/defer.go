package main

import "fmt"

func obterValorAprovado(num int) int {
	// essa é a instrução apresentada nesta aula;
	// ela posterga a execução do 'println' para o momento antes de
	// retornar o resultado desta função;
	// essa instrução é executada independente de onde ocorre o
	// retorno da função;
	defer fmt.Println("fim!")

	if num > 5_000 {
		fmt.Println("valor alto...")
		return 5_000
	}

	fmt.Println("valor baixo...")
	return num
}

func main() {
	fmt.Println(obterValorAprovado(6_000))
	fmt.Println(obterValorAprovado(3_000))
}
