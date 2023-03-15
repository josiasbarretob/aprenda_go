// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	saboresSorvetes []string	
}
func main(){
	mapa1 := map[string]pessoa{}
	pessoa1 := pessoa{"Josias", "Barreto", []string{"Flocos", "Maracujá","Chocolate"}}
	// pessoa2 := pessoa{
	// 	nome: "Majuh",
	// 	sobrenome: "Barreto",
	// 	saboresSorvetes: []string{"Aptamil rss", "Materno", "NAN"}}
	
	mapa1[pessoa1.sobrenome] = pessoa1
	fmt.Println(mapa1)
	}