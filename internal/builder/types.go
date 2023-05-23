// eesl - embeddable, extendable scripting language
// Copyright (c) 2018-2023 Michael D Henderson
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

package builder

import "github.com/mdhender/eesl/internal/tree"

// Tree is a parse tree node. Symbol can either be a terminal (Lexeme) or a non-terminal
// (see Builder's Enter method). Lexemes matched using Builder's Match method or added
// using Builder's Add method, can be retrieved by type asserting Symbol.
// Subtrees are child nodes of the current node.
type Tree struct {
	Symbol   interface{}
	Subtrees []*Tree
}

func NewTree(symbol interface{}, subtrees ...*Tree) *Tree {
	t := Tree{Symbol: symbol}
	for _, subtree := range subtrees {
		if subtree != nil {
			t.Subtrees = append(t.Subtrees, subtree)
		}
	}
	return &t
}

func (t *Tree) Data() interface{} {
	if t == nil {
		return ""
	}
	return t.Symbol
}

func (t *Tree) Children() (c []tree.Node) {
	for _, subtree := range t.Subtrees {
		c = append(c, subtree)
	}
	return
}

// Add adds a subtree as a child to t.
func (t *Tree) Add(subtree *Tree) {
	t.Subtrees = append(t.Subtrees, subtree)
}

// Detach removes a subtree as a child of t.
func (t *Tree) Detach(subtree *Tree) {
	for i, st := range t.Subtrees {
		if st == subtree {
			t.Subtrees = append(t.Subtrees[:i], t.Subtrees[i+1:]...)
			break
		}
	}
}

func (t *Tree) String() string {
	return tree.SprintHrn(t)
}

// DebugTree is a debug tree node. Can be printed to help tracing the
// parsing flow.
type DebugTree struct {
	data     string
	subtrees []*DebugTree
}

func newDebugTree(data string) *DebugTree {
	return &DebugTree{
		data:     data,
		subtrees: []*DebugTree{},
	}
}

func (dt *DebugTree) add(subtree *DebugTree) {
	dt.subtrees = append(dt.subtrees, subtree)
}

func (dt *DebugTree) Data() interface{} {
	return dt.data
}

func (dt *DebugTree) Children() (c []tree.Node) {
	for _, child := range dt.subtrees {
		c = append(c, child)
	}
	return
}

func (dt *DebugTree) String() string {
	return tree.SprintHrn(dt)
}
