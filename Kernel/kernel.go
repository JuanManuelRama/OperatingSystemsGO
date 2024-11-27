package main

import (
	"bufio"
	"os"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/consola"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/planner"
)

func main() {
	memory := connection.Connect("kernel", "8081")

	/*cpu := connection.Connect("kernel", "08080")
	connection.SendInt(cpu, 10)
	defer cpu.Close()*/
	go planner.LongTermPlanner(memory)
	go planner.ShortTermPlanner()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		consola.ProcessInput(consola.ReadInput(scanner))
	}
}
