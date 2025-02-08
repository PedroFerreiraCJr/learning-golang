package main

import "fmt"

func main() {
	// essa é uma declaração de map dentro de map;
	funcionariosPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	// remove uma chave de dentro do primeiro map;
	delete(funcionariosPorLetra, "P")

	for letra, funcionario := range funcionariosPorLetra {
		fmt.Println(letra, funcionario)
	}
}
