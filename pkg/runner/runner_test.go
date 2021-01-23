package runner

import (
	"testing"

	"github.com/yuchiki/YuchikiStackMachine/pkg/instruction"
	"github.com/yuchiki/YuchikiStackMachine/pkg/stack"
)

func TestRunnerRun(t *testing.T) {
	sampleCode := []uint64{
		uint64(instruction.OpPushI)<<32 + 5, // pushi 5
		uint64(instruction.OpPushI)<<32 + 7, // pushi 7
		uint64(instruction.OpAdd) << 32,     // add
		uint64(instruction.OpRet) << 32,     // ret

	}

	runner := Runner{
		code:         sampleCode,
		workingStack: stack.NewUint64Stack(),
	}

	returnValue := runner.Run()
	t.Log(returnValue)
	if returnValue != 12 {
		t.Errorf("expected 5, but actual return value is %v", returnValue)
	}
}
