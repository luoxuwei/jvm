package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _iload(frame *rtda.Frame, index uint)  {
	frame.OperandStack().PushInt(frame.LocalVars().GetInt(index))
}

type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
    _iload(frame, self.Index)
}
 