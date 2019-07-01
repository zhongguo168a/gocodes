package arrayutil

type ByteArray []byte

func (arr *ByteArray) RemoveAt(i int) {
	s := *arr
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0 // GC
	s2 := s[:len(s)-1]
	*arr = s2
}

func (arr *ByteArray) IndexOf(item byte) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if s[i] == item {
			return i
		}
	}
	return -1
}

func (arr *ByteArray) IndexOfFunc(predicate func(item byte) bool) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if predicate(s[i]) {
			return i
		}
	}
	return -1
}

func (arr *ByteArray) Push(item byte) {
	s := *arr
	s = append(s, item)
	*arr = s
}

func (arr *ByteArray) Pop() interface{} {
	s := *arr
	last := s[len(s)-1]
	s[len(s)-1] = 0 // GC
	s2 := s[:len(s)-1]
	*arr = s2

	return last
}

func (arr *ByteArray) Unshift(item byte) {
	s := *arr
	l := len(s) + 1
	ss := make([]byte, l, l)
	ss[0] = item
	copy(ss[1:], s[:])
	*arr = ss
}

func (arr *ByteArray) Shift() interface{} {
	s := *arr
	top := s[0]
	s[0] = 0 // GC
	s2 := s[1:]
	*arr = s2

	return top
}

func (arr *ByteArray) Replace(i int, item byte) {
	s := *arr
	over := i - len(s)
	if over > -1 {
		ss := make([]byte, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = item
	*arr = s
}
