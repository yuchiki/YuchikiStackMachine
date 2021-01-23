package runner

import (
	"fmt"

	"github.com/yuchiki/YuchikiStackMachine/pkg/instruction"
	"github.com/yuchiki/YuchikiStackMachine/pkg/stack"
)

// Runner とは、仮想機械本体である。
type Runner struct {
	code           []uint64
	workingStack   stack.Uint64Stack
	programCounter uint64
	isFinished     bool
}

func NewRunner(code []uint64) *Runner {
	return &Runner{
		code:         code,
		workingStack: stack.NewUint64Stack(),
	}
}

func (r *Runner) Run() uint64 {
	for !r.isFinished {
		r.Step()
	}
	n, _ := r.workingStack.Pop()
	return n
}

func (r *Runner) Step() error {
	opCode, immediate := instruction.Instruction(r.code[r.programCounter]).Decomposite()

	switch opCode {
	case instruction.OpNop:
		r.Next()
	case instruction.OpAdd:
		left, err := r.workingStack.Pop()
		if err != nil {
			return err
		}

		right, err := r.workingStack.Pop()
		if err != nil {
			return err
		}

		r.workingStack.Push(left + right)
		r.Next()
	case instruction.OpPushI:
		r.workingStack.Push(uint64(immediate))
		r.Next()
	case instruction.OpRet:
		r.isFinished = true
	case instruction.OpPrintInt:
		n, err := r.workingStack.Pop()
		if err != nil {
			return err
		}
		fmt.Println(n)
		r.Next()
	default:
		panic("unknown op")
	}

	return nil
}

func (r *Runner) Next() {
	r.programCounter = r.programCounter + 1
}
