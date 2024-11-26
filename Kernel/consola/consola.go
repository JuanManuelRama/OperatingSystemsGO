package consola

import (
	"bufio"
	"strings"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
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

func ProcessInput(input []string) {
	if fn, exists := commands[input[0]]; exists {
		fn(input[1:])
	} else {
		logger.Warning("Unknown command")
	}
}
