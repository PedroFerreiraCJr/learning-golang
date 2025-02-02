package main

import (
	"fmt"
	"time"
)

func main() {
	// declara e inicializa a variável 't', com o resultado da função
	// time.Now(); essa função retorna um objeto do tipo time.Time, que
	// representa a data/hora atual da máquina;
	t := time.Now()

	// esse é um 'switch', que avalia expressões booleanas;
	// então, ele executa o 'case' relacionado a primeira expressão
	// booleana que for retornar o valor 'true';
	// esse é um switch sem nenhuma argumento, portanto, ele procura
	// o primeiro case que dê um valor verdadeiro;
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Bom tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
