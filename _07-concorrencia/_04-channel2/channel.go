package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines;
// é importante saber que um 'chan', em Golang, é um tipo nativo (embutido) da linguagem;

// essa função declara um parâmetro do tipo 'chan' de 'int';
// o parâmetro 'ch' é um canal de comunicação entre duas ou mais goroutines, onde para este canal, foi
// definido que nele é esperado que seja trafegado apenas valores do tipo de dado 'int';
func doisTresQuatroVezes(base int, ch chan int) {
	time.Sleep(time.Second) // faz com que a goroutine seja suspensa por 1s;
	ch <- 2 * base          // enviando informações para o canal 'ch';

	time.Sleep(time.Second) // faz com que a goroutine seja suspensa por 1s;
	ch <- 3 * base          // enviando informações para o canal 'ch';

	time.Sleep(3 * time.Second) // faz com que a goroutine seja suspensa por 3s;
	ch <- 4 * base              // enviando informações para o canal 'ch';
}

func main() {
	ch := make(chan int)          // instancia um channel (canal) que é usado para comunicação entre goroutines;
	go doisTresQuatroVezes(2, ch) // inicia uma goroutine responsável por executar a função informada;

	a, b := <-ch, <-ch // declara duas variáveis, e lê do canal 'ch' dois valores de maneira consecutiva;
	fmt.Println(a, b)

	fmt.Println(<-ch) // lê um novo valor do canel 'ch';

	// é importante mencionar que caso seja feita uma tentativa de ler outro valor do canal 'ch'
	// o que irá acontecer neste cenário é um erro em tempo de execução;
	// este erro é porque a goroutine da função principal não tem como ler um novo valor do canal
	// porque a goroutine que envia informações para o canal já foi concluída, e não tem outra
	// enviando informações para o canal 'ch'; e isso gera um deadlock, situação conhecida na
	// programação concorrente onde duas threads não conseguem fazer progresso, e permanecem desta
	// forma sem que hava como realizar um processamento;
}
