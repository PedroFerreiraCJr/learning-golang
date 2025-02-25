package main

import "fmt"

// declara uma struct para representar o item de um pedido;
type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

// declara uma struct para representar um pedido de compra, que é composto por
// outra struct;
type pedido struct {
	userID int
	itens  []item
}

// declara um método na struct pedido, que deve ser utilizado para calcular o valor
// total de determinado pedido;
func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	// declaração de variável de maneira sucinta; declara e já inicializa ela
	// na mesma instrução, além disso, utilizando a inferência de tipos;
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		}, // essa ',' (vírgula), é obrigatória;
	}

	fmt.Printf("valor total do pedido é R$ %.2f", pedido.valorTotal())
}
