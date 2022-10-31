package main

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

type GStack[T any] []T

func (s *GStack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *GStack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *GStack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *GStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var empty T
		return empty, false
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]
		return elem, true
	}
}
