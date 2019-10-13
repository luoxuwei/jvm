package classfile
/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
    AccessFlags    uint16
    NameIndex      uint16
    DescriptorIndex uint16
    AttributeInfos  []AttributeInfo
}

func readMembers(reader *ClassReader, cp *ConstantPool) []MemberInfo {
	n := reader.readUint16()
	infos := make([]MemberInfo, n)
	for i := range infos {
		infos[i] = MemberInfo{
			AccessFlags:     reader.readUint16(),
			NameIndex:       reader.readUint16(),
			DescriptorIndex: reader.readUint16(),
			AttributeInfos:  readAttributes(reader, cp),
		}
	}
	return infos
}

func (self *MemberInfo) Name(cp *ConstantPool) string {
	return cp.getUtf8(self.NameIndex)
}

