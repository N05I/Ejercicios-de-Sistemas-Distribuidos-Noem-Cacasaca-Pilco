//package main

//import "fmt"

//func main() {
//nombre := "Alex"
//saludo := saludo(nombre)
//fmt.Println(saludo)
//}
//func saludo(nombre string) string {
//return "BIENVENIDOS A SISTEMAS DISTRIBUIDOS " + nombre
//}

package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	saludo := saludar(nombre)
	fmt.Println(saludo)
}
func saludar(nombre string) string {
	return "BIENVENIDOS A SISTEMAS DISTRIBUIDOS: " + nombre
}
