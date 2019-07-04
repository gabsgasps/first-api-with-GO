package main

import "fmt"

// Teste - contain name and sobrenome
type Teste struct {
	name      string
	sobrenome string
}

// com * reserva na memoria "uma unidade"
func changeName(n *Teste) {
	n.name = "gabriel"
}

func main() {
	// com o sinal	"&" acessamos o mesmo endereco na memoria em que o "*" foi colocado
	person := Teste{name: "joao", sobrenome: "gasparino"}
	fmt.Println(person)

	changeName(&person)
	// fmt.Println(person)
	// fmt.Println(&person)

	// sem referenciar o pointer
	// tt := person

	// referenciando o pointer
	tt := &person
	fmt.Println("tt", tt)
	tt.name = "gadhjsdjhjh"
	fmt.Println("person", person)
	fmt.Println("tt", tt)

}
