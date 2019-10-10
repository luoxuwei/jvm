package classfile

type ClassFile struct {
	magic          uint32
	MinorVersion   uint32
	MajorVersion   uint32
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
    self.magic = reader.readUint32()
    if self.magic != 0xCAFEBABE {
    	panic("java.lang.ClassFormatError: magic!")
	}
}