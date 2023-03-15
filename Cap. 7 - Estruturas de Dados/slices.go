package main

import "fmt"

var quantidade int
var nota float64

func main(){
	carro := [] string{"Camaro", "Golf", "Hilux", "Corsa"}
	fmt.Println(carro) // Imprime o Slice inteiro

	for i := 0 ; i < len(carro) ; i++{
		fmt.Println(i, carro[i])
	}

	fmt.Println("Calculadora de Média:")
	fmt.Println("Informe a quantidade de notas que deseja calcular a Média:")
	fmt.Scan(&quantidade)
	media := [] float64{}
	for i := 1 ; i <= quantidade ; i++{
		fmt.Println("Favor digite a ", i ,"nota:")
		fmt.Scan(&nota)
		media = append(media, nota)
	}
	for _, valor := range(media){
		fmt.Println(valor)
	}
	fatia := media[:]
	fmt.Println(fatia)
}