package main

import "fmt"

// essa é uma implementação do calculo de fatorial de um número inteiro;
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	// invocação de sucesso da função fatorial;
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	// invocação com um argumento incorreto, o que gera um erro;
	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
