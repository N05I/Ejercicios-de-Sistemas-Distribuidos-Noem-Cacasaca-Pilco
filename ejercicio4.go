// BUCLE EN GO
// EL CLASICO for
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("NUMERO:", i)
	}

	NUMERO := 2
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", NUMERO, i, NUMERO*i)
	}
}
