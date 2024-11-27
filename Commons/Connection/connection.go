package connection

import (
	"encoding/binary"
	"net"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

func StartServer(port string) net.Listener {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		logger.Error(err.Error() + "while starting server")
	}
	logger.Info("Server on port " + port + " started succesfully")
	return listener
}

func AcceptConection(who string, listener net.Listener) net.Conn {
	logger.Info("Server is waiting for " + who)
	conn, err := listener.Accept()
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

func SendCode(conn net.Conn, code uint8) {
	err := binary.Write(conn, binary.BigEndian, code)
	if err != nil {
		logger.Error(err.Error() + "while sending code")
	}
}

func ReciveCode(conn net.Conn) uint8 {
	var code uint8
	err := binary.Read(conn, binary.BigEndian, &code)
	if err != nil {
		logger.Error(err.Error() + "while reading code")
	}
	return code
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
