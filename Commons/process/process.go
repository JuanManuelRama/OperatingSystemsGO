package process

type PCB struct {
	PID       int32
	Quantum   int
	Registers Registers
	State     string
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
