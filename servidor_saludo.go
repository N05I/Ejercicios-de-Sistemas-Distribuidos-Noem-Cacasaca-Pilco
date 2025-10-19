package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Este es el servidor saludos - Computadora A")
	fmt.Println("Esperando conexiones...")

	listener, _ := net.Listen("tcp", ":9000")
	for {
		conn, _ := listener.Accept()
		fmt.Println("Cliente conectado:")
		//Leer el nombre del cliente
		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		nombre := string(buffer[:n])
		fmt.Println("Recibí el nombre:", nombre)

		//Enviar saludo al cliente
		respuesta := "Hola " + nombre + " BIENVENIDOS A SISTEMAS DISTRIBUIDOS"
		conn.Write([]byte(respuesta))
		fmt.Println("Envie saludo")
		conn.Close()
	}
}
