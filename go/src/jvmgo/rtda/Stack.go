package rtda

type Stack struct {
	maxSize uint
	size    uint
   _top    *Frame
}

func newStack(size uint) *Stack {
	return &Stack{maxSize:size}
}

func (self *Stack) push(frame *Frame) {
	if self.size > self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	ret := self._top
	self._top = ret.lower
	ret.lower = nil
	self.size--
	return ret
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
