package main

import (
	//"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"bufio"
	"os"

	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/consola"
)

func main() {
	/*cpu := connection.Connect("kernel", "08080")
	connection.SendInt(cpu, 10)
	defer cpu.Close()*/

	scanner := bufio.NewScanner(os.Stdin)
	for {
		consola.ProcessInput(consola.ReadInput(scanner))
	}
}
