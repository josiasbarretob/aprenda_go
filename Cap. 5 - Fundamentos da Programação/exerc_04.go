package main

import(
	"fmt"
)

func main(){
	numero := 200
	fmt.Printf("%d \n%#x \n%b", numero, numero, numero)

	x := numero << 1
	fmt.Printf("\n%d \n%#x \n%b", x, x, x)
}