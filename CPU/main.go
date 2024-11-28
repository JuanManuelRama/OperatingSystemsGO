package main

import (
	"encoding/gob"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/process"
	"github.com/JuanManuelRama/OperatingSyStemsGO/cpu/instructions"
)

func main() {
	gob.Register(process.PCB{})
	memory := connection.Connect("cpu", "8080")
	server := connection.StartServer("8081")
	kernel := connection.AcceptConection("kernel", server)
	for {
		instructions.Exit = false
		process := connection.RecivePCB(kernel)
		connection.SendCode(memory, connection.RUN)
		connection.SendInt(memory, process.PID)
		for !instructions.Exit {
			instruction := instructions.Fetch(process, memory)
			args := instructions.Decode(instruction)
			instructions.Execute(&process, args)
			process.Registers.PC++
		}
		connection.SendPCB(kernel, process)
		connection.SendCode(kernel, connection.EXIT)
		println(process.Registers.AX)
	}
}
