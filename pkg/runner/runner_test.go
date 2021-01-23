package runner

import (
	"testing"

	"github.com/yuchiki/YuchikiStackMachine/pkg/instruction"
)

func TestRunnerRun(t *testing.T) {
	sampleCode := []uint64{
		uint64(instruction.OpPushI)<<32 + 5, // pushi 5
		uint64(instruction.OpRet) << 32,     // ret

	}

	runner := Runner{
		code: sampleCode,
	}

	returnValue := runner.Run()
	t.Log(returnValue)
	if returnValue != 5 {
		t.Errorf("expected 5, but actual return value is %v", returnValue)
	}
}
