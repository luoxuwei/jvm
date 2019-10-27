package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _lstore(frame *rtda.Frame, index uint) {
	frame.LocalVars().SetLong(index, frame.OperandStack().PopLong())
}

type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, self.Index)
}
