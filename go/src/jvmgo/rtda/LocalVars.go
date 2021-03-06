package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self *LocalVars) SetInt(index uint, val int32) {
	(*self)[index].Val = val
}

func (self *LocalVars) GetInt(index uint) int32 {
	return (*self)[index].Val
}

func (self *LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	(*self)[index].Val = int32(bits)
}

func (self *LocalVars) GetFloat(index uint) float32 {
	bits := uint32((*self)[index].Val)
	return math.Float32frombits(bits)
}

func (self *LocalVars) SetLong(index uint, val int64) {
	(*self)[index].Val = int32(val)
	(*self)[index+1].Val = int32(val>>32)
}

func (self *LocalVars) GetLong(index uint) int64 {
	low := uint32((*self)[index].Val)
	high := uint32((*self)[index+1].Val)
    return int64(high)<<32 | int64(low)
}

func (self *LocalVars) SetDouble(index uint, val float64) {
    bits := math.Float64bits(val)
    self.SetLong(index, int64(bits))
}

func (self *LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self *LocalVars) SetRef(index uint, ref *Object) {
	(*self)[index].Ref = ref
}

func (self *LocalVars) GetRef(index uint) *Object {
	return (*self)[index].Ref
}