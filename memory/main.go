package main

import (
	"sync"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/connection"
	"github.com/JuanManuelRama/OperatingSyStemsGO/memory/attend"
)

func main() {
	var wait sync.WaitGroup
	go attend.Kernel(connection.AcceptConection("kernel", "8081"))
	wait.Add(1)
	wait.Wait()
}
