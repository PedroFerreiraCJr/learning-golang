package main

import "fmt"

// essa função utiliza o conceito de retorno nomeado;
// que são variáveis declaradas na assinatura da função;
// neste caso, a função 'trocar' declara 2 parâmetros do tipo 'int';
// esta função também declara, na parte de retorno da função, 2 valores
// do tipo 'int'; isso é chamado de Retorno Nomeado;
// como os dois valores nomeados são do mesmo tipo, também é possível
// utilizar uma notação mais sucinta, omitindo o tipo do primeiro parâmetro,
// já que ambos são do mesmo tipo, que neste caso é 'int';
// os identificadores declarados na assinatura da função, através do conceito
// de retorno nomeado, tem escopo de bloco da função onde foram declarados;
// também é importante notar que, a ordem de declaração dos identificadores
// do retorno nomeado é importante, porque caso seja utilizado o retorno limpo
// o ordem destes identificadore será mantida no retorno da função;
func trocar(param1, param2 int) (segundo, primeiro int) {
	// os identificadores declarados, ou retorno nomeado, são variáveis normais;
	// portanto, é possível atribuir valores do mesmo tipo das variáveis;
	segundo = param2
	primeiro = param1

	// neste caso, está sendo utilizado o conceito de Retorno Limpo;
	// quando os valores do retorno nomeado já tiverem valores definidos, e a
	// função está pronta para retornar esses valores, é possível utilizar,
	// simplesmente, a palavra-chave 'return' para concluir a execução da
	// função, mas, retornando dela os valores informados nas variáveis definidas
	// no retorno nomeado;
	// mas, também poderia ser feito retorno explícito dos identificadores
	// declarados no retorno nomeado; também daria certo;
	return
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
