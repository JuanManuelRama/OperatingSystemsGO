package attend

import (
	"net"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

func Kernel(kernel net.Conn) {
	for {
		switch connection.ReciveCode(kernel) {
		case connection.NEW_PROCESS:
			logger.Info(connection.ReciveString(kernel))
			connection.SendCode(kernel, connection.SUCCES)
			connection.SendString(kernel, "Instructions")
		default:
			logger.Warning("Unexpected code from Kernel")
		}
	}
}
