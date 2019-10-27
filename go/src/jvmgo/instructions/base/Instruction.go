package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}