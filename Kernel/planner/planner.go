package planner

import (
	"net"
	"strconv"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
)

var NewChan = make(chan Process)
var ReadyChan = make(chan Process)
var KillChan = make(chan Process)

func LongTermPlanner(memory net.Conn) {
	for {
		process := <-NewChan

		if sendNewProcess(process, memory) == connection.FAILIURE {
			logger.Warning("Memory couldn't open file")
			continue
		}
		process.PCB.State = "READY"
		ReadyChan <- process
	}

}

func sendNewProcess(process Process, memory net.Conn) uint8 {
	connection.SendCode(memory, connection.NEW_PROCESS)
	connection.SendInt(memory, process.PCB.PID)
	connection.SendString(memory, process.Joker)
	return connection.ReciveCode((memory))
}

func ShortTermPlanner(cpu net.Conn) {
	for {
		process := <-ReadyChan
		process.PCB.State = "RUNNING"
		connection.SendPCB(cpu, process.PCB)
		logger.Info("Running process: " + process.Joker)

		process.PCB = connection.RecivePCB(cpu)
		switch connection.ReciveCode(cpu) {
		case connection.EXIT:
			deathRow(process, "Succes")
		default:
			logger.Warning("Unexpected exit code")
		}
	}
}

func deathRow(process Process, motive string) {
	process.PCB.State = "EXIT"
	process.Joker = motive
	KillChan <- process
}

func Killer(memory net.Conn) {
	for {
		process := <-KillChan
		connection.SendCode(memory, connection.REMOVE_PROCESS)
		connection.SendInt(memory, process.PCB.PID)
		log_finish(process.PCB.PID, process.Joker)
	}
}

func log_finish(pid int32, motive string) {
	logger.Info("PID: " + strconv.Itoa(int(pid)) + " Finishes - Motive: " + motive)
}
