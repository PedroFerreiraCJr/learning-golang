package main

/**
para compilar e executar esse arquivo, execute o seguinte comando
em qualquer terminal, dentro desta pasta:
	go run *.go
*/
// a função main é obrigatório porque estamos no package main
func main() {
	// essa linha declara e já inicializa a variável 'resultado' com o valor
	// retornado pela função 'somar'; a função 'somar' declara dois parâmetros;
	// neste caso, dois parâmetros do tipo 'int'; portanto, é preciso informar
	// dois valores do tipo 'int' no momento da invocação da função 'somar';
	resultado := somar(1, 2)

	// imprime no console o resultado da operação realizada pela função 'somar';
	imprimir(resultado)
}
