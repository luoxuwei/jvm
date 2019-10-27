package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _istore(frame *rtda.Frame, index uint) {
	frame.LocalVars().SetInt(index, frame.OperandStack().PopInt())
}
type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, self.Index)
}