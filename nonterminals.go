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

func chunk(tokens []Token) bool {
	return block(tokens)
}

func block(tokens []Token) bool {
	return sequence(zeroOrMore(stat), zeroOrOne(retstat))
}

func stat(tokens []Token) bool {
	if SemiColon(tokens) {
		return true
	}
	if sequence(varlist, Equals, explist) {
		return true
	}
	if functioncall(tokens) {
		return true
	}
	if label(tokens) {
		return true
	}
	if Break(tokens) {
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

func attnamelist(tokens []Token) bool {
	return sequence(Name, attrib, zeroOrOne(Comma, Name, attrib))
}

func attrib(tokens []Token) bool {
	return zeroOrOne(LessThan, Name, GreaterThan)(tokens)
}

func retstat(tokens []Token) bool {
	return sequence(Return, zeroOrOne(explist), zeroOrOne(SemiColon))
}

func label(tokens []Token) bool {
	return sequence(ColonColon, Name, ColonColon)
}

func funcname(tokens []Token) bool {
	return sequence(Name, zeroOrMore(Dot, Name), zeroOrOne(Colon, Name))
}

func varlist(tokens []Token) bool {
	return sequence(variable, zeroOrMore(Comma, variable))
}

func variable(tokens []Token) bool {
	if Name(tokens) {
		return true
	}
	if sequence(prefixexp, OpenBracket, exp, CloseBracket) {
		return true
	}
	return sequence(prefixexp, Dot, Name)
}

func namelist(tokens []Token) bool {
	return sequence(Name, zeroOrMore(Comma, Name))
}

func explist(tokens []Token) bool {
	return sequence(exp, zeroOrMore(Comma, Name))
}

func exp(tokens []Token) bool {
	if Nil(tokens) {
		return true
	}
	if False(tokens) {
		return true
	}
	if True(tokens) {
		return true
	}
	if Numeral(tokens) {
		return true
	}
	if LiteralString(tokens) {
		return true
	}
	if DotDotDot(tokens) {
		return true
	}
	if functiondef(tokens) {
		return true
	}
	if prefixexp(tokens) {
		return true
	}
	if tableconstructor(tokens) {
		return true
	}
	if sequence(exp, binop, exp) {
		return true
	}
	return sequence(unop, exp)
}

func prefixexp(tokens []Token) bool {
	if variable(tokens) {
		return true
	}
	if functioncall(tokens) {
		return true
	}
	return sequence(OpenParen, exp, CloseParen)
}

func functioncall(tokens []Token) bool {
	if sequence(prefixexp, args) {
		return true
	}
	return sequence(prefixexp, Colon, Name, args)
}

func args(tokens []Token) bool {
	if sequence(OpenParen, zeroOrOne(explist), CloseParen) {
		return true
	}
	if tableconstructor(tokens) {
		return true
	}
	return LiteralString(tokens)
}

func functiondef(tokens []Token) bool {
	return sequence(Function, funcbody)
}

func funcbody(tokens []Token) bool {
	return sequence(OpenParen, zeroOrOne(parlist), CloseParen, block, End)
}

func parlist(tokens []Token) bool {
	if sequence(namelist, zeroOrOne(Comma, DotDotDot)) {
		return true
	}
	return DotDotDot(tokens)
}

func tableconstructor(tokens []Token) bool {
	return sequence(OpenCurly, zeroOrOne(fieldlist), CloseCurly)
}

func fieldlist(tokens []Token) bool {
	return sequence(field, zeroOrMore(fieldsep, field), zeroOrOne(fieldsep))
}

func field(tokens []Token) bool {
	if sequence(OpenBracket, exp, CloseBracket, Equals, exp) {
		return true
	}
	if sequence(Name, Equals, exp) {
		return true
	}
	return exp(tokens)
}

func fieldsep(tokens []Token) bool {
	return or(Comma, SemiColon)
}

func binop(tokens []Token) bool {
	return or(Plus, Hyphen, Star, Slash, SlashSlash, Caret, Percent, Ampersand, Tilde, Bar, GreaterThanGreaterThan, LessThanLessThan, DotDot, LessThan, LessThanEquals, GreaterThan, GreaterThanEquals, EqualsEquals, TildeEquals, And, Or)
}

func unop(tokens []Token) bool {
	return or(Hyphen, Not, Hash, Tilde)
}
