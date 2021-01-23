package main

import (
	"github.com/yuchiki/YuchikiStackMachine/pkg/instruction"
	"github.com/yuchiki/YuchikiStackMachine/pkg/runner"
)

func main() {
	sampleCode := []uint64{
		uint64(instruction.OpPushI)<<32 + 5,  // pushi 5
		uint64(instruction.OpPushI)<<32 + 7,  // pushi 7
		uint64(instruction.OpAdd) << 32,      // add
		uint64(instruction.OpPrintInt) << 32, // printint
		uint64(instruction.OpPushI)<<32 + 0,  // pushi 0
		uint64(instruction.OpRet) << 32,      // ret

	}

	r := runner.NewRunner(sampleCode)

	r.Run()
}
