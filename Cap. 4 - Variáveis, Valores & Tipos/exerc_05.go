package main

import ("fmt")

type newType int
var x newType
var y int
func main(){
	fmt.Printf("%v \n", x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Printf("%v \n", x)
	y = int(x)
	fmt.Printf("%T \n", y)
}