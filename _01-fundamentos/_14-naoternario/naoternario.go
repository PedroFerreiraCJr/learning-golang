package main

import "fmt"

// a linguagem Go não tem operador ternário!
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
