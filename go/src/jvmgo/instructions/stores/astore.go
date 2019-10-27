package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _astore(frame *rtda.Frame, index uint) {
	frame.LocalVars().SetRef(index, frame.OperandStack().PopRef())
}

type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, self.Index)
}
