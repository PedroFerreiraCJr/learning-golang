package main

// aqui está sendo feito a importação de um pacote chamado "fmt";
// esse pacote é uma funcionalidade da própria biblioteca da linguagem Go;
// é possível atribuir um alias para um pacote que foi importado;
// para isso, simplesmente adicione o alias antes do nome do pacote;
// neste caso, está sendo atribuido o alias 'm' para o pacote 'math';
// desta forma, é possível se referir as funções e objetos dentro deste pacote
// através do alias informado;
// vale mencionar que, não pode ser feita a atribuição de alias e não utilizar,
// porque gera um erro de compilação, e o próprio compilador se encarrega de
// remover o valor não utilizado, deixando assim, somente o nome do pacote;
import (
	"fmt"
	m "math"
)

func main() {
	// a linha seguinte é uma declaração de uma constante em Go;
	// é preciso notar que, neste caso, está sendo feita a declaração
	// explícita do tipo da constante, que neste caso, está sendo feito
	// depois de declarar o nome da constante, com o tipo 'float64';
	// vale notar que, um literal numérico com decimal, ou seja, um valor
	// numérico definido com o uso de ponto (.), e um valor numérico é
	// considerado do tipo: 'float64';
	// a linguagem Go é estaticamente tipada e compilada;
	const PI float64 = 3.1415

	// a linha seguinte é uma declaração de uma variável em Go;
	// na linguagem Go, é emitido um erro de compilação caso seja feito a
	// declaração de uma variável e ela não tenha sido usada;
	// para declarar uma variável em Go, é possível usar a palavra-chave
	// 'var', conforme o seguinte exemplo;
	// essa linha está declarando uma variável e no mesmo momento
	// inicializando a variável;
	// neste caso, não é preciso informar o tipo da variável porque o
	// próprio compilador é capaz de inferir a partir do valor informado
	// no momento da inicialização da variável, que neste caso, está sendo
	// feita na mesma instrução;
	var raio = 3.2

	// na linha seguinte está sendo feita a impressão na saída padrão de
	// um valor textual, mais a concatenação do valor da variável 'raio';
	// é preciso falar que, automaticamente é adiciondo um espaço em branco
	// para separar o valor da variável informada do valor do literal string;
	fmt.Println("O valor da variável 'raio' é", raio)

	// essa é uma forma reduzida de de criar uma variável;
	// nessa forma reduzida, é utilizado um outro operador de atribuição, o
	// operador dois pontos igual (:=);
	// nessa instrução está sendo feita a declaração e, na mesma instrução,
	// também está sendo feita a inicialização da variável;
	// neste caso, não é preciso informar o tipo da variável porque o próprio
	// compilador consegue inferir o tipo dela a partir do valor atribuído;
	// como neste caso está sendo utilizado o operador ':=', que o papel é
	// declarar e inicializar variáveis; após essa linha, a variável 'area'
	// já está declarada, e portanto, está pronta para receber um eventual
	// novo valor; para fazer isso, não se pode usar o operador ':=' porque
	// a variável já foi declarada e inicializada; neste caso, para atualizar o
	// valor da variável é preciso usar o operador de atribuição '=';
	area := PI * m.Pow(raio, 2)

	// essa linha não compila;
	//area := 1

	// a linha seguinte faz a impressão de um valor textual no console da
	// aplicação; neste caso, está sendo feita a concatenação de um valor
	// literal string, com o valor da variável 'area';
	fmt.Println("A área da circunferência é", area)

	// essa linha é outra forma de declarar constantes na linguagem Go;
	const (
		a = 1
		b = 2
	)

	// essa linha é outra forma de declarar variáveis na linguagem Go;
	// para que não haja erro de compilação, é preciso utilizar, de alguma
	// maneira as variáveis declaradas;
	var (
		c = 3
		d = 4
	)

	// faz uso das variáveis declaradas acima;
	fmt.Println(a, b, c, d)

	// outra forma de declarar variáveis na linguagem Go;
	// nesta linha está sendo feita a declaração de duas variáveis do tipo
	// 'bool' (booleano), e está sendo feita a inicialização delas;
	// neste caso, está sendo atribuído para a variável 'e' o valor 'bool'
	// 'true'; e para a variável 'f' está sendo atribuído o valor 'bool' 'false';
	var e, f bool = true, false

	// faz uso das variáveis declaradas acima;
	fmt.Println(e, f)

	// outra forma de declarar variáveis na linguagem Go;
	// neste caso, está sendo feita a declaração e inicialização de 3 variáveis;
	// onde, a variável 'g' recebe o valor literal inteiro '2';
	// onde, a variável 'h' recebe o valor literal bool 'false';
	// onde, a variável 'i' recebe o valor literal string '"epa!"'
	g, h, i := 2, false, "epa!"

	// faz uso das variáveis declaradas acima;
	fmt.Println(g, h, i)
}
