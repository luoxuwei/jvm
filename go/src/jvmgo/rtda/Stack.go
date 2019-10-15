package rtda

type Stack struct {
   _top    *Frame
}

func newStack(size int) *Stack {
	return &Stack{}
}

func (self *Stack) push(frame *Frame) {

}

func (self *Stack) pop() *Frame {
	return self._top
}

func (self *Stack) top() *Frame {
	return self._top
}
