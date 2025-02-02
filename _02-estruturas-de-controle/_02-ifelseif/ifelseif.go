package main

import "fmt"

// essa é uma declaração de função;
// ela recebe como parâmetro um valor decimal, do tipo 'float64';
// ela, também estabelece que deve retornar um resultado do tipo 'string';
// portanto, todas as ramificações de execução desta função tem que retornar
// um valor do tipo 'string';
func notaParaConceito(n float64) string {
	// estabelece um critério para devolver a string "A";
	// neste caso, está sendo estabelecido um intervalo de valor aceitável
	// para cada ramificação; se o valor do parâmetro informado estiver em
	// determinado intervalo, então um valor do tipo 'string', corespondente,
	// será retornado pela função;
	if n >= 9 && n <= 10 {
		// retorna a string "A", quando a nota estiver entre 9 e 10;
		return "A"
	} else if n >= 8 && n < 9 {
		// retorna a string "A", quando a nota estiver entre 8 e menor que 9;
		return "B"
	} else if n >= 5 && n < 8 {
		// retorna a string "A", quando a nota estiver entre 5 e menor que 8;
		return "C"
	} else if n >= 3 && n < 5 {
		// retorna a string "A", quando a nota estiver entre 3 e menor que 5;
		return "D"
	} else {
		// caso o valor da nota informado no parâmetro for menor que 3;
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
