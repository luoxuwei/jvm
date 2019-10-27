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

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}