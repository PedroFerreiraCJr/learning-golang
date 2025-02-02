package main

import (
	"fmt"
	"math/rand"
	"time"
)

// declara uma função que retorna, aleatóriamente, um valor inteiro no
// intervalo de 0 a 9;
func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// essa linha é a parte principal desta aula;
	// a linha abaixo, está fazendo o seguinte:
	// antes de entrar no bloco do 'if', é declarada uma variável chamada 'i';
	// essa variável é declarada e inicializada (conforme já sabemos), com o
	// uso do operador ':=' (dois-pontos igual);
	// desta forma, antes de saber se essa instrução 'if' deve ser executada,
	// primeiro, é feito a declaração e inicialização da variável 'i', de acordo
	// com o valor retornado pela função 'numeroAleatorio()';
	// depois disso, é colocado o operador ';' (ponto-e-vírgula), para separar
	// a declaração e inicialização da variável, da expressão que avalia se
	// é possível executar o bloco da instrução 'if';
	// neste caso, devemos interpretar a seguinte instrução como:
	// inicializa a variável 'i' com o valor retornado pela função
	// 'numeroAleatorio()'; se o valor da variável 'i' for maior que '5', então
	// execute o bloco da instrução 'if';
	// a variável 'i', fica visível somente dentro do bloco;
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("O valor da variável 'i' é:", i)
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("O valor da variável 'i' é:", i)
		fmt.Println("Perdeu!")
	}
}
