package connection

import (
	"encoding/binary"
	"net"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

func AcceptConection(who string, port string) net.Conn {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error(err.Error() + "while starting server")
	}
	logger.Info("Server is waiting for " + who)
	conn, err := listener.Accept()
	listener.Close()
	if err != nil {
		logger.Error(err.Error() + "while accepting connection")
	}
	buf := make([]byte, 1024)
	conn.Read(buf)
	logger.Info(string(buf))
	return conn
}

func Connect(who string, port string) net.Conn {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		logger.Error(err.Error() + "while connecting to server")
	}
	message := "Hello from " + who
	conn.Write([]byte(message))
	return conn
}

func SendCode(conn net.Conn, code byte) {
	_, err := conn.Write([]byte{code})
	if err != nil {
		logger.Error(err.Error() + "while sending code")
	}
}

func ReciveCode(conn net.Conn) byte {
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err != nil {
		logger.Error(err.Error() + "while reading code")
	}
	return buf[0]
}

func SendInt(conn net.Conn, number int32) {
	err := binary.Write(conn, binary.BigEndian, number)
	if err != nil {
		logger.Error(err.Error() + "while sending int")
	}
}

func ReciveInt(conn net.Conn) int32 {
	var number int32
	err := binary.Read(conn, binary.BigEndian, &number)
	if err != nil {
		logger.Error(err.Error() + "while reading int")
	}
	return number
}

func SendString(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		logger.Error(err.Error() + "while sending message")
	}
}

func ReciveString(conn net.Conn) string {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		logger.Error(err.Error() + "while reading message")
	}
	return string(buf)
}
