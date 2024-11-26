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
	return conn
}
