package main

import "fmt"

func main() {
	// um channel em Golang é uma forma de estabelecer um meio de comunicação entre duas ou
	// mais goroutines;
	// um canal (channel), que em Golang é definido através do tipo embutido 'chan', precisa
	// ter um determinado tipo de dado vinculado a ele; onde isso, estabelece qual o tipo de
	// informação (o tipo do dado), que pode ser transferido através deste canal, entre
	// os involvidos na comunicação, geralmente, sendo uma goroutine que transfere uma
	// informação para outa goroutine, desta forma, elas se comunicam;
	// para obter uma instância de um channel, basta fazer uso da função 'make'; onde é
	// necessário informar o tipo do dado capaz de ser enviado e lido do canal;
	ch := make(chan int, 1)

	// neste trecho de código, está sendo feito o envio de uma informação do tipo de
	// dado 'int' através do canal 'ch';
	ch <- 1 // enviando dados para o canal (escrevendo no canal 'ch' o valor literal '1');

	// nesta instrução está sendo feito a leitura de um eventual valor do canal 'ch';
	// como o valor literal '1' foi enviado anteriormente para o canal 'ch', então o valor
	// será lido; mas repare que o valor que está sendo lido do canal 'ch' não está sendo
	// atribuído a variável nenhuma, e nem mesmo, passado como argumento de nenhuma função,
	// portanto, o valor está sendo lido mas não está sendo usado;
	<-ch

	// nesta instrução, está sendo feito o envio do valor literal inteiro '2', para o
	// canal 'ch';
	ch <- 2

	// desta vez, o valor que é lido do canal 'ch' está sendo enviado, diretamente, como
	// argumento da função 'Println', para que seja impresso no console;
	fmt.Println(<-ch)
}
