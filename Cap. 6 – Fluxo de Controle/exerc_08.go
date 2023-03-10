package main

import "fmt"

func main(){
	cardapio := 1

	switch{
	case cardapio == 0:
		fmt.Println("Strogonof + Arroz")
	case cardapio == 1:
		fmt.Println("Feijoada")
	case cardapio == 2:
		fmt.Println("Carne assada com Purê de Batata")
	case cardapio == 3:
		fmt.Println("Mocotó")
	case cardapio == 4:
		fmt.Println("Caldo Verde")
	case cardapio == 5:
		fmt.Println("Caldo de Pinto")
	case cardapio == 6:
		fmt.Println("Dobradinha")
	case cardapio == 7:
		fmt.Println("Filé de Tilápia + Fritas")
	default:
		fmt.Println("Opção inválida!")

	}
}