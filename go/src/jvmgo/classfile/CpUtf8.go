package classfile

import (
	"jvmgo/vmutils"
)

type ConstantUtf8Info struct {
	val string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	n := reader.readUint16()
	bytes := reader.readBytes(uint32(n))
	self.val = vmutils.DecodeMUTF8(bytes)
}