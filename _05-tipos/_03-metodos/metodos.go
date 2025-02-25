package main

import (
	"fmt"
	"strings"
)

// declaração de uma struct, uma estrutura de dados, para representar uma
// pessoa, que neste contexto, terá somente os atributos 'nome', e 'sobrenome';
type pessoa struct {
	nome      string
	sobrenome string
}

// método que retorna o valor do nome completo, de acordo com um receptor do
// tipo 'pessoa';
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// nesta definição de "método" da struct pessoa, tem um ponto importante;
// neste caso, está sendo informado que a struct do tipo pessoa, deve ser passada
// para o "método" por REFERêNCIA; desta forma, será possível alterar o valor
// contido dentro da struct pessoa, que for o receiver (receptor) corrente na
// invocação do "método" setNomeCompleto; caso não tivesse sido informado o operador
// '.' (ponto), no receiver da função 'setNomeCompleto', a instância corrente onde
// está sendo invocado o método 'setNomeCompleto', seria passado por valor; e desta
// forma, os atributos da struct pessoa, não seriam alterados; para ter esse
// comportamento, basta remover o operador '*' (asterísco) da declaração do
// método 'setNomeCompleto', que foi definido logo abaixo;
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())
}
