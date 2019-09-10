package arrayutil

import (
	"sort"
)

type ArrayListInt []int

func (arr *ArrayListInt) At(index int) int {
	return (*arr)[index]
}

func (arr *ArrayListInt) Clear() {
	*arr = (*arr)[:0]
}

func (arr *ArrayListInt) Clone() ArrayListInt {
	s := *arr
	ss := make([]int, len(s), cap(s))
	copy(ss, s)

	return ss
}

func (arr *ArrayListInt) Contains(item int) bool {
	return arr.IndexOf(item) != -1
}

func (arr *ArrayListInt) ContainsCond(cond func(item int) bool) bool {
	return arr.IndexOfConf(cond) != -1
}

// 每一个项都符合条件就返回true
func (arr *ArrayListInt) Every(cond func(item int) bool) bool {
	s := arr.Clone()
	for i := 0; i < len(s); i++ {
		val := s[i]
		if cond(val) == false {
			return false
		}
	}
	return true
}

func (arr *ArrayListInt) First(cond func(item int) bool) (int, bool) {
	s := arr.Clone()
	for i := 0; i < len(s); i++ {
		val := s[i]
		if cond(val) {
			return val, true
		}
	}
	return 0, false
}

// 过滤数组, 返回符合条件[conf]的新数组
func (arr *ArrayListInt) Filter(cond func(index int, elem int) bool) (r ArrayListInt) {
	for i, x := range arr.Clone() {
		if cond(i, x) {
			r = append(r, x)
		}
	}
	return r
}

func (arr *ArrayListInt) ForRange(handler func(item int)) {
	s := arr.Clone()
	for i := 0; i < len(s); i++ {
		handler(s[i])
	}
}

func (arr *ArrayListInt) IndexOfConf(cond func(item int) bool) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if cond(s[i]) {
			return i
		}
	}
	return -1
}

func (arr *ArrayListInt) IndexOf(item int) int {
	s := *arr
	for i := 0; i < len(s); i++ {
		if s[i] == item {
			return i
		}
	}
	return -1
}

func (arr *ArrayListInt) Last(cond func(item int) bool) int {
	s := arr.Clone()
	for i := len(s) - 1; i >= 0; i-- {
		val := s[i]
		if cond(val) {
			return val
		}
	}
	return 0
}

func (arr *ArrayListInt) Slice() []int {
	return []int(*arr)
}

func (arr *ArrayListInt) Length() int {
	return len(*arr)
}

func (arr *ArrayListInt) Pop() int {
	s := *arr
	last := s[len(s)-1]
	s[len(s)-1] = 0 // GC
	s2 := s[:len(s)-1]
	*arr = s2

	return last
}

func (arr *ArrayListInt) Push(item int) {
	s := *arr
	s = append(s, item)
	*arr = s
}

func (arr *ArrayListInt) PushList(list ArrayListInt) {
	s := *arr
	s = append(s, list...)
	*arr = s
}

func (arr *ArrayListInt) Remove(item int) {
	i := arr.IndexOf(item)
	if i != -1 {
		arr.RemoveAt(i)
	}
}

func (arr *ArrayListInt) RemoveAt(i int) {
	s := *arr
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0 // GC
	s2 := s[:len(s)-1]
	*arr = s2
}

func (arr *ArrayListInt) Replace(i int, item int) {
	s := *arr
	over := i - len(s)
	if over > -1 {
		ss := make([]int, i+1)
		copy(ss[0:], s[:])
		s = ss
	}
	s[i] = item
	*arr = s
}

func (arr *ArrayListInt) Reverse() {
	for i := len(*arr)/2 - 1; i >= 0; i-- {
		opp := len(*arr) - 1 - i
		(*arr)[i], (*arr)[opp] = (*arr)[opp], (*arr)[i]
	}
}

func (arr *ArrayListInt) Shift() int {
	s := *arr
	top := s[0]
	s[0] = 0 // GC
	s2 := s[1:]
	*arr = s2

	return top
}

func (arr *ArrayListInt) Sort(compare func(a, b int) int) {
	l := *arr
	sort.Slice(l, func(i, j int) bool {
		return compare(l[i], l[j]) >= 0
	})
}

func (arr *ArrayListInt) Unshift(item int) {
	s := *arr
	l := len(s) + 1
	ss := make([]int, l, l)
	ss[0] = item
	copy(ss[1:], s[:])
	*arr = ss
}

// 去重操作, 返回去重后的数组
func (arr *ArrayListInt) Unique(getKey func(a int) string) (r ArrayListInt) {
	l := *arr
	m := map[string]int{} // 存放不重复主键
	for _, e := range l {
		length := len(m)
		m[getKey(e)] = 0
		if len(m) != length { // 加入map后，map长度变化，则元素不重复
			r = append(r, e)
		}
	}
	return
}

// 并集
func (arr *ArrayListInt) Union(a ArrayListInt, getKey func(a int) string) ArrayListInt {
	arr.PushList(a)
	return arr.Unique(getKey)
}
