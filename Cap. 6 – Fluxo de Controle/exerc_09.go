package main

import "fmt"

var esporte string
func main(){
	fmt.Println("Favor insira o seu esporte favorito")
	fmt.Scan(&esporte)
	switch esporte{
	case "Futebol":
		fmt.Println("Ótima escolha! O futebol é a alma do Brasil")
	case "Voleibol":
		fmt.Println(esporte, "É bom tanto na praia quanto em quadra")
	case "Natação":
		fmt.Println(esporte, "É um esporte que precisa de fôlego e resistência!")
	}
}