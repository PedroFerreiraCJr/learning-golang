package main

import "fmt"

func main() {
	// como a struct 'Ponto' foi declarada com a primeira letra maiúscula
	// é possível acessar ela em qualquer pacote;
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	// como a função 'catetos' foi declarada com a primeira letra minúscula
	// e este arquivo de código-fonte foi declarado no mesmo pacote, então
	// é possível acessar a função declarada;
	// mas caso estivêssemos em um outro arquivo mas declarado dentro de
	// outro pacote, como a função foi declarada com a primeira letra minúscula
	// não seria possível acessar ela, porque ela estaria privada ao pacote onde
	// ela foi declarada, e estaríamos dentro de outro pacote, e ela não estaria
	// visível;
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
