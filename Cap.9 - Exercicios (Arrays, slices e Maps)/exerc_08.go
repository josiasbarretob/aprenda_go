package main

import
"fmt"

var nome string
var numero, opcao int

func main(){
	agenda := map[string]int {}
	fmt.Println("Agenda Telefônica!")
	status := true
	for status{
		fmt.Println("Digite o Nome e o Sobrenome do contato:")
		fmt.Scan(&nome)
		fmt.Println("Digite o Número do contato:")
		fmt.Scan(&numero)
		agenda[nome] = numero
		fmt.Println("Contato salvo! Digite [1] para cadastrar um novo contato e [2] para sair")
		fmt.Scan(&opcao)
		if opcao == 2{
			status = false
			break
		}
		}
		agenda["Josias"] = 12345678
		delete(agenda, "Josias")
		for key, value := range agenda{
			fmt.Println(key, " - ", value)
		}
	}
