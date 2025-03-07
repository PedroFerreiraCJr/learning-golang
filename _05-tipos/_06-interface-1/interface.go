package main

import "fmt"

// a seguinte declaração define uma interface, que na realidade é apenas um
// contrato, que estabele algumas assinaturas de método;
// neste exemplo, a interface é chamada de 'imprimivel';
// dentro dela existe uma única assinatura de função, como se fosse um protótipo
// de função, que é bastante conhecido nas linguagens C, e C++;
// o principal diferencial da linguagem Go, relacionado ao conceito de interfaces
// é que, em Golang, o próprio compilador verifica se uma determinada struct tem
// a assinatura de função que a interface solicita; desta forma, não é preciso
// informar, e nem é possível, dizer que uma struct implementa uma interface;
// para que haja compatibilidade entre uma struct e uma interface, basta que a
// struct tenha um método com a assinatura solicitada pela interface, conforme
// veremos logo abaixo;
type imprimivel interface {
	// define que esta interface declara uma função que tem a seguinte
	// assinatura: tem identificador 'toString', não declara nenhum parâmetro e
	// tem a responsabilidade de sempre retornar uma string;
	toString() string
}

// essa é uma declaração de struct, que na realidade é apenas uma estrutura de
// dados, que tem o papel de representar uma abstração, ou conceito do mundo real;
// neste caso, essa struct define uma representação de uma pessoa, onde neste
// contexto, a pessoa representada por essa struct tem dois atributos (campos);
type pessoa struct {
	nome      string
	sobrenome string
}

// essa é uma declaração de struct, que na realidade é apenas uma estrutura de
// dados, que tem o papel de representar uma abstração, ou conceito do mundo real;
// neste caso, essa struct define uma representação de um produto, onde neste
// contexto, o produto representado por essa struct tem dois atributos (campos);
type produto struct {
	nome  string
	preco float64
}

// IMPORTANTE: interfaces são implementadas implícitamente;

// essa é uma declaração de método, que tem como assinatura: uma função chamada de
// 'toString', que não declara nenhum parâmetro, e que deve sempre retornar uma
// string; mas, é preciso notar que está sendo feita a declaração de um receptor
// (receiver), que neste caso, é do tipo 'pessoa', e que dentro deste método pode
// ser referenciado pelo identificador 'p';
// outra coisa importante a se notar, é que neste caso, está sendo feita a
// declaração de um receptor por valor, ou seja, quando esse método for invocado
// em uma instância do tipo 'pessoa', esse objeto que está recebendo a invocação
// do método, será passado para dentro desta função utilizando a passagem de
// parâmetro para função por valor; e portanto, não poderá ser alterado dentro
// desta função;
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

// essa é apenas outra definição de método que está compatível com a assinatura da
// função declarada dentro da interface 'imprimivel'; mas neste caso, está sendo
// feita a declaração de um receptor do tipo 'produto';
func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// essa é uma declaração de função que espera receber um argumento do tipo
// 'imprimivel'; ou seja, essa função espera receber um objeto (instância), que
// tenha uma assinatura compatível com o que foi estabelecido na interface
// 'imprimivel'; desta forma, é possível passar como argumento para essa função
// qualquer objeto que possua essa assinatura de função, que é feita através de
// um método, com determinado receptor (receiver);
func imprimir(obj imprimivel) {
	fmt.Println(obj.toString())
}

func main() {
	// essa é uma declaração de variável que utiliza Polimorfismo;
	// isso porque ela tem o tipo de uma interface, e interfaces podem ter mais
	// de um tipo que seja compatível com ela, e portanto, é permitido a
	// atribuição;
	var coisa imprimivel = pessoa{"Pedro", "Junior"}
	fmt.Println(coisa.toString())

	// como a variável 'coisa' tem um tipo interface, ele tem a capacidade de poder
	// receber uma instância de uma struct que tenha a assinatura correta, conforme
	// o que foi preestabelecido na declaração da interface 'imprimivel';
	coisa = produto{"Caneta", 2.29}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// mesmo que a linguagem Go não seja Orientada a Objetos, ela possui alguns
	// dos melhores conceitos deste paradigma;
	p2 := produto{"Lápis", 1.59}
	imprimir(p2)
}
