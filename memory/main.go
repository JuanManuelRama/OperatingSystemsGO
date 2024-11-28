package main

import (
	"sync"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/memory/attend"
)

func main() {
	var wait sync.WaitGroup
	server := connection.StartServer("8080")
	go attend.CPU(connection.AcceptConection("cpu", server))
	go attend.Kernel(connection.AcceptConection("kernel", server))
	wait.Add(1)
	wait.Wait()
}
