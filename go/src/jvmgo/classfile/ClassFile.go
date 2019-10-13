package classfile

import "fmt"

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
	ConstantPool   ConstantPool
	accessFlags    uint16
	thisClass      uint16
	superClass     uint16
	interfaces     []uint16
	FieldInfos      []MemberInfo
	MethodInfos     []MemberInfo
	AttributeInfos  []AttributeInfo
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

func Parse(bytes []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{data:bytes}
    cf = &ClassFile{}
    cf.read(cr)
    return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
    self.ConstantPool = readConstantPool(reader)
    self.accessFlags = reader.readUint16()
    self.thisClass = reader.readUint16()
    self.superClass = reader.readUint16()
    self.interfaces = reader.readUint16s()
    self.FieldInfos = readMembers(reader, &self.ConstantPool)
    self.MethodInfos = readMembers(reader, &self.ConstantPool)
    self.AttributeInfos = readAttributes(reader, &self.ConstantPool)
}

func (self *ClassFile) AccessFlags() uint16 {
    return self.accessFlags
}

func (self *ClassFile) ClassName() string {
	return self.ConstantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	return self.ConstantPool.getClassName(self.superClass)
}

func (self *ClassFile) InterfaceNames() interface{} {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.ConstantPool.getClassName(cpIndex)
	}
	return interfaceNames
}