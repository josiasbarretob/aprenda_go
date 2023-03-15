/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
  - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
  - Do quinto ao último item do slice (incluindo o último item!)
  - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
  - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
  - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/
package main

import "fmt"

var x int

func main(){
	numerosSlice := []int{}
	for i := 0 ; i < 10 ; i++{
		fmt.Scan(&x)
		numerosSlice = append(numerosSlice, x)
	} 
	fmt.Println(numerosSlice[:3])
	fmt.Println(numerosSlice[4:])
	fmt.Println(numerosSlice[1:7])
	fmt.Println(numerosSlice[2:9])
	fmt.Println(numerosSlice[2:len(numerosSlice)-1])
}