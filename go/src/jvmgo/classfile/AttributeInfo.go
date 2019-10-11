package classfile
/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	count := reader.readUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	name := cp.getUtf8(nameIndex)
	l := reader.readUint32()
	info := newAttributeInfo(name, l, cp)
	info.readInfo(reader)
	return info
}

func newAttributeInfo(name string, l uint32, cp ConstantPool) AttributeInfo {
	switch name {
	case "Code":
		return &CodeAttribute{cp:cp}
	default:
		return &UnparseAttribute{
			name:name,
			length:l,
			info:nil,
		}
	}
}
