package attend

import (
	"bufio"
	"net"
	"os"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

func Kernel(kernel net.Conn) {
	for {
		switch connection.ReciveCode(kernel) {
		case connection.NEW_PROCESS:
			newProcess(kernel)
		case connection.REMOVE_PROCESS:
			pid := connection.ReciveInt(kernel)
			processList.Remove((find(pid)))
		default:
			logger.Warning("Unexpected code from Kernel")
		}
	}
}

func newProcess(kernel net.Conn) {
	pid := connection.ReciveInt(kernel)
	path := connection.ReciveString(kernel)
	println(path)
	file, err := os.Open("test.txt")
	if err != nil {
		logger.Warning(err.Error())
		connection.SendCode(kernel, connection.FAILIURE)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	processList.PushBack(&memoryProcess{pid, lines})
	connection.SendCode(kernel, connection.SUCCES)
}
