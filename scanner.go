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

package eesl

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// Token is the tuple returned from the scanner.
type Token struct {
	Line  int
	Text  []byte
	Error error
}

// Scan returns all the tokens in the input buffer.
// Errors such as unterminated blocks are included in the token.
// Nil is returned only if the input buffer is empty.
func Scan(buffer []byte, ignoreSpaces, ignoreComments bool) []Token {
	var delims = []byte{'+', '-', '*', '/', '%', '^', '#', ',', '=', '<', '>', '(', ')', '[', ']', '{', '}', '.', ':', ';'}

	line, tokens := 1, []Token{}

	for len(buffer) != 0 {
		t := Token{Line: line}
		switch ch := buffer[0]; ch {
		case '\n':
			t.Text, buffer = []byte{ch}, buffer[1:]
			line++
			if !ignoreSpaces {
				tokens = append(tokens, t)
			}
			continue
		case '+', '*', '/', '%', '^', '#', ',', '(', ')', '[', ']', '{', '}', ':', ';':
			t.Text, buffer = []byte{ch}, buffer[1:]
			tokens = append(tokens, t)
			continue
		case '-':
			if len(buffer) == 0 || buffer[0] != '-' {
				t.Text, buffer = []byte{'-'}, buffer[1:]
				tokens = append(tokens, t)
			} else if bytes.HasPrefix(buffer, []byte{'-', '-', '[', '['}) { // found block comment
				eoc := bytes.Index(buffer, []byte{'-', '-', ']', ']'})
				if eoc == -1 {
					// error, unterminated block comment
					t.Text, buffer = buffer, nil
					t.Error = ErrUnterminated
				} else {
					t.Text, buffer = buffer[:eoc+4], buffer[eoc+4:]
				}
				line += bytes.Count(t.Text, []byte{'\n'})
				if !ignoreComments {
					tokens = append(tokens, t)
				}
			} else { // found comment!
				start, length := buffer, 0
				for len(buffer) != 0 && buffer[0] != '\n' {
					buffer, length = buffer[1:], length+1
				}
				t.Text = start[:length]
				if !ignoreComments {
					tokens = append(tokens, t)
				}
			}
			continue
		case '<', '>', '=', '~':
			t.Text, buffer = []byte{ch}, buffer[1:]
			if len(buffer) != 0 && buffer[0] == '=' {
				t.Text, buffer = []byte{ch, '='}, buffer[1:]
			}
			tokens = append(tokens, t)
			continue
		case '.': // can be . or .. or ...
			if bytes.HasPrefix(buffer, []byte{'.', '.', '.'}) {
				t.Text, buffer = []byte{'.', '.', '.'}, buffer[3:]
			} else if bytes.HasPrefix(buffer, []byte{'.', '.'}) {
				t.Text, buffer = []byte{'.', '.'}, buffer[2:]
			} else {
				t.Text, buffer = []byte{'.'}, buffer[1:]
			}
			tokens = append(tokens, t)
			continue
		}

		start, length := buffer, 0
		r, w := utf8.DecodeRune(buffer)
		buffer, length = buffer[w:], length+w

		// spaces
		if r == utf8.RuneError || unicode.IsSpace(r) {
			r, w = utf8.DecodeRune(buffer)
			for len(buffer) != 0 && (r == utf8.RuneError || unicode.IsSpace(r)) {
				buffer, length = buffer[w:], length+w
				r, w = utf8.DecodeRune(buffer)
			}
			t.Text = start[:length]
			if !ignoreSpaces {
				tokens = append(tokens, t)
			}
			continue
		}

		// quoted text
		if qm := start[0]; qm == '"' || qm == '\'' {
			for len(buffer) != 0 && buffer[0] != qm {
				if buffer[0] == '\\' && len(buffer) > 1 {
					buffer, length = buffer[1:], length+1
				}
				if buffer[0] == '\n' {
					line++
				}
				buffer, length = buffer[1:], length+1
			}
			if len(buffer) == 0 || buffer[0] != qm {
				// unterminated quoted text
				t.Error = ErrUnterminated
				buffer = nil
			} else {
				buffer, length = buffer[1:], length+1
			}
			t.Text = start[:length]
			tokens = append(tokens, t)
			continue
		}

		// other tokens
		for len(buffer) != 0 && bytes.IndexByte(delims, buffer[0]) == -1 {
			r, w = utf8.DecodeRune(buffer)
			if r == utf8.RuneError || unicode.IsSpace(r) {
				// treat bad UTF-8 and spaces as delimiters
				break
			}
			buffer, length = buffer[w:], length+w
		}
		t.Text = start[:length]
		tokens = append(tokens, t)
	}

	return tokens
}
