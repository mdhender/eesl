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

func chunk() bool {
	return block()
}

func block() bool {
	return sequence(zeroOrMore(stat), zeroOrOne(retstat))
}

func stat() bool {
	if SemiColon() {
		return true
	}
	if sequence(varlist, Equals, explist) {
		return true
	}
	if functioncall() {
		return true
	}
	if label() {
		return true
	}
	if Break() {
		return true
	}
	if sequence(Goto, Name) {
		return true
	}
	if sequence(Do, block, End) {
		return true
	}
	if sequence(While, exp, Do, block, End) {
		return true
	}
	if sequence(Repeat, block, Until, exp) {
		return true
	}
	if sequence(If, exp, Then, block, zeroOrMore(ElseIf, exp, Then, block), zeroOrOne(Else, block), End) {
		return true
	}
	if sequence(For, Name, Equals, exp, Comma, exp, zeroOrOne(Comma, exp), Do, block, End) {
		return true
	}
	if sequence(For, namelist, In, explist, Do, block, End) {
		return true
	}
	if sequence(Function, funcname, funcbody) {
		return true
	}
	if sequence(Local, Function, Name, funcbody) {
		return true
	}
	return sequence(Local, attnamelist, zeroOrOne(Equals, explist))
}

func attnamelist() bool {
	return sequence(Name, attrib, zeroOrOne(Comma, Name, attrib))
}

func attrib() bool {
	return zeroOrOne(LessThan, Name, GreaterThan)()
}

func retstat() bool {
	return sequence(Return, zeroOrOne(explist), zeroOrOne(SemiColon))
}

func label() bool {
	return sequence(ColonColon, Name, ColonColon)
}

func funcname() bool {
	return sequence(Name, zeroOrMore(Dot, Name), zeroOrOne(Colon, Name))
}

func varlist() bool {
	return sequence(var_, zeroOrMore(Comma, var_))
}

func var_() bool {
	if Name() {
		return true
	}
	if sequence(prefixexp, OpenBracket, exp, CloseBracket) {
		return true
	}
	return sequence(prefixexp, Dot, Name)
}

func namelist() bool {
	return sequence(Name, zeroOrMore(Comma, Name))
}

func explist() bool {
	return sequence(exp, zeroOrMore(Comma, Name))
}

func exp() bool {
	if Nil() {
		return true
	}
	if False() {
		return true
	}
	if True() {
		return true
	}
	if Numeral() {
		return true
	}
	if LiteralString() {
		return true
	}
	if DotDotDot() {
		return true
	}
	if functiondef() {
		return true
	}
	if prefixexp() {
		return true
	}
	if tableconstructor() {
		return true
	}
	if sequence(exp, binop, exp) {
		return true
	}
	return sequence(unop, exp)
}

func prefixexp() bool {
	if var_() {
		return true
	}
	if functioncall() {
		return true
	}
	return sequence(OpenParen, exp, CloseParen)
}

func functioncall() bool {
	if sequence(prefixexp, args) {
		return true
	}
	return sequence(prefixexp, Colon, Name, args)
}

func args() bool {
	if sequence(OpenParen, zeroOrOne(explist), CloseParen) {
		return true
	}
	if tableconstructor() {
		return true
	}
	return LiteralString()
}

func functiondef() bool {
	return sequence(Function, funcbody)
}

func funcbody() bool {
	return sequence(OpenParen, zeroOrOne(parlist), CloseParen, block, End)
}

func parlist() bool {
	if sequence(namelist, zeroOrOne(Comma, DotDotDot)) {
		return true
	}
	return DotDotDot()
}

func tableconstructor() bool {
	return sequence(OpenCurly, zeroOrOne(fieldlist), CloseCurly)
}

func fieldlist() bool {
	return sequence(field, zeroOrMore(fieldsep, field), zeroOrOne(fieldsep))
}

func field() bool {
	if sequence(OpenBracket, exp, CloseBracket, Equals, exp) {
		return true
	}
	if sequence(Name, Equals, exp) {
		return true
	}
	return exp()
}

func fieldsep() bool {
	return or(Comma, SemiColon)
}

func binop() bool {
	return or(Plus, Hyphen, Star, Slash, SlashSlash, Caret, Percent, Ampersand, Tilde, Bar, GreaterThanGreaterThan, LessThanLessThan, DotDot, LessThan, LessThanEquals, GreaterThan, GreaterThanEquals, EqualsEquals, TildeEquals, And, Or)
}

func unop() bool {
	return or(Hyphen, Not, Hash, Tilde)
}
