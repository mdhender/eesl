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

//type Terminal int
//
//const (
//	Ampersand Terminal = iota
//	And
//	Bar
//	Break
//	Caret
//	CloseBracket
//	CloseCurly
//	CloseParen
//	Colon
//	ColonColon
//	Comma
//	Do
//	Dot
//	DotDot
//	DotDotDot
//	Else
//	ElseIf
//	End
//	Equals
//	EqualsEquals
//	False
//	For
//	Function
//	Goto
//	GreaterThan
//	GreaterThanEquals
//	GreaterThanGreaterThan
//	Hash
//	Hyphen
//	If
//	In
//	LessThan
//	LessThanEquals
//	LessThanLessThan
//	LiteralString
//	Local
//	Name
//	Nil
//	Not
//	Numeral
//	OpenBracket
//	OpenCurly
//	OpenParen
//	Or
//	Percent
//	Plus
//	Repeat
//	Return
//	SemiColon
//	Slash
//	SlashSlash
//	Star
//	Then
//	Tilde
//	TildeEquals
//	True
//	Until
//	While
//)

/*
func tAmpersand(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'&'}) }
func tAnd(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'a', 'n', 'd'}) }
func tBar(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'|'}) }
func tBreak(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'b', 'r', 'e', 'a', 'k'})
}
func tCaret(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'^'}) }
func tCloseBracket(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{']'}) }
func tCloseCurly(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{'}'}) }
func tCloseParen(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{')'}) }
func tColon(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{':'}) }
func tColonColon(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{':', ':'}) }
func tComma(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{','}) }
func tDo(b *builder.Builder) bool           { return bytes.Equal(tokens[0].Text, []byte{'d', 'o'}) }
func tDot(b *builder.Builder) bool          { return bytes.Equal(tokens[0].Text, []byte{'.'}) }
func tDotDot(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'.', '.'}) }
func tDotDotDot(b *builder.Builder) bool    { return bytes.Equal(tokens[0].Text, []byte{'.', '.', '.'}) }
func tElse(b *builder.Builder) bool         { return bytes.Equal(tokens[0].Text, []byte{'e', 'l', 's', 'e'}) }
func tElseIf(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'e', 'l', 's', 'e', 'i', 'f'})
}
func tEnd(b *builder.Builder) bool          { return bytes.Equal(tokens[0].Text, []byte{'e', 'n', 'd'}) }
func tEquals(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'='}) }
func tEqualsEquals(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'=', '='}) }
func tFalse(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'f', 'a', 'l', 's', 'e'})
}
func tFor(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'f', 'o', 'r'}) }
func tFunction(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'f', 'u', 'n', 'c', 't', 'i', 'o', 'n'})
}
func tGoto(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'g', 'o', 't', 'o'}) }
func tGreaterThan(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'>'}) }
func tGreaterThanEquals(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'>', '='})
}
func tGreaterThanGreaterThan(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'>', '>'})
}
func tHash(b *builder.Builder) bool             { return bytes.Equal(tokens[0].Text, []byte{'#'}) }
func tHyphen(b *builder.Builder) bool           { return bytes.Equal(tokens[0].Text, []byte{'-'}) }
func tIf(b *builder.Builder) bool               { return bytes.Equal(tokens[0].Text, []byte{'i', 'f'}) }
func tIn(b *builder.Builder) bool               { return bytes.Equal(tokens[0].Text, []byte{'i', 'n'}) }
func tLessThan(b *builder.Builder) bool         { return bytes.Equal(tokens[0].Text, []byte{'<'}) }
func tLessThanEquals(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{'<', '='}) }
func tLessThanLessThan(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'<', '<'}) }
func tLiteralString(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'"', '.', '.', '.', '"'})
}
func tLocal(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'l', 'o', 'c', 'a', 'l'})
}
func tName(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'n', 'a', 'm', 'e'}) }
func tNil(b *builder.Builder) bool  { return bytes.Equal(tokens[0].Text, []byte{'n', 'i', 'l'}) }
func tNot(b *builder.Builder) bool  { return bytes.Equal(tokens[0].Text, []byte{'n', 'o', 't'}) }
func tNumeral(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'0', '.', '.', '.', '9'})
}
func tOpenBracket(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'['}) }
func tOpenCurly(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{'{'}) }
func tOpenParen(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{'('}) }
func tOr(b *builder.Builder) bool          { return bytes.Equal(tokens[0].Text, []byte{'o', 'r'}) }
func tPercent(b *builder.Builder) bool     { return bytes.Equal(tokens[0].Text, []byte{'%'}) }
func tPlus(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'+'}) }
func tRepeat(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'r', 'e', 'p', 'e', 'a', 't'})
}
func tReturn(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'r', 'e', 't', 'u', 'r', 'n'})
}
func tSemiColon(b *builder.Builder) bool   { return bytes.Equal(tokens[0].Text, []byte{';'}) }
func tSlash(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'/'}) }
func tSlashSlash(b *builder.Builder) bool  { return bytes.Equal(tokens[0].Text, []byte{'/', '/'}) }
func tStar(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'*'}) }
func tThen(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'t', 'h', 'e', 'n'}) }
func tTilde(b *builder.Builder) bool       { return bytes.Equal(tokens[0].Text, []byte{'~'}) }
func tTildeEquals(b *builder.Builder) bool { return bytes.Equal(tokens[0].Text, []byte{'~', '='}) }
func tTrue(b *builder.Builder) bool        { return bytes.Equal(tokens[0].Text, []byte{'t', 'r', 'u', 'e'}) }
func tUntil(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'u', 'n', 't', 'i', 'l'})
}
func tWhile(b *builder.Builder) bool {
	return bytes.Equal(tokens[0].Text, []byte{'w', 'h', 'i', 'l', 'e'})
}
*/
