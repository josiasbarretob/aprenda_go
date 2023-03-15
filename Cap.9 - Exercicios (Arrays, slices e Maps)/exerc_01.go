/*
- Usando uma literal composta:
  - Crie um array que suporte 5 valores to tipo int
  - Atribua valores aos seus índices

- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
*/
package main

import "fmt"

func main(){
	numeros := [5]int{}
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

for indice, valor := range(numeros){
	fmt.Println(indice, valor)
}
fmt.Printf("Tipo: %T", numeros)
}