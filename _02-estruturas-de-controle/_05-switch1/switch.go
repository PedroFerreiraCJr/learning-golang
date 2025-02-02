package main

import "fmt"

// essa função estabelece na sua assinatura, seu contrato, que ela deve
// receber como argumento um valor do tipo 'float64', e deve devolver um
// valor do tipo 'string';
func notaParaConceito(n float64) string {
	// essa linha, converte o valor do parâmetro 'n' para um valor do tipo
	// inteiro (int);
	var nota = int(n)

	// essa linha, avalia o valor da variável 'nota', para saber se ela
	// atende a determinado valor de comparação;
	// o argumento da instrução 'switch', não é um valor booleano, como em
	// outras instruções da linguagem, como a expressão de um 'if', ou a
	// expressão avaliada no caso do laço 'for';
	// caso haja uma correspondência do valor do argumento do switch
	// com o valor deste 'case', então ele executa o 'case' correspondente;
	switch nota {
	// caso o valor da variável 'nota', for igual ao valor '10', então,
	// esse 'case' será executado;
	case 10:
		// na linguagem Go, diferente de outras linguagens C-like, não
		// precisamos utilizar a palavra-chave 'break' para concluir a
		// execução de determinado 'case', quando há uma correspondência
		// de valores; na linguagem Go, o comportamento padrão quando há uma
		// correspondência e a seleção de um determinado 'case' da instrução
		// 'switch', após as instruções do 'case' serem executadas,
		// automaticamente, todo o restante de cases é ignorado; e portanto,
		// não serão executados;
		// já com o uso da palavra-chave 'fallthrough', o contrário acontece;
		// desta forma, caso exista dentro de determinado 'case' a instrução
		// 'fallthrough', o 'case' logo a baixo será executado também;
		fallthrough
	case 9:
		return "A"
		// esse 'case', estabelece que se o valor do argumento do 'switch' for
		// igual a 8 ou 7, então esse 'case' será executado;
	case 8, 7:
		return "B"
		// esse 'case', estabelece que se o valor do argumento do 'switch' for
		// igual a 6 ou 5, então esse 'case' será executado;
	case 6, 5:
		return "C"
		// esse 'case', estabelece que se o valor do argumento do 'switch' for
		// igual a 4 ou 3, então esse 'case' será executado;
	case 4, 3:
		return "D"
		// esse 'case', estabelece que se o valor do argumento do 'switch' for
		// igual a 2 ou 1 ou 0, então esse 'case' será executado;
	case 2, 1, 0:
		return "E"
		// esse é o caso default, ou seja, se não houver nenhuma correspondência
		// nos cases anteriores, então essa instrução será executada;
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
