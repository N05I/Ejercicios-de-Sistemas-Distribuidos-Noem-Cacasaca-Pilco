package main

import "fmt"

func main() {
	contador := 0
	for contador < 10 {
		fmt.Println("COntador: ", contador)
		contador++
		//contador +=1
		//contador = contador+1
	}
	//FOR infinito con BREAK
	contador1 := 0
	for {
		if contador1 >= 10 {
			break
		}
		fmt.Println("COntador1: ", contador1)
		contador1++
	}
}
