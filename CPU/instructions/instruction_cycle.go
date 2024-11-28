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
	"SET":  SET,
	"SUM":  SUM,
	"SUB":  SUB,
	"JNZ":  JNZ,
	"EXIT": EXIT,
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
	log_execute(process.PID, instruction, args[1:])
	instructions[instruction](process, args[1:])
}

func log_fetch(pid int32, pc uint32) {
	logger.Info("PID: " + fmt.Sprint(pid) + " - FETCH - Program Counter: " + fmt.Sprint(pc))
}

func log_execute(pid int32, instruction string, args []string) {
	logger.Info("PID: " + fmt.Sprint(pid) + " - Executing: " + instruction + " - " + fmt.Sprint(args))
}
