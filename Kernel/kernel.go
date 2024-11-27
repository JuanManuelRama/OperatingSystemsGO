package main

import (
	"bufio"
	"os"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/consola"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/planner"
)

func main() {
	memory := connection.Connect("kernel", "8080")
	//cpu := connection.Connect("kernel", "08081")

	//connection.SendInt(cpu, 1)
	go planner.LongTermPlanner(memory)
	go planner.ShortTermPlanner()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		consola.ProcessInput(consola.ReadInput(scanner))
	}
}
