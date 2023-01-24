package t_stack

import "fmt"

type Stack struct {
	max   int
	top   int
	value []int
}

func NewStack(max int) *Stack {
	return &Stack{
		max:   max,
		top:   -1,
		value: make([]int, max),
	}
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top == s.max-1
}

func (s *Stack) Push(newValue int) {
	if s.isFull() {
		return
	}

	s.top++
	s.value[s.top] = newValue
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return -1
	}

	value := s.value[s.top]
	s.top--

	return value
}

func (s *Stack) Top() int {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}

	return s.value[s.top]
}

func (s *Stack) Print() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	for i := s.top; i >= 0; i-- {
		fmt.Println("-----")
		fmt.Println("|", s.value[i], "|")
		fmt.Println("-----")
	}
}
