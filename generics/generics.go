package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Printf("a: %d, b: %d\n", a, b)
	x := "John"
	y := "Joe"
	x, y = swap(x, y)
	fmt.Printf("x: %s, y: %s\n", x, y)

	t := Stack[int]{elements: []int{1}}
	t.push(2)
	fmt.Println(t)
	fmt.Println(t.pop())
	fmt.Println(t)
	fmt.Println(t.pop())
	fmt.Println(t.pop())
	fmt.Println(t.isEmpty())
}
