package main

import "fmt"

func main() {
	// mapas devem ser inicializados, caso contrário, ao tentar utilizar essa
	// estrutura, será gerado um erro em tempo de execução;
	//var aprovados map[int]string

	// para instânciar um map, é possível utilizar a função 'make';
	// o argumento para ela deve representar o tipo do map que queremos instanciar;
	// neste caso, estamos informando que o map tem chaves do tipo 'int', e
	// os valores relacionados a uma chave qualquer, tem o tipo 'string';
	aprovados := make(map[int]string)

	// as instruções abaixo estão adicionando valores ao map 'aprovados';
	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"

	// essa é uma forma de percorrer os pares, chave e valor, de dentro de uma
	// estrutura do tipo 'map';
	// a variável 'cpf', recebe o valor da chave, que neste caso é um 'int';
	// a variável 'nome', recebe o valor relacionado a determinada chave, que
	// neste caso é uma 'string';
	for cpf, nome := range aprovados {
		// imprime no console, através da saída formatada;
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// obtem determinado valor da estrutura map, através da chave;
	fmt.Println(aprovados[95135745682])

	// utiliza a função 'delete', para remover da estrutura map um determinado
	// valor, que tem a chave igual ao segundo argumento;
	delete(aprovados, 95135745682)

	// obtem um valor na estrutura map, a partir da chave;
	fmt.Println(aprovados[95135745682])
}
