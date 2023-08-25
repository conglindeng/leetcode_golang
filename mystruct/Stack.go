package mystruct

type Stack struct {
	size int
	nums []interface{}
}

func Stack_Init() *Stack {
	return &Stack{size: 0, nums: make([]interface{}, 0)}
}
func (s *Stack) Push(val interface{}) {
	s.nums = append(s.nums, val)
	s.size++
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.nums[s.size-1], true
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.size == 0 {
		return nil, false
	}
	res := s.nums[s.size-1]
	s.nums = s.nums[0 : s.size-1]
	s.size--
	return res, true
}

func (s *Stack) Size() int {
	return s.size
}