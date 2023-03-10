package main

import (
	"fmt"
)

func main(){
	for horas := 0 ; horas < 12; horas ++ {
		// fmt.Println(horas)
		for minutos := 0 ; minutos < 60 ; minutos ++ {
			// fmt.Println(minutos)
			for segundos := 0 ; segundos < 60 ; segundos ++{
				// fmt.Println(segundos)
				fmt.Println(horas,":", minutos,":", segundos)
			}
		}
	
	}
}