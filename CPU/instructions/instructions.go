package instructions

import (
	"reflect"

	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/process"
	"github.com/JuanManuelRama/OperatingSyStemsGO/commons/utils"
)

var Exit bool

func SET(process *process.PCB, args []string) {
	register := args[0]
	value := uint(utils.Atoi(args[1]))
	setRegister(register, &process.Registers, value)
}

func SUM(process *process.PCB, args []string) {
	destiny := args[0]
	source := args[1]
	destinyValue := getRegister(destiny, process.Registers)
	sourceValue := getRegister(source, process.Registers)
	setRegister(destiny, &process.Registers, destinyValue+sourceValue)
}

func SUB(process *process.PCB, args []string) {
	destiny := args[0]
	source := args[1]
	destinyValue := getRegister(destiny, process.Registers)
	sourceValue := getRegister(source, process.Registers)
	setRegister(destiny, &process.Registers, destinyValue-sourceValue)
}

func JNZ(process *process.PCB, args []string) {
	register := args[0]
	value := uint(utils.Atoi(args[1]))
	if getRegister(register, process.Registers) != 0 {
		setRegister("PC", &process.Registers, value)
	}
}

func EXIT(process *process.PCB, args []string) {
	Exit = true
}

func getRegister(register string, registers process.Registers) uint {
	println(register)
	v := reflect.ValueOf(registers)
	field := v.FieldByName(register)
	return uint(field.Uint())
}

func setRegister(register string, registers *process.Registers, value uint) {
	v := reflect.ValueOf(registers).Elem()
	field := v.FieldByName(register)
	switch field.Kind() {
	case reflect.Uint8:
		field.SetUint(uint64(uint8(value)))
	case reflect.Uint32:
		field.SetUint(uint64(uint32(value)))
	}
}
