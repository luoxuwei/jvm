package classfile

type ConstantValueAttribute struct {
	ConstantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.ConstantValueIndex = reader.readUint16()
}