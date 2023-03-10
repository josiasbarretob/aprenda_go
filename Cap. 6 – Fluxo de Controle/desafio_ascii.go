package main

import (
	"fmt"
)

func main(){
	for i := 32 ; i <= 122 ; i++{
		fmt.Printf("Decimal: %d | Hexadecimal: %#x | Unicode: %#U \n", i, i, i)
	}
}