package main

import (
	"bufio"
	"encoding/gob"
	"os"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/process"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/consola"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/planner"
)

func main() {
	gob.Register(process.PCB{})
	memory := connection.Connect("kernel", "8080")
	cpu := connection.Connect("kernel", "08081")

	go planner.LongTermPlanner(memory)
	go planner.ShortTermPlanner(cpu)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		consola.ProcessInput(consola.ReadInput(scanner))
	}
}
