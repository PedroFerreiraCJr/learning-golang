package main

import (
	"fmt"
	"time"
)

// essa função estabelece em sua assinatura, que ela espera receber
// um tipo 'interface';
// então, de acordo com o tipo da informação passada, ela retorna uma
// 'string' que informa qual é o tipo passado como argumento para a função;
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func(): // tipo função;
		return "função"
	default:
		return "Desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
