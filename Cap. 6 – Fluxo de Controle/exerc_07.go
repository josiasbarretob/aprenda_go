package main

import "fmt"

func main(){
	ativo := false
	if ativo{
		fmt.Println("Sistema ativo! Manutenção adiada.")
	}else if ativo == false{
		fmt.Println("Sistema inativo! Ordem de Manutenção iniciada!")
	}else{
		fmt.Println("Status inválido! Error")
	}
}