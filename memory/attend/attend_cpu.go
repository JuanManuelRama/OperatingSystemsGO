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
		case connection.RUN:
			currentProcess = findPID(connection.ReciveInt(cpu))
		}
	}
}
