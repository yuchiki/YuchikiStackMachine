package instruction

import "fmt"

// OpCode はInstructionの上半分で、命令種別を表す
type OpCode uint32

const (
	OpNop OpCode = iota
	OpAdd
	OpPushI
	OpRet
)

func (c OpCode) String() string {
	switch c {
	case OpNop:
		return "nop"
	case OpAdd:
		return "add"
	case OpPushI:
		return "pushI"
	case OpRet:
		return "ret"
	default:
		return "unknown"
	}
}

// Instruction は OpCodeと Immediateからなる
type Instruction uint64

func (i Instruction) OpCode() OpCode {
	return OpCode(uint64(i) >> 32)
}

func (i Instruction) Immediate() uint32 {
	return uint32(uint64(i) % (2 << 32))
}

func (i Instruction) Decomposite() (OpCode, uint32) {
	return i.OpCode(), i.Immediate()
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %d", i.OpCode().String(), i.Immediate())
}
