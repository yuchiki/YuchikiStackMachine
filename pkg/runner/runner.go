package runner

import "github.com/yuchiki/YuchikiStackMachine/pkg/instruction"

// Runner とは、仮想機械本体である。
type Runner struct {
	code           []uint64
	workingStack   uint64
	programCounter uint64
	isFinished     bool
}

func (r *Runner) Run() uint64 {
	for !r.isFinished {
		r.Step()
	}
	return r.workingStack
}

func (r *Runner) Step() {
	opCode, immediate := instruction.Instruction(r.code[r.programCounter]).Decomposite()

	switch opCode {
	case instruction.OpNop:
		r.Next()
	case instruction.OpAdd:
		panic("unimplemented!")
	case instruction.OpPushI:
		r.workingStack = uint64(immediate)
		r.Next()
	case instruction.OpRet:
		r.isFinished = true
	default:
		panic("unknown op")
	}
}

func (r *Runner) Next() {
	r.programCounter = r.programCounter + 1
}
