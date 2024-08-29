package dsstack

type arrayStack struct {
	// 用数组来模拟栈
	data []int
}

func NewArrayStack() *arrayStack {

	return &arrayStack{
		data: make([]int, 0, 16),
	}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return s.size() == 0
}

func (s *arrayStack) push(v int) {
	// 不需要判断是否满了,go切片会自动扩容
	s.data = append(s.data, v)
}

func (s *arrayStack) pop() any {
	val := s.peek()
	s.data = s.data[:s.size()-1]
	return val
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	val := s.data[s.size()-1]
	return val
}

func (s *arrayStack) toSlice() []int {
	return s.data
}
