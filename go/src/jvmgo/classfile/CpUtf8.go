package classfile

import (
	"jvmgo/vmutils"
)
/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	val string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	n := reader.readUint16()
	bytes := reader.readBytes(uint32(n))
	self.val = vmutils.DecodeMUTF8(bytes)
}