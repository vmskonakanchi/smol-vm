package main

type StackItem struct {
	name  string
	value interface{}
}

type Stack[T comparable] struct {
	items        []T
	maxSize      int
	currentIndex int
}

func CreateNewStack[T comparable](size int) *Stack[T] {
	return &Stack[T]{
		items:        make([]T, size),
		maxSize:      size,
		currentIndex: 0,
	}
}

func (s *Stack[T]) Push(item T) {
	if s.currentIndex >= s.maxSize {
		log.Fatal("Stack size exceeded")
		return
	}
	s.items[s.currentIndex] = item
	s.currentIndex++
}

func (s *Stack[T]) Pop() T {
	if s.currentIndex == 0 {
		log.Fatal("Cannot pop from empty stack")
	}
	s.currentIndex--
	item := s.items[s.currentIndex]

	// clearing the slot to help GC
	var zero T
	s.items[s.currentIndex] = zero

	return item
}

func (s *Stack[T]) Peek() T {
	if s.currentIndex == 0 {
		log.Fatal("Cannot peek empty stack")
	}
	return s.items[s.currentIndex-1]
}

func (s *Stack[T]) Empty() bool {
	return s.currentIndex == 0
}

func (s *Stack[T]) Size() int {
	return s.currentIndex
}
