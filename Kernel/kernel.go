package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	message := "Hola"
	conn.Write([]byte(message))
}
