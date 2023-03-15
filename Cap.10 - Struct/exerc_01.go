// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	saboresSorvetes []string	
}
func main(){
	pessoa1 := pessoa{"Josias", "Barreto", []string{"Flocos", "Maracujá","Chocolate"}}
	pessoa2 := pessoa{
		nome: "Majuh",
		sobrenome: "Barreto",
		saboresSorvetes: []string{"Aptamil rss", "Materno", "NAN"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for indice, valor := range pessoa1.saboresSorvetes{
		fmt.Println(indice, valor)
	}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for indice, valor := range pessoa2.saboresSorvetes{
		fmt.Println(indice, valor)
	}
	

}