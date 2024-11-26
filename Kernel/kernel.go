package main

import "github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"

func main() {
	cpu := connection.Connect("kernel", "08080")
	connection.SendInt(cpu, 10)
	defer cpu.Close()
	for {
	}
}
