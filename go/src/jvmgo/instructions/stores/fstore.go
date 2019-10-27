package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _fstore(frame *rtda.Frame, index uint) {
	frame.LocalVars().SetFloat(index, frame.OperandStack().PopFloat())
}

type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, self.Index)
}
