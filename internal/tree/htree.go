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

package tree

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// PrintHr prints the horizontal formatted tree to standard output.
func PrintHr(root Node) {
	fmt.Print(SprintHr(root))
}

// SprintHr returns the horizontal formatted tree.
func SprintHr(root Node) (s string) {
	for _, line := range lines(root) {
		// ignore runes before root node
		line = string([]rune(line)[2:])
		s += strings.TrimRight(line, " ") + "\n"
	}
	return
}

func lines(root Node) (s []string) {
	data := fmt.Sprintf("%s %v ", BoxHor, root.Data())
	l := len(root.Children())
	if l == 0 {
		s = append(s, data)
		return
	}

	w := utf8.RuneCountInString(data)
	for i, c := range root.Children() {
		for j, line := range lines(c) {
			if i == 0 && j == 0 {
				if l == 1 {
					s = append(s, data+BoxHor+line)
				} else {
					s = append(s, data+BoxDownHor+line)
				}
				continue
			}

			var box string
			if i == l-1 && j == 0 {
				// first line of the last child
				box = BoxUpRight
			} else if i == l-1 {
				box = " "
			} else if j == 0 {
				box = BoxVerRight
			} else {
				box = BoxVer
			}
			s = append(s, strings.Repeat(" ", w)+box+line)
		}
	}
	return
}

// PrintHrn prints the horizontal-newline formatted tree to standard output.
func PrintHrn(root Node) {
	fmt.Print(SprintHrn(root))
}

// SprintHrn returns the horizontal-newline formatted tree.
func SprintHrn(root Node) (s string) {
	return strings.Join(lines2(root), "\n") + "\n"
}

func lines2(root Node) (s []string) {
	s = append(s, fmt.Sprintf("%v", root.Data()))
	l := len(root.Children())
	if l == 0 {
		return
	}

	for i, c := range root.Children() {
		for j, line := range lines2(c) {
			// first line of the last child
			if i == l-1 && j == 0 {
				s = append(s, BoxUpRight+BoxHor+" "+line)
			} else if j == 0 {
				s = append(s, BoxVerRight+BoxHor+" "+line)
			} else if i == l-1 {
				s = append(s, "   "+line)
			} else {
				s = append(s, BoxVer+"  "+line)
			}
		}
	}
	return
}
