/*
- Usando uma literal composta:
  - Crie uma slice de tipo int
  - Atribua 10 valores a ela

- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
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
	for indice, valor := range(numerosSlice){
		fmt.Printf("Indice = [%d] - Valor = %d\n", indice, valor)
	}
	fmt.Printf("%T",numerosSlice)
}