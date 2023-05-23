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

package tokenizer_test

import (
	"github.com/mdhender/eesl/internal/tokenizer"
	"testing"
)

func TestScanner(t *testing.T) {
	input := []byte(`
		-- defines a factorial function
		function fact (n)
		  if n == 0 then
			return 1
		  else
			return n * fact(n-1)
		  end
		end
	
		print("enter a number:")
		a = io.read("*number")        -- read a number
		print(fact(a))
	`)

	for _, token := range tokenizer.Scan(input, false, false) {
		t.Logf("%3d: %-30q %v\n", token.Line, string(token.Text), token.Error)
	}

	input = []byte(`
		-- file 'lib1.lua'
	
		function norm (x, y)
		  local n2 = x^2 + y^2
		  return math.sqrt(n2)
		end
	
		function twice (x)
		  return 2*x
		end
	`)

	for _, token := range tokenizer.Scan(input, true, true) {
		t.Logf("%3d: %-30q %v\n", token.Line, string(token.Text), token.Error)
	}

	input = []byte(`
		--[[
		print(10)         -- no action (comment)
		--]]
	
		---[[
		print(10)         -- no action (comment)
		--]]
	`)

	for _, token := range tokenizer.Scan(input, true, true) {
		t.Logf("%3d: %-30q %v\n", token.Line, string(token.Text), token.Error)
	}

	input = []byte(`
		print(type("Hello world"))  --> string
		print(type(10.4*3))         --> number
		print(type(print))          --> function
		print(type(type))           --> function
		print(type(true))           --> boolean
		print(type(nil))            --> nil
		print(type(type(X)))        --> string
		print(type(a))              --> nil   (a is not initialized)
		a = print                   -- yes, this is valid!
		print(type(a))              --> number
	`)

	for _, token := range tokenizer.Scan(input, true, true) {
		t.Logf("%3d: %-30q %v\n", token.Line, string(token.Text), token.Error)
	}

	input = []byte(`
		print("one line\nnext line\n\"in quotes\", 'in quotes'")
		print('a backslash inside quotes: \'\\\'')
		print("a simpler way: '\\'")
	`)

	for _, token := range tokenizer.Scan(input, true, true) {
		t.Logf("%3d: %-30q %v\n", token.Line, string(token.Text), token.Error)
	}

	type expect struct {
		line int
		text string
	}
	for _, tc := range []struct {
		id     int
		input  string
		expect []expect
	}{
		{1, `"a"`, []expect{
			{1, `"a"`},
		}},
		{2, `'a'`, []expect{
			{1, `'a'`},
		}},
		{3, `"a\nb \"c\"\n'in quotes'"`, []expect{
			{1, `"a\nb \"c\"\n'in quotes'"`},
		}},
		{4, `'a\nb \'c\'\n"in quotes"'`, []expect{
			{1, `'a\nb \'c\'\n"in quotes"'`},
		}},
		{5, `print(type("Hello world"))`, []expect{
			{1, "print"},
			{1, "("},
			{1, "type"},
			{1, "("},
			{1, "\"Hello world\""},
			{1, ")"},
			{1, ")"},
		}},
		{6, `b = string.gsub(a, "one", "another")`, []expect{
			{1, "b"},
			{1, "="},
			{1, "string"},
			{1, "."},
			{1, "gsub"},
			{1, "("},
			{1, "a"},
			{1, ","},
			{1, `"one"`},
			{1, ","},
			{1, `"another"`},
			{1, ")"},
		}},
	} {
		tokens := tokenizer.Scan([]byte(tc.input), true, true)
		if len(tokens) != len(tc.expect) {
			t.Errorf("%d: want tokens: %d, got %d\n", tc.id, len(tc.expect), len(tokens))
		} else {
			for i := range tokens {
				if tokens[i].Error != nil {
					t.Errorf("%d: want err: nil, got %v\n", tc.id, tokens[i].Error)
				}
				if tokens[i].Line != tc.expect[i].line {
					t.Errorf("%d: want line: 1, got %d\n", tc.id, tokens[i].Line)
				}
				if string(tokens[i].Text) != tc.expect[i].text {
					t.Errorf("%d: want text: %q, got %q\n", tc.id, tc.expect, tokens[i].Text)
				}
			}
		}
	}
}
