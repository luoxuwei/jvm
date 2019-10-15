package rtda

import "math"

type OperandStack struct {
	size uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
    if maxStack>0 {
    	return &OperandStack{slots:make([]Slot, maxStack)}
	}

	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].Val = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].Val
}

func (self *OperandStack) PushFloat(val float32) {
	self.slots[self.size].Val = int32(math.Float32bits(val))
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	return math.Float32frombits(uint32(self.slots[self.size].Val))
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].Val = int32(val)
	self.size++
	self.slots[self.size].Val = int32(val>>32)
	self.size++
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].Val)
	high := uint32(self.slots[self.size+1].Val)
	return int64(low) | int64(high<<32)
}

func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].Ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].Ref
	self.slots[self.size].Ref = nil
	return ref
}
