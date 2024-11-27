package main

import (
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/process"
	"github.com/JuanManuelRama/OperatingSyStemsGO/cpu/instructions"
)

func main() {
	/*memory := connection.Connect("cpu", "8080")
	server := connection.StartServer("8081")
	kernel := connection.AcceptConection("kernel", server)
	x := connection.ReciveInt(kernel)
	connection.SendInt(memory, x)
	log.Printf("Recived %d\n", x)
	defer kernel.Close()*/
	process := process.PCB{PID: 1, Registers: process.Registers{PC: 0}}
	instructions.Execute(&process, instructions.Decode("SET AX 10"))
	instructions.Execute(&process, instructions.Decode("SET BX 5"))
	instructions.Execute(&process, instructions.Decode("ADD AX BX"))
	println(process.Registers.AX)
}
