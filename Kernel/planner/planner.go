package planner

import (
	"net"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

var NewChan = make(chan Process)
var ReadyChan = make(chan Process)

func LongTermPlanner(memory net.Conn) {
	for {
		process := <-NewChan
		process.PCB.Instructions = sendNewProcess(process, memory)
		if process.PCB.Instructions == nil {
			logger.Warning("Memory couldn't open file")
			continue
		}
		process.PCB.State = "READY"
		ReadyChan <- process
	}

}

func sendNewProcess(process Process, memory net.Conn) interface{} {
	connection.SendCode(memory, connection.NEW_PROCESS)
	connection.SendString(memory, process.Joker)
	if connection.ReciveCode((memory)) == connection.SUCCES {
		return connection.ReciveString(memory)
	}
	return nil
}

func ShortTermPlanner() {
	for {
		process := <-ReadyChan
		process.PCB.State = "RUNNING"
		logger.Info("Running process: " + process.Joker)
	}
}
