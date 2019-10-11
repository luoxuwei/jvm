package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
    n := int(reader.readUint16())
    constantInfos := make([]ConstantInfo, n)
    for i := 1; i < n; i++ {
    	constantInfos[i] = readConstantInfo(reader, constantInfos)
		switch constantInfos[i].(type) {
		case *ConstantDoubleInfo, *ConstantLongInfo:
            i++ //long 和 double 类型占两个位置
		}
	}
	return constantInfos
}

func (self *ConstantPool) getConstantInfo(i uint16) ConstantInfo {
	if constantInfo := (*self)[i]; constantInfo != nil {
		return constantInfo
	}
	panic("Invalid constant pool index!")
}

func (self *ConstantPool) getUtf8(i uint16) string {
    utf8Info := (*self)[i].(*ConstantUtf8Info)
    return utf8Info.val
}