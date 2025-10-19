package main

import "fmt"

func main() {
	//definicion de variable corta
	nombre := "Noemi"

	//definicion de variable con var

	var apellidos string = "Cacasaca Pilco"

	//definicion de variable multiple
	var (
		peso   = 52
		talla  = 1.64
		estado = true
	)
	//tipos de variables
	var edad int = 42
	var temperatura float64 = 20.5
	var aprobado bool = true

	fmt.Println(nombre, apellidos, peso, talla, estado)
	fmt.Println(edad, temperatura, aprobado)
}
