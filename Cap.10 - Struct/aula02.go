package main

import "fmt"

type pessoa struct{
	nome string
	dataNascimento string
	sexo string
}
type profissional struct{
	pessoa
	cargo string
	salario float64
}
func main(){
	pessoa1 := pessoa{
		nome: "Jacinto Dores",
		dataNascimento: "23/06/2000",
		sexo: "M",
	}
	pessoa2 := profissional{
		pessoa : pessoa{
			nome: "Maricota",
			dataNascimento: "23/07/1990",
			sexo: "P",
		},
		cargo: "Professora",
		salario: 2876.89,
	}
	fmt.Println("Nome da 1ª pessoa:", pessoa1.nome)
	fmt.Println("Sexo da 1ª pessoa:", pessoa1.sexo)
	fmt.Println("Todas as informações da 1ª pessoa:", pessoa1)
	fmt.Println("Nome da 2ª pessoa:", pessoa2.pessoa.nome)
	fmt.Println("Todas as informações da 2ª pessoa:", pessoa2)
}