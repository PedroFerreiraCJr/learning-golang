package main

import (
	"encoding/json"
	"fmt"
)

// essa declaração de struct define, adicionalmente, como ela pode ser convertida para uma
// string no formato de texto JSON; neste caso, o texto delimitado por crase '“', é como uma
// instância desta struct deve ser convertido, ou seja, quais são os nomes dos atributos de
// destino do JSON; da mesma forma, eles informam como deve ser feita a leitura de determinado
// texto no formato JSON para uma struct deste tipo;
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// transforma uma struct do tipo 'produto' para um texto no formato json;
	p1 := produto{ID: 1, Nome: "Notebook", Preco: 1899.90, Tags: []string{"Promoção", "Eletônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(p1Json)

	// transforma uma string, em texto no formato JSON, para uma struct do tipo 'produto';
	var p2 produto // o valor-zero de um tipo agregado é uma instância dele com os atributos com valores-zero;
	jsonString := `{"id":2, "nome":"Caneta", "preco":8.9, "tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
