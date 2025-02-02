package main

import "fmt"

// essa função estabelece em sua assinatura que ela deve receber um 'float64'
// e deve retornar uma informações do tipo 'string';
func notaParaConceito(n float64) string {
	// utilização da instrução 'switch' com argumento omitido;
	// desta forma, esse se torna um 'switch' booleano;
	// ele avalia uma expressão, e executa a primeira que der verdadeiro;
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
