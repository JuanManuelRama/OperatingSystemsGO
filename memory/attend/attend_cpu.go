package attend

import (
	"net"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
)

func CPU(cpu net.Conn) {
	for {
		switch connection.ReciveCode(cpu) {
		case connection.FETCH:
			connection.SendString(cpu, currentProcess.instructions[connection.ReciveInt(cpu)])
			println("FETCH")
		case connection.RUN:
			pid := connection.ReciveInt(cpu)
			currentProcess = findPID(pid)
		}
	}
}
