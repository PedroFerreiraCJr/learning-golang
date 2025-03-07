package main

import "fmt"

// a seguinte declaração, está definindo um tipo personalizado, que seria
// quase como um alias, fornecido por outras linguagens como C, e C++;
// dessa forma, é possível definir um método que tenha como receptor (receiver)
// esse tipo que acabamos de declarar;
type nota float64

// esse método, conforme mencionado no comentário acima, faz uso do tipo 'nota'
// para declarar um método que pode ser invocado neste tipo personalizado;
func (n nota) entre(inicio, fim float64) bool {
	// perceba que, neste trecho de código está sendo feita a conversão do alias
	// (neste caso, o tipo 'nota'), para um tipo float64;
	return float64(n) >= inicio && float64(n) <= fim
}

// com esta declaração de função é possível notar que, podemos fazer uso do tipo
// personalizado também como um tipo que é definido para um parâmetro de função,
// conforme o parâmetro 'n' do tipo 'nota';
// então, é possível passar para essa função um valor que seja um float64;
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
