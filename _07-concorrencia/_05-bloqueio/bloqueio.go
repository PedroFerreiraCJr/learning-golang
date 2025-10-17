package main

import (
	"fmt"
	"time"
)

// a função abaixo declara um parâmetro do tipo 'chan' de 'int', um canal em que é possível
// trafegar valores do tipo de dado 'int';
func rotina(ch chan int) {
	time.Sleep(time.Second) // faz com que a goroutine seja suspensa por 1s;
	ch <- 1                 // envia o valor literal inteiro '1' para o canal 'ch';
	// é preciso mencionar o comportamento que é inerente ao código acima, e, para este caso,
	// a goroutine da função 'main';
	// quando esta função é iniciada através de uma outra goroutine a partir da função
	// 'main', uma nova linha de execução começa a rodar esta função de maneira
	// independente da linha de execução da goroutine da função 'main'; imediatamente após
	// iniciar a goroutine para esta função, a goroutine da função 'main' prossegue para
	// dar continuidade a execução do restante de código da função 'main';
	// então, a função 'rotina' envia uma informação através do canal que foi instanciado
	// e passado como argumento para esta função;
	// neste momento, a goroutine da função 'rotina' fica bloqueada, esperando que o valor
	// que foi enviado através do canal que foi fornecido como argumento para ela;

	fmt.Println("somente após a leitura do valor enviado para o canal, esta linha será executada")
}

func main() {
	// a função 'make' consegue criar uma instância do tipo embutido 'chan';
	// neste caso, o canal instanciado tem o tipo informado como 'int';
	// é importante saber que, é possível criar um canal que tenha um buffer;
	// na invocação da função 'make' abaixo, não está sendo informado um valor inteiro
	// relacionado ao tamanho do buffer do canal, portanto, o canal criado neste caso é
	// um canal sem buffer;
	// quanto a questão do buffer, é preciso saber que quando existe um buffer envolvido
	// na instanciação do canal, os valores podem ser enviado para o canal sem que seja
	// feito um bloqueio da possível goroutine que escreve neste canal;
	// já quando houver esgotado o buffer, ou seja, o buffer estiver cheio, uma possível
	// operação de escrita neste mesmo canal acarreta um bloqueio na escrita;
	// portanto, para que seja possível escrever mais valores neste canal, é preciso que
	// seja feita a leitura de valores do buffer, liberam espaço para uma posterior operação
	// de escrita no buffer do canal;
	ch := make(chan int)
	go rotina(ch) // roda esta função em uma nova goroutine;

	fmt.Println("esperando para ler o valor enviado para o canal 'ch'")
	fmt.Println(<-ch) // operação bloqueante;
	fmt.Println("o valor do canal 'ch' acabou de ser lido")
	fmt.Println(<-ch) // o runtime da linguagem encontra um cenário de deadlock!
	fmt.Println("nunca serei executado - por causa de um deadlock!")
}
