package classfile

type ConstantStringInfo struct {
	cp ConstantPool
	i  uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader)  {
	self.i = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.i)
}