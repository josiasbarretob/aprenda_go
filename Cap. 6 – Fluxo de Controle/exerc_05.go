package main

import "fmt"

func main(){
	for i := 10; i <= 100 ; i++{
		x := i % 4
		fmt.Println("Número:", i,"| Divisão por 4:", x)
	}
}