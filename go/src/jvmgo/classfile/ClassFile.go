package classfile
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magic          uint32
	MinorVersion   uint16
	MajorVersion   uint16
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
    self.magic = reader.readUint32()
    if self.magic != 0xCAFEBABE {
    	panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
    self.MinorVersion = reader.readUint16()
    self.MajorVersion = reader.readUint16()
	switch self.MajorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.MinorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}