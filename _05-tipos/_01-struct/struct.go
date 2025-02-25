package main

import "fmt"

// essa é a maneira de se definir um novo tipo de dados na linguagem Go;
// uma struct é um conjunto de campos; todo campo tem um tipo de dado;
// conforme como quando queremos declarar uma variável, basta definir o
// identificador do campo, assim como, o tipo de dado que o campo pode
// armazenar; não é permitido colocar ';' (ponto-e-vírgula); mas é possível
// definir mais de um atributo do mesmo tipo, na mesma linha, apenas separando
// pelo operador ',' (vírgula);
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: função com receiver (recetor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	// declaração de variável para ser utilizada logo em seguida;
	// mas poderia ser declarada conforme o exemplo do 'produto2';
	var produto1 produto

	// essa é a maneira de informar na linguagem Go que queremos uma instância
	// da struct produto; neste caso, foi utilizado uma maneira de inicializar a
	// struct, que nomeia os atributos que estão sendo informados;
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05, // essa vírgula é obrigatória;
	}

	// através do operador '.' (ponto), é possível acessar atributos da struct, e
	// também invocar um "método", que tenha previamente sido declarado para a
	// struct produto;
	fmt.Println(produto1, produto1.precoComDesconto())

	// essa é outra alternativa para a inicialização da struct;
	// primeiramente, está sendo utilizada a forma de declaração de variável de
	// maneira sucinta; além disso, a forma como passamos os valores iniciais para
	// a struct, é através da maneira "ordenada", de acordo com a ordem de definição
	// dos atributos da struct produto;
	produto2 := produto{"Notebook", 1989.90, 0.10}

	// invoca o "método", na instância da struct 'produto2';
	fmt.Println(produto2.precoComDesconto())
}
