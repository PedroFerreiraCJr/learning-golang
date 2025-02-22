package main

import "fmt"

// essa é uma função variática;
// declara uma função que pode receber um número variádo de argumentos;
// o operador para declarar um parâmetro como sendo um argumento variável é o
// operador '...' (reticências), que é semelhante a outras linguagens, como
// por exemplo, Java, JavaScript, C, C++;
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9, 9.9))
}
