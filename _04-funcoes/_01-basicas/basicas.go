package main

import "fmt"

// essa função não declara nenhum parâmetro, nem um tipo de retorno;
// essa função é conhecida como uma função 'void', em linguagens como, Java
// C, e C++; pois ela não tem nenhum valor de retorno;
func f1() {
	fmt.Println("f1")
}

// essa função declara 2 parâmetros, cada um deles, do tipo 'string';
// essa função é uma função 'void', porque ela não declara que retorna um
// valor qualquer;
func f2(param1 string, param2 string) {
	fmt.Printf("F2: %s %s\n", param1, param2)
}

// essa função "promete" retornar um valor do tipo 'string';
// neste caso, quando essa função for executada ela terá como resultado da
// sua execução, o valor textual: F3;
func f3() string {
	return "F3"
}

// essa função declara 2 parâmetros do tipo 'string';
// essa é outra forma de fazer a mesma declaração da função 'f2', que declara
// dois parâmetros; mas neste caso, como os parâmetros são do mesmo tipo, podemos
// fazer a declaração deles de uma forma idiomática da linguagem Golang;
// essa função retorna uma 'string';
func f4(param1, param2 string) string {
	return fmt.Sprintf("F4: %s %s", param1, param2)
}

// essa função declara mais de um valor de retorno;
// essa é uma característica específica da linguagem Golang;
// uma função em Golang, pode declarar que retorna mais de um valor após a
// execução dela;
// neste exemplo de função, ela declara que retorna 2 valores do tipo 'string';
// para declarar que uma função em Golang retorna mais de um valor quando
// executada, basta definir quais são os tipos dos valores que serão retornados
// pela função, e eles devem ser declarados dentro de parênteses, conforme o
// seguinte exemplo;
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param 1", "Param 2")

	r3, r4 := f3(), f4("Param 1", "Param 2")
	fmt.Printf("%s %s", r3, r4)

	// neste caso, a função 'f5' retorna 2 valores, portanto, está sendo feita
	// a atribuição do valor retornado para as variáveis locais a função 'main',
	// e que foram chamadas de 'r51', e 'r52';
	// é importante notar que, também poderiamos ter utilizado o operador '_'
	// (underline), para informar ao compilador que não estamos interessados
	// no valor retornado pela função;
	// também é importante notar que, caso a função retorne valores e não seja
	// de interesse do chamador o valor retornado, então, não seria necessário
	// atribuir os valores retornados a variáveis;
	r51, r52 := f5()
	fmt.Println("F5: ", r51, r52)
}
