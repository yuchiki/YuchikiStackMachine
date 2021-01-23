package instruction

import "testing"

func TestInstructionOpCode(t *testing.T) {
	i := Instruction(1<<32 + 3)
	if uint32(i.OpCode()) != uint32(OpAdd) {
		t.Errorf("expected opAdd, but actually %v", i.OpCode())
	}
}
