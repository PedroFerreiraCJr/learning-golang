package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	// aqui está sendo feito o uso de uma variável do mesmo tipo do objeto
	// instânciado, que neste exemplo é do tipo 'ferrari';
	// portanto, não é preciso ter um ponteiro para poder alterar o valor
	// do atributo 'turboLigado';
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// o ponto principal desta aula está neste trecho de código;
	// como a implementação do método da interface 'esportivo', no caso da struct
	// ferrari foi feito através de um receptor do tipo ponteiro, é necessário
	// utilizar o operador '&' (e-comercial) para que o código funcione corretamente
	// porque neste exemplo, a referência é polimórfica (está sendo atribuído o
	// tipo 'ferrari' para a interface 'esportivo'); mas isso somente é necessário
	// caso seja preciso alterar um atributo de uma struct através de uma referência
	// polimórfica;
	// outra coisa importante, é que caso o e-comercial seja removido, o código
	// passa a não mais compilar;
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
