package main

import "fmt"

type alunos struct{
	nome string
	matricula int
	serie string
	tipoSanguineo string
	ativo bool
}

func main(){
	alunos1 := alunos{
		nome: "Josias Barreto",
		matricula: 683583,
		serie: "7º série",
		tipoSanguineo: "O+",
		ativo: false,
	}
	alunos2 := alunos{"Majuh Peixoto", 897653, "2º ano", "AB+", true}
	fmt.Println(alunos1, alunos2)

}