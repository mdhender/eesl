// eesl - embeddable, extendable scripting language
// Copyright (c) 2023 Michael D Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package stack

// Stack implements a generic stack structure.
type Stack[T any] struct {
	data []T
}

// New returns a new stack for type T.
func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Peek returns the top element on the stack.
// If the stack is empty, returns a zero-value T and false.
func (s *Stack[T]) Peek() (tos T, ok bool) {
	if len(s.data) == 0 {
		return tos, false
	}
	return s.data[len(s.data)-1], true
}

// Pop returns the top element on the stack.
// If the stack is empty, returns a zero-value T and false.
func (s *Stack[T]) Pop() (tos T, ok bool) {
	if tos, ok = s.Peek(); ok {
		tos = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
	}
	return tos, ok
}

// Push adds a new element to the stack.
func (s *Stack[T]) Push(n T) {
	s.data = append(s.data, n)
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}
