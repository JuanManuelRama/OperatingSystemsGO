package consola

import (
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/logger"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/utils"
	"github.com/JuanManuelRama/OperatingSyStemsGO/kernel/planner"
)

var pid = 0

func ExecuteScript(args []string) {
	if len(args) == 0 {
		logger.Warning("No script provided")
		return
	}
	logger.Info("Executing script:" + args[0])
}

func StartProcess(args []string) {
	if len(args) == 0 {
		logger.Warning("No process provided")
		return
	}
	logger.Info("Starting process: " + args[0])
	go func() {
		planner.NewChan <- planner.Process{planner.PCB{pid, 0, planner.Registers{PC: 0}, "NEW", 0}, args[0]}
		pid++
	}()
}

func KillProcess(args []string) {
	if len(args) == 0 {
		logger.Warning("No process provided")
		return
	}
	n := utils.Atoi(args[0])
	if n < 0 {
		logger.Warning("Invalid process number")
		return
	}
	logger.Info("Killing process: " + args[0])
}

func StopPlanification(_ []string) {
	logger.Info("Stopping planification")
}

func StartPlanification(_ []string) {
	logger.Info("Starting planification")
}

func ProcessList(_ []string) {
	logger.Info("Process list")
}
