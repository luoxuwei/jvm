package classfile

type ConstantPool []ConstantInfo
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
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

func (self *ConstantPool) getNameAndType(i uint16) (string, string)  {
    info := (*self)[i].(*ConstantNameAndTypeInfo)
    return self.getUtf8(info.nameIndex), self.getUtf8(info.descriptorIndex)
}

func (self *ConstantPool) getClassName(i uint16) string {
    info := (*self)[i].(*ConstantClassInfo)
    return self.getUtf8(info.nameIndex)
}