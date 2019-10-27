package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _aload(frame *rtda.Frame, index uint) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(index))
}

type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}
