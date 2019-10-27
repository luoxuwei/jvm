package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _lload(frame *rtda.Frame, index uint) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(index))
}

type LLOAD struct {
	base.Index8Instruction
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, self.Index)
}
