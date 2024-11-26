package consola

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadInput(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Fields(scanner.Text())
}

var commands = map[string]func([]string){
	"EXECUTE_SCRIPT":      ExecuteScript,
	"START_PROCESS":       StartProcess,
	"KILL_PROCESS":        KillProcess,
	"STOP_PLANIFICATION":  StopPlanification,
	"START_PLANIFICATION": StartPlanification,
	"PROCESS_LIST":        ProcessList,
}

func ProcesInput(input []string) {
	if fn, exists := commands[input[0]]; exists {
		fn(input[1:])
	} else {
		fmt.Println("Unknown command")
	}
}
