package main

import "fmt"

func main() {
	// essa é a forma de inicializar literal um map;
	// o último elemento deve conter uma vírgula;
	funcioanriosESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	// adiciona um novo elemento ao map de funcionários e salários;
	funcioanriosESalarios["Rafael Filho"] = 1350.0

	// faz uma tentativa de excluir um elemento, mas ele não existe então
	// não acontece nada, pois não existe chave para excluir;
	delete(funcioanriosESalarios, "Inexistente")

	fmt.Println("----------------------------")

	// uma forma de percorrer o map;
	for nome, salario := range funcioanriosESalarios {
		fmt.Println(nome, salario)
	}

	fmt.Println("----------------------------")

	// outra forma de percorrer o map;
	// imprime somente o valor relacionado a chave;
	// a chave foi omitida, pois utilizamos o operador '_' (underscore);
	for _, salario := range funcioanriosESalarios {
		fmt.Println(salario)
	}

	fmt.Println("----------------------------")

	// outra forma de percorrer o map;
	// obtem somente o valor da chave
	for nome := range funcioanriosESalarios {
		fmt.Println(funcioanriosESalarios[nome])
	}
}
