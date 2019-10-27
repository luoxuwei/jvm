package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _dstore(frame *rtda.Frame, index uint) {
	frame.LocalVars().SetDouble(index, frame.OperandStack().PopDouble())
}

type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, self.Index)
}
