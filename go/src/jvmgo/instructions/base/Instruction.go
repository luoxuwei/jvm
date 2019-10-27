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

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader)  {
    self.Index = uint(reader.ReadInt8())
}
