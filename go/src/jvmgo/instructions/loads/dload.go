package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _dload(frame *rtda.Frame, index uint) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(index))
}

type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}