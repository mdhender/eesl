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

import "bytes"

func Ampersand(tokens []Token) bool    { return bytes.Equal(tokens[0].Text, []byte{'&'}) }
func And(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'a', 'n', 'd'}) }
func Bar(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'|'}) }
func Break(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'b', 'r', 'e', 'a', 'k'}) }
func Caret(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'^'}) }
func CloseBracket(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{']'}) }
func CloseCurly(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'}'}) }
func CloseParen(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{')'}) }
func Colon(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{':'}) }
func ColonColon(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{':', ':'}) }
func Comma(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{','}) }
func Do(tokens []Token) bool           { return bytes.Equal(tokens[0].Text, []byte{'d', 'o'}) }
func Dot(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'.'}) }
func DotDot(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'.', '.'}) }
func DotDotDot(tokens []Token) bool    { return bytes.Equal(tokens[0].Text, []byte{'.', '.', '.'}) }
func Else(tokens []Token) bool         { return bytes.Equal(tokens[0].Text, []byte{'e', 'l', 's', 'e'}) }
func ElseIf(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'e', 'l', 's', 'e', 'i', 'f'})
}
func End(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'e', 'n', 'd'}) }
func Equals(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'='}) }
func EqualsEquals(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'=', '='}) }
func False(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'f', 'a', 'l', 's', 'e'}) }
func For(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'f', 'o', 'r'}) }
func Function(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'f', 'u', 'n', 'c', 't', 'i', 'o', 'n'})
}
func Goto(tokens []Token) bool              { return bytes.Equal(tokens[0].Text, []byte{'g', 'o', 't', 'o'}) }
func GreaterThan(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'>'}) }
func GreaterThanEquals(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'>', '='}) }
func GreaterThanGreaterThan(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'>', '>'})
}
func Hash(tokens []Token) bool             { return bytes.Equal(tokens[0].Text, []byte{'#'}) }
func Hyphen(tokens []Token) bool           { return bytes.Equal(tokens[0].Text, []byte{'-'}) }
func If(tokens []Token) bool               { return bytes.Equal(tokens[0].Text, []byte{'i', 'f'}) }
func In(tokens []Token) bool               { return bytes.Equal(tokens[0].Text, []byte{'i', 'n'}) }
func LessThan(tokens []Token) bool         { return bytes.Equal(tokens[0].Text, []byte{'<'}) }
func LessThanEquals(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'<', '='}) }
func LessThanLessThan(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'<', '<'}) }
func LiteralString(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'"', '.', '.', '.', '"'})
}
func Local(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'l', 'o', 'c', 'a', 'l'}) }
func Name(tokens []Token) bool  { return bytes.Equal(tokens[0].Text, []byte{'n', 'a', 'm', 'e'}) }
func Nil(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'n', 'i', 'l'}) }
func Not(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'n', 'o', 't'}) }
func Numeral(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'0', '.', '.', '.', '9'})
}
func OpenBracket(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'['}) }
func OpenCurly(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'{'}) }
func OpenParen(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{'('}) }
func Or(tokens []Token) bool          { return bytes.Equal(tokens[0].Text, []byte{'o', 'r'}) }
func Percent(tokens []Token) bool     { return bytes.Equal(tokens[0].Text, []byte{'%'}) }
func Plus(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'+'}) }
func Repeat(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'r', 'e', 'p', 'e', 'a', 't'})
}
func Return(tokens []Token) bool {
	return bytes.Equal(tokens[0].Text, []byte{'r', 'e', 't', 'u', 'r', 'n'})
}
func SemiColon(tokens []Token) bool   { return bytes.Equal(tokens[0].Text, []byte{';'}) }
func Slash(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'/'}) }
func SlashSlash(tokens []Token) bool  { return bytes.Equal(tokens[0].Text, []byte{'/', '/'}) }
func Star(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'*'}) }
func Then(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'t', 'h', 'e', 'n'}) }
func Tilde(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'~'}) }
func TildeEquals(tokens []Token) bool { return bytes.Equal(tokens[0].Text, []byte{'~', '='}) }
func True(tokens []Token) bool        { return bytes.Equal(tokens[0].Text, []byte{'t', 'r', 'u', 'e'}) }
func Until(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'u', 'n', 't', 'i', 'l'}) }
func While(tokens []Token) bool       { return bytes.Equal(tokens[0].Text, []byte{'w', 'h', 'i', 'l', 'e'}) }
