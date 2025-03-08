package main

import "fmt"

// define uma interface, que tem apenas uma única assinatura de método;
type Esportivo interface {
	ligarTurbo()
}

// define uma interface, que tem apenas uma única assinatura de método;
type Luxuoso interface {
	fazerBaliza()
}

// define uma interface, que é composta das interfaces anteriores, 'Esportivo'
// e 'Luxuoso';
type EsportivoLuxuoso interface {
	Esportivo
	Luxuoso
	// podemos adicionar mais método aqui, sem problemas;
}

// define uma estrutura de dados que não possui nenhum atributo, mas é usada
// para implementar as interfaces 'Esportivo', 'Luxuoso', e 'EsportivoLuxuoso';
type Bmw7 struct{}

// faz com que a struct 'Bmw7', implemente o método da interface 'Esportivo';
func (b Bmw7) ligarTurbo() {
	fmt.Println("ligando turbo...")
}

// faz com que a struct 'Bmw7', implemente o método da interface 'Luxuoso';
func (b Bmw7) fazerBaliza() {
	fmt.Println("fazendo baliza...")
}

// como a struct 'Bmw7' tem todas as assinaturas de função estabelecidas em
// todas as 3 interfaces, ela pode ser usada de maneira polimorfica, ou seja,
// onde é esperado uma instância de 'Esportivo', posso utilizar uma instância
// da struct 'Bmw7';
// assim como, onde é esperado uma instância de 'Luxuoso', posso utilizar uma
// instância da struct 'Bmw7';
// da mesma forma, onde é esperado uma instância de 'EsportivoLuxuoso', posso
// utilizar uma instância da struct 'Bmw7';
// isso somente é possível porque a struct 'Bmw7' implementa todas as assinaturas
// de método preestabelecidas em todos os "contratos" (tipos interfaces);

func main() {
	// declara uma variável do tipo 'EsportivoLuxuoso', que neste caso, é um
	// tipo interface;
	// como a struct 'Bmw7' satisfaz as assinaturas de método de ambas as
	// interfaces 'Esportivo' e 'Luxuoso', ela também satisfaz a interface
	// 'EsportivoLuxuoso' de maneira implícita também;
	var b EsportivoLuxuoso = Bmw7{}

	b.ligarTurbo()
	b.fazerBaliza()
}
