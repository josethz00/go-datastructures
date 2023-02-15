package t_stack

import (
	"errors"
	"fmt"
)

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

type StackStr struct {
	max   int
	top   int
	value []string
}

func NewStackStr(max int) *StackStr {
	return &StackStr{
		max:   max,
		top:   -1,
		value: make([]string, max),
	}
}

func (s *StackStr) isEmpty() bool {
	return s.top == -1
}

func (s *StackStr) isFull() bool {
	return s.top == s.max-1
}

func (s *StackStr) Push(newValue string) {
	if s.isFull() {
		return
	}

	s.top++
	s.value[s.top] = newValue
}

func (s *StackStr) Pop() string {
	if s.isEmpty() {
		return "-1"
	}

	value := s.value[s.top]
	s.top--

	return value
}

func (s *StackStr) Top() string {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return "-1"
	}

	return s.value[s.top]
}

func (s *StackStr) Print() {
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

type StackRune struct {
	max   int
	top   int
	value []rune
}

func NewStackRune(max int) *StackRune {
	return &StackRune{
		max:   max,
		top:   -1,
		value: make([]rune, max),
	}
}

func (s *StackRune) IsEmpty() bool {
	return s.top == -1
}

func (s *StackRune) IsFull() bool {
	return s.top == s.max-1
}

func (s *StackRune) Push(newValue rune) {
	if s.IsFull() {
		return
	}

	s.top++
	s.value[s.top] = newValue
}

func (s *StackRune) Pop() (rune, error) {
	if s.IsEmpty() {
		return ' ', errors.New("Stack is empty")
	}

	value := s.value[s.top]
	s.top--

	return value, nil
}

func (s *StackRune) Top() (rune, error) {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return ' ', errors.New("Stack is empty")
	}

	return s.value[s.top], nil
}

func (s *StackRune) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	for i := s.top; i >= 0; i-- {
		fmt.Println("-----")
		fmt.Println("|", s.value[i], "|")
		fmt.Println("-----")
	}
}
