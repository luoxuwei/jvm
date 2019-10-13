package classfile
/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}

type LineNumberTableAttribute struct {
	LineNumberTable    []*LineNumberTableEntry
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
    length := reader.readUint16()
    self.LineNumberTable = make([]*LineNumberTableEntry, length)
    for i := range self.LineNumberTable {
    	self.LineNumberTable[i] = &LineNumberTableEntry{
			StartPc:    reader.readUint16(),
			LineNumber: reader.readUint16(),
		}
	}
}