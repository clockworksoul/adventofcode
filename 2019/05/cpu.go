package main

import "fmt"

type ParameterMode int

const ioBufferSize = 1024

const (
	ModePosition  ParameterMode = 0
	ModeImmediate ParameterMode = 1
)

type CPU struct {
	intcode       []int
	ip            int
	halted        bool
	parameterMode ParameterMode
	istream       chan int
	ostream       chan int
	halt          chan bool
}

func NewCPU(intcode []int) *CPU {
	return &CPU{
		intcode: intcode,
		istream: make(chan int, ioBufferSize),
		ostream: make(chan int),
		halt:    make(chan bool)}
}

func (cpu *CPU) Peek() (OpCode, []ParameterMode, []int, error) {
	if cpu.ip >= len(cpu.intcode) {
		return CodeNull, []ParameterMode{}, []int{}, fmt.Errorf("stack overrun")
	}

	// Get intcode and parse off the opcode
	intcode := cpu.intcode[cpu.ip]
	opcode := OpCode(intcode % 100)
	inst := Instructions[opcode]

	// Parse off the parameter modes list
	modes := make([]ParameterMode, inst.plength)
	intcode /= 100
	for i := 0; i < inst.plength; i++ {
		modes[i] = ParameterMode(intcode % 10)
		intcode /= 10
	}

	params := cpu.intcode[cpu.ip+1 : cpu.ip+inst.plength+1]

	if len(params) != len(modes) || len(params) != inst.plength {
		return CodeNull,
			[]ParameterMode{},
			[]int{},
			fmt.Errorf("parameter length mismatch")
	}

	return OpCode(opcode), modes, params, nil
}

func (cpu *CPU) Execute() error {
	opcode, modes, params, err := cpu.Peek()
	if err != nil {
		return fmt.Errorf("execution error: %w", err)
	}

	fmt.Println("Before:", cpu, opcode, modes, params)

	inst := Instructions[opcode]
	execute := inst.function
	execute(cpu, params, modes)

	if inst.autoIncrement {
		cpu.ip += (1 + inst.plength)
	}

	fmt.Println("After: ", cpu)

	return nil
}

func (cpu *CPU) ExecuteAll() error {
	var err error

	for err == nil && !cpu.halted {
		err = cpu.Execute()
	}

	return err
}

func (cpu *CPU) read(mode ParameterMode, value int) int {
	if mode == ModePosition {
		return cpu.intcode[value]
	}

	return value
}

func (cpu *CPU) output(out int) {
	// fmt.Println(out)
	cpu.ostream <- out
}
