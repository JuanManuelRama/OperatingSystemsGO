package connection

import (
	"fmt"
	"log"
	"net"
)

func AcceptConection(who string, port string) net.Conn {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is waiting for ...", who)

	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
	}
	buf := make([]byte, 1024)
	conn.Read(buf)
	log.Println(string(buf))
	return conn
}

func Connect(who string, port string) net.Conn {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	message := "Hello from " + who
	conn.Write([]byte(message))
	return conn
}
