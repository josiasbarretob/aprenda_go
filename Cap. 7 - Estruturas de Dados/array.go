package main

import "fmt"

var x [10] int
var num int

func main(){
	x[0] = 5
	x[1] = 3
	x[2] = 2
	x[3] = 6
	x[4] = 8
	x[5] = 1
	x[6] = 0
	x[7] = 7
	x[8] = 8
	x[9] = 9
	tamanho := len(x)
	fmt.Println(x, "\n",tamanho)

	vetor := [10] int{}
	for i := 0 ; i < len(vetor) ; i++{
		fmt.Println("Insira um valor para a posição ", i ,"do vetor:")
		fmt.Scan(&num)
		vetor[i] = num
	}
	
	fmt.Println(vetor)

}