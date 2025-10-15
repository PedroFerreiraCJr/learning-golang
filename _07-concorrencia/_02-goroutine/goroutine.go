package main

import (
	"fmt"
	"time"
)

// declara uma função que poderá ser executada através de uma goroutine;
// observe que ela não tem nada de especial, é uma declaração de função normal da
// linguagem Go;
func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// o seguinte trecho de código tem a invocação da função que foi declarada logo acima;
	// essa é uma invocação convencional, porque não tem nenhuma outra sintaxe que não tenha sido
	// apresentada até o momento;
	// a chamada de função abaixo, tem o comportamento esperado, onde a segunda chamada de função
	// somente será invocada após o término da primeira;
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// já o seguinte trecho de código, apresenta uma invocação que não tinha sido vista até o
	// momento, onde é feito o uso da palavra-chave 'go';
	// uma invocação feita desta forma, é como a linguagem Go estabelece para que uma função
	// seja invocada em sua própria linha de execução;
	// descomente o trecho de código abaixo e veja que o resultado da execuçaõ do seguinte código
	// não tem um resultado que talvez fosse o esperado;
	// mas por quê? porque a "thread main" que executa a função main não espera que o código das
	// linhas abaixo termine de executar;
	// fazendo uma comparação com o comportamento da linguagem Java, que usa o conceito de thread
	// é como se as threads iniciadas pelos códigos abaixo tivessem sido configuradas para que elas
	// fossem threads daemon, e threads daemon somente permanecem executando normalmente enquanto
	// a thread main estivem rodando;
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	// neste caso, adicionando a palavra-chave 'go' apenas na primeira chamada de função, faz com
	// que a primeira invocação seja executada em outra goroutine; já a segunda invocação faz com
	// que a função seja invocada na mesma "thread" que a função main;
	// desta forma, é possível notar que a enquanto a segunda função estiver sendo executada, a
	// primeira invocação de função também está sendo executada;
	// então, apenas enquanto a função main, que é executada na "thread" main estiver em execução
	// é que a função que estivem sendo invocada através da palavra-chave 'go', que por sua vez,
	// é executada em uma outra "thread", que na verdade na linguagem Go é conhecida por
	// Gorountine, vai estar em execução;
	// portanto, é preciso fazer com que a thread da função main não termine, caso contrário,
	// qualquer função que tenha sido disparada para ser executada por uma goroutine, vai ser
	// concluída abruptamente;
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
