package classfile

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