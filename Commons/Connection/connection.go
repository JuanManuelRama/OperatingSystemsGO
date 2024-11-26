package connection

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func AcceptConection(who string, port string) net.Conn {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Println("Server is waiting for ...", who)
	conn, err := listener.Accept()
	listener.Close()
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

func SendCode(conn net.Conn, code byte) {
	_, err := conn.Write([]byte{code})
	if err != nil {
		log.Fatal("Error sending code:", err)
	}
}

func ReciveCode(conn net.Conn) byte {
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading code:", err)
	}
	return buf[0]
}

func SendInt(conn net.Conn, number int32) {
	err := binary.Write(conn, binary.BigEndian, number)
	if err != nil {
		log.Fatal("Error sending int:", err)
	}
}

func ReciveInt(conn net.Conn) int32 {
	var number int32
	err := binary.Read(conn, binary.BigEndian, &number)
	if err != nil {
		log.Fatal("Error reading int:", err)
	}
	return number
}

func SendString(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
}

func ReciveString(conn net.Conn) string {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading message:", err)
	}
	return string(buf)
}
