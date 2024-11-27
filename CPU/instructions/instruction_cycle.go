package instructions

import (
	"fmt"
	"net"
	"strings"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/process"
)

var instructions = map[string]func(*process.PCB, []string){
	"SET": SET,
	"ADD": ADD,
	"SUB": SUB,
	"JNZ": JNZ,
}

func Fetch(process process.PCB, memory net.Conn) string {
	log_fetch(process.PID, process.Registers.PC)
	connection.SendCode(memory, connection.FETCH)
	connection.SendInt(memory, int32(process.Registers.PC))
	return connection.ReciveString(memory)
}

func Decode(buff string) []string {
	return strings.Fields(buff)

}

func Execute(process *process.PCB, args []string) {
	instruction := args[0]
	instructions[instruction](process, args[1:])
}

func log_fetch(pid int, pc uint32) {
	logger.Info("PID: " + fmt.Sprint(pid) + " - FETCH - Program Counter: " + fmt.Sprint(pc))
}
