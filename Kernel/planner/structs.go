package planner

type Process struct {
	PCB   PCB
	Joker string
}

type PCB struct {
	PID          int
	Quantum      int
	Registers    Registers
	State        string
	Instructions interface{}
}

type Registers struct {
	PC  uint32
	AX  uint8
	BX  uint8
	CX  uint8
	DX  uint8
	EAX uint32
	EBX uint32
	ECX uint32
	EDX uint32
	SI  uint32
	DI  uint32
}
