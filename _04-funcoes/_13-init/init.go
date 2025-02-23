package main

import "fmt"

// essa é uma função especial, porque ela é lida e executada antes de
// começar a executar o código contido na função 'main';
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Função main...")
}
