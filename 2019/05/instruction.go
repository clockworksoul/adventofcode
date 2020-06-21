package main

type OpCode byte

const (
	CodeNull        OpCode = 0
	CodeAdd         OpCode = 1
	CodeMul         OpCode = 2
	CodeIn          OpCode = 3
	CodeOut         OpCode = 4
	CodeJumpIfTrue  OpCode = 5
	CodeJumpIfFalse OpCode = 6
	CodeLessThan    OpCode = 7
	CodeEquals      OpCode = 8
	CodeHalt        OpCode = 99
)

type Instruction struct {
	plength       int
	autoIncrement bool
	function      InstructionFunc
}

var Instructions = map[OpCode]Instruction{
	CodeAdd:         Instruction{3, true, executeAdd},
	CodeMul:         Instruction{3, true, executeMul},
	CodeIn:          Instruction{1, true, executeIn},
	CodeOut:         Instruction{1, true, executeOut},
	CodeJumpIfTrue:  Instruction{2, false, executeJumpIfTrue},
	CodeJumpIfFalse: Instruction{2, false, executeJumpIfFalse},
	CodeLessThan:    Instruction{3, true, executeLessThan},
	CodeEquals:      Instruction{3, true, executeEquals},
	CodeHalt:        Instruction{0, true, executeHalt},
}

type InstructionFunc func(cpu *CPU, params []int, modes []ParameterMode)

func executeAdd(cpu *CPU, params []int, modes []ParameterMode) {
	pnoun, pverb, paddress := params[0], params[1], params[2]
	mnoun, mverb := modes[0], modes[1]
	noun, verb := cpu.read(mnoun, pnoun), cpu.read(mverb, pverb)
	cpu.intcode[paddress] = noun + verb
}

func executeMul(cpu *CPU, params []int, modes []ParameterMode) {
	pnoun, pverb, paddress := params[0], params[1], params[2]
	mnoun, mverb := modes[0], modes[1]
	noun, verb := cpu.read(mnoun, pnoun), cpu.read(mverb, pverb)
	cpu.intcode[paddress] = noun * verb
}

func executeIn(cpu *CPU, params []int, modes []ParameterMode) {
	paddress := params[0]
	input := <-cpu.istream
	cpu.intcode[paddress] = input
}

func executeOut(cpu *CPU, params []int, modes []ParameterMode) {
	// cpu.ostream <- cpu.read(modes[0], params[0])
	cpu.output(cpu.read(modes[0], params[0]))
}

// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the
// instruction pointer to the value from the second parameter. Otherwise, it does nothing.
func executeJumpIfTrue(cpu *CPU, params []int, modes []ParameterMode) {
	if cpu.read(modes[0], params[0]) != 0 {
		pos := cpu.read(modes[1], params[1])
		cpu.ip = pos
	}
}

// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction
// pointer to the value from the second parameter. Otherwise, it does nothing.
func executeJumpIfFalse(cpu *CPU, params []int, modes []ParameterMode) {
	if cpu.read(modes[0], params[0]) == 0 {
		pos := cpu.read(modes[1], params[1])
		cpu.ip = pos
	}
}

// Opcode 7 is less than: if the first parameter is less than the second parameter,
// it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
func executeLessThan(cpu *CPU, params []int, modes []ParameterMode) {
	ma, mb, mpos := modes[0], modes[1], modes[2]
	pa, pb, ppos := params[0], params[1], params[2]
	a, b, pos := cpu.read(ma, pa), cpu.read(mb, pb), cpu.read(mpos, ppos)

	value := 0
	if a < b {
		value = 1
	}

	cpu.intcode[pos] = value
}

// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores
// 1 in the position given by the third parameter. Otherwise, it stores 0.
func executeEquals(cpu *CPU, params []int, modes []ParameterMode) {
	ma, mb, mpos := modes[0], modes[1], modes[2]
	pa, pb, ppos := params[0], params[1], params[2]
	a, b, pos := cpu.read(ma, pa), cpu.read(mb, pb), cpu.read(mpos, ppos)

	value := 0
	if a == b {
		value = 1
	}

	cpu.intcode[pos] = value
}

func executeHalt(cpu *CPU, params []int, modes []ParameterMode) {
	cpu.halted = true
	cpu.halt <- true
}
