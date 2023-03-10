package main

import ("fmt")

type newType int
var x newType

func main(){
	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Printf("%v \n", x)
}