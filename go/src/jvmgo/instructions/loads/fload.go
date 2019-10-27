package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}