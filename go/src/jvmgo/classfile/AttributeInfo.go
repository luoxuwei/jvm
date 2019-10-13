package classfile
const (
	ConstantValue                       = "ConstantValue"                        //	1.0.2
	Code                                = "Code"                                 //	1.0.2
	Exceptions                          = "Exceptions"                           //	1.0.2
	SourceFile                          = "SourceFile"                           //	1.0.2
	LineNumberTable                     = "LineNumberTable"                      //	1.0.2
	LocalVariableTable                  = "LocalVariableTable"                   //	1.0.2
	InnerClasses                        = "InnerClasses"                         //	1.1
	Synthetic                           = "Synthetic"                            //	1.1
	Deprecated                          = "Deprecated"                           //	1.1
	EnclosingMethod                     = "EnclosingMethod"                      //	5.0
	Signature                           = "Signature"                            //	5.0
	SourceDebugExtension                = "SourceDebugExtension"                 //	5.0
	LocalVariableTypeTable              = "LocalVariableTypeTable"               //	5.0
	RuntimeVisibleAnnotations           = "RuntimeVisibleAnnotations"            //	5.0
	RuntimeInvisibleAnnotations         = "RuntimeInvisibleAnnotations"          //	5.0
	RuntimeVisibleParameterAnnotations  = "RuntimeVisibleParameterAnnotations"   //	5.0
	RuntimeInvisibleParameterAnnotation = "RuntimeInvisibleParameterAnnotations" //	5.0
	AnnotationDefault                   = "AnnotationDefault"                    //	5.0
	StackMapTable                       = "StackMapTable"                        //	6
	BootstrapMethods                    = "BootstrapMethods"                     //	7
	RuntimeVisibleTypeAnnotations       = "RuntimeVisibleTypeAnnotations"        //	8
	RuntimeInvisibleTypeAnnotations     = "RuntimeInvisibleTypeAnnotations"      //	8
	MethodParameters                    = "MethodParameters"                     //	8
	Module                              = "Module"                               // 9
	ModulePackages                      = "ModulePackages"                       // 9
	ModuleMainClass                     = "ModuleMainClass"                      // 9
	NestHost                            = "NestHost"                             // 11
	NestMembers                         = "NestMembers"                          // 11
)
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
	case Code:
		return &CodeAttribute{cp:cp}
	default:
		return &UnparseAttribute{
			name:name,
			length:l,
			info:nil,
		}
	}
}
