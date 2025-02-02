package main

import (
	"fmt"
	"math"
)

func main() {
	// as duas variáveis são do tipo 'int';
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E lógico bit a bit =>", a&b)
	fmt.Println("Ou lógico bit a bit =>", a|b)
	fmt.Println("Xor lógico bit a bit =>", a^b)

	c := 3.0
	d := 2.0

	// outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
