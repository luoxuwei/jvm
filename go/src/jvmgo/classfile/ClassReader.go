package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	ret := self.data[0]
	self.data = self.data[1:]
    return ret
}

func (self *ClassReader) readUint16() uint16 {
	ret := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return ret
}

func (self *ClassReader) readUint32() uint32 {
	ret := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return ret
}

func (self *ClassReader) readUint64() uint64 {
	ret := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return ret
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	ret := self.data[:n]
	self.data = self.data[n:]
	return ret
}



