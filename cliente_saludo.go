package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Cliente computadora B")
	//Pedir nombre del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese su nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	//Conectarse al servidor
	fmt.Println("Conectándose con el servidor...")
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error no se pudo conectar con el servidor:", err)
		return
	}
	//Enviar el nombre al servidor
	fmt.Println("Se está enviando el nombre al servidor...")
	conn.Write([]byte(nombre))
	//Recibir el saludo del servidor
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	respuesta := string(buffer[:n])
	fmt.Println("El servidor respondió con éxito:", respuesta)
	conn.Close()
}
