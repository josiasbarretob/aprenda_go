package main

import (
	"fmt"
)

func main(){
	var numero int 
	x := 7
	a := (10 == 10)
	b := (10 > 10)
	c := (10 < 10)
	d := (10 >= 10)
	e := (10 <= 10)
	f := (10 != 10)
	fmt.Printf("%v, %v, %v, %v, %v, %v \n", a, b, c, d, e, f)
	println()
	fmt.Println("************** Descubra o valor de X: **************")
	fmt.Println("Favor digite um número inteiro de 1 a 10:")
	fmt.Scan(&numero)
	
	if numero > x {
		fmt.Println("O número digitado é maior que X!")
	} else if numero == x {
		fmt.Println("Parabéns! O número digitado é igual a X!")
	}else{
		fmt.Println("O número digitado é menor que X!")
	}

}