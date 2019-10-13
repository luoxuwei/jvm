package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/

type ExceptionAttribute struct {
	ExceptionIndexTable []uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	n := reader.readUint16()
	self.ExceptionIndexTable = make([]uint16, n)
	for i := range self.ExceptionIndexTable {
		self.ExceptionIndexTable[i] = reader.readUint16()
	}
}