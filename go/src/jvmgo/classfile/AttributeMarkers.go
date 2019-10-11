package classfile

type MarkerAttribute struct {

}

type DeprecatedAttribut struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}
