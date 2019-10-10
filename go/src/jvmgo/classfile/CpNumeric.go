package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	self.val = int32(reader.readUint32())
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	self.val = math.Float32frombits(reader.readUint32())
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	self.val = int64(reader.readUint64())
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	self.val = math.Float64frombits(reader.readUint64())
}


