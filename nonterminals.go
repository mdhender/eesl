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
	"github.com/mdhender/eesl/internal/builder"
	"github.com/mdhender/eesl/internal/tokenizer"
)

func ntChunk(b *builder.Builder) (ok bool) {
	b.Enter("chunk")
	defer b.Exit(&ok)

	return ntBlock(b)
}

func ntBlock(b *builder.Builder) (ok bool) {
	b.Enter("block")
	defer b.Exit(&ok)

	for ntStat(b) {
		// ... //
	}
	if ntRetstat(b) {
		// ... //
	}
	return true
}

func ntStat(b *builder.Builder) (ok bool) {
	b.Enter("stat")
	defer b.Exit(&ok)

	if b.Match(tokenizer.SemiColon) {
		// ... //
		return true
	}
	if ntVarlist(b) && b.Match(tokenizer.Equals) && ntExplist(b) {
		// ... //
		return true
	}
	if ntFunctioncall(b) {
		// ... //
		return true
	}
	if ntLabel(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Break) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Goto) && b.Match(tokenizer.Name) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Do) && ntBlock(b) && b.Match(tokenizer.End) {
		// ... //
		return true
	}
	if b.Match(tokenizer.While) && ntExp(b) && b.Match(tokenizer.Do) && ntBlock(b) && b.Match(tokenizer.End) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Repeat) && ntBlock(b) && b.Match(tokenizer.Until) && ntExp(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.If) && ntExp(b) && b.Match(tokenizer.Then) && ntBlock(b) {
		for b.Match(tokenizer.ElseIf) && ntExp(b) && b.Match(tokenizer.Then) && ntBlock(b) {
			// ... //
		}
		if b.Match(tokenizer.Else) && ntBlock(b) {
			// ... //
		}
		if b.Match(tokenizer.End) {
			// ... //
			return true
		}
		return false
	}
	if b.Match(tokenizer.For) && b.Match(tokenizer.Name) && b.Match(tokenizer.Equals) && ntExp(b) && b.Match(tokenizer.Comma) && ntExp(b) {
		if b.Match(tokenizer.Comma) && ntExp(b) {
			// ... //
		}
		if b.Match(tokenizer.Do) && ntBlock(b) && b.Match(tokenizer.End) {
			// ... //
			return true
		}
		return false
	}
	if b.Match(tokenizer.For) && ntNamelist(b) && b.Match(tokenizer.In) && ntExplist(b) && b.Match(tokenizer.Do) && ntBlock(b) && b.Match(tokenizer.End) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Function) && ntFuncname(b) && ntFuncbody(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Local) && b.Match(tokenizer.Function) && b.Match(tokenizer.Name) && ntFuncbody(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Local) && ntAttnamelist(b) {
		// ... //
		if b.Match(tokenizer.Equals) && ntExplist(b) {
			// ... //
		}
		return true
	}
	return false
}

func ntAttnamelist(b *builder.Builder) (ok bool) {
	b.Enter("attnamelist")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Name) && ntAttrib(b) {
		// ... //
		for b.Match(tokenizer.Comma) && b.Match(tokenizer.Name) && ntAttrib(b) {
			// ... //
		}
	}
	return false
}

func ntAttrib(b *builder.Builder) (ok bool) {
	b.Enter("attrib")
	defer b.Exit(&ok)

	if b.Match(tokenizer.LessThan) && b.Match(tokenizer.Name) && b.Match(tokenizer.GreaterThan) {
		// ... //
		return true
	}
	// epsilon
	return true
}

func ntRetstat(b *builder.Builder) (ok bool) {
	b.Enter("retstat")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Return) {
		// ... //
		if ntExplist(b) {
			// ... //
		}
		if b.Match(tokenizer.SemiColon) {
			// ... //
		}
		return true
	}
	return false
}

func ntLabel(b *builder.Builder) (ok bool) {
	b.Enter("label")
	defer b.Exit(&ok)

	if b.Match(tokenizer.ColonColon) && b.Match(tokenizer.Name) && b.Match(tokenizer.ColonColon) {
		// ... //
		return true
	}
	return false
}

func ntFuncname(b *builder.Builder) (ok bool) {
	b.Enter("funcname")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Name) {
		// ... //
		for b.Match(tokenizer.Dot) && b.Match(tokenizer.Name) {
			// ... //
		}
		if b.Match(tokenizer.Colon) && b.Match(tokenizer.Name) {
			// ... //
		}
		return true
	}
	return false
}

func ntVarlist(b *builder.Builder) (ok bool) {
	b.Enter("varlist")
	defer b.Exit(&ok)

	if ntVar(b) {
		// ... //
		for b.Match(tokenizer.Comma) && ntVar(b) {
			// ... //
		}
		return true
	}
	return false
}

func ntVar(b *builder.Builder) (ok bool) {
	b.Enter("var")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Name) {
		// ... //
		return true
	}
	if ntPrefixexp(b) && b.Match(tokenizer.OpenBracket) && ntExp(b) && b.Match(tokenizer.CloseBracket) {
		// ... //
		return true
	}
	if ntPrefixexp(b) && b.Match(tokenizer.Dot) && b.Match(tokenizer.Name) {
		// ... //
		return true
	}
	return false
}

func ntNamelist(b *builder.Builder) (ok bool) {
	b.Enter("namelist")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Name) {
		// ... //
		for b.Match(tokenizer.Comma) && b.Match(tokenizer.Name) {
			// ... //
		}
		return true
	}
	return false
}

func ntExplist(b *builder.Builder) (ok bool) {
	b.Enter("explist")
	defer b.Exit(&ok)

	if ntExp(b) {
		// ... //
		for b.Match(tokenizer.Comma) && b.Match(tokenizer.Name) {
			// ... ///
		}
		return true
	}
	return false
}

func ntExp(b *builder.Builder) (ok bool) {
	b.Enter("exp")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Nil) {
		// ... //
		return true
	}
	if b.Match(tokenizer.False) {
		// ... //
		return true
	}
	if b.Match(tokenizer.True) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Numeral) {
		// ... //
		return true
	}
	if b.Match(tokenizer.LiteralString) {
		// ... //
		return true
	}
	if b.Match(tokenizer.DotDotDot) {
		// ... //
		return true
	}
	if ntFunctiondef(b) {
		// ... //
		return true
	}
	if ntPrefixexp(b) {
		// ... //
		return true
	}
	if ntTableconstructor(b) {
		// ... //
		return true
	}
	if ntExp(b) && ntBinop(b) && ntExp(b) {
		// ... //
		return true
	}
	if ntUnop(b) && ntExp(b) {
		// ... //
		return true
	}
	return false
}

func ntPrefixexp(b *builder.Builder) (ok bool) {
	b.Enter("prefixexp")
	defer b.Exit(&ok)

	if ntVar(b) {
		// ... //
		return true
	}
	if ntFunctioncall(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.OpenParen) && ntExp(b) && b.Match(tokenizer.CloseParen) {
		// ... //
		return true
	}
	return false
}

func ntFunctioncall(b *builder.Builder) (ok bool) {
	b.Enter("functioncall")
	defer b.Exit(&ok)

	if ntPrefixexp(b) && ntArgs(b) {
		// ... //
		return true
	}
	if ntPrefixexp(b) && b.Match(tokenizer.Colon) && b.Match(tokenizer.Name) && ntArgs(b) {
		// ... //
		return true
	}
	return false
}

func ntArgs(b *builder.Builder) (ok bool) {
	b.Enter("args")
	defer b.Exit(&ok)

	if b.Match(tokenizer.OpenParen) {
		// ... //
		if ntExplist(b) {
			// ... //
		}
		if b.Match(tokenizer.CloseParen) {
			// ... //
			return true
		}
		return false
	}
	if ntTableconstructor(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.LiteralString) {
		// ... //
		return true
	}
	return false
}

func ntFunctiondef(b *builder.Builder) (ok bool) {
	b.Enter("functiondef")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Function) && ntFuncbody(b) {
		// ... //
		return true
	}
	return false
}

func ntFuncbody(b *builder.Builder) (ok bool) {
	b.Enter("funcbody")
	defer b.Exit(&ok)

	if b.Match(tokenizer.OpenParen) {
		// ... //
		if ntParlist(b) {
			// ... //
		}
		if b.Match(tokenizer.CloseParen) && ntBlock(b) && b.Match(tokenizer.End) {
			// ... //
			return true
		}
		return false
	}
	return false
}

func ntParlist(b *builder.Builder) (ok bool) {
	b.Enter("parlist")
	defer b.Exit(&ok)

	if ntNamelist(b) {
		// ... //
		if b.Match(tokenizer.Comma) && b.Match(tokenizer.DotDotDot) {
			// ... //
		}
		return true
	}
	if b.Match(tokenizer.DotDotDot) {
		// ... //
		return true
	}
	return false
}

func ntTableconstructor(b *builder.Builder) (ok bool) {
	b.Enter("tableconstructor")
	defer b.Exit(&ok)

	if b.Match(tokenizer.OpenCurly) && ntFieldlist(b) && b.Match(tokenizer.CloseCurly) {
		// ... //
		return true
	}
	return false
}

func ntFieldlist(b *builder.Builder) (ok bool) {
	b.Enter("fieldlist")
	defer b.Exit(&ok)

	if ntField(b) {
		// ... //
		for ntFieldsep(b) && ntField(b) {
			// ... //
		}
		if ntFieldsep(b) {
			// ... //
		}
		return true
	}
	return false
}

func ntField(b *builder.Builder) (ok bool) {
	b.Enter("field")
	defer b.Exit(&ok)

	if b.Match(tokenizer.OpenBracket) && ntExp(b) && b.Match(tokenizer.CloseBracket) && b.Match(tokenizer.Equals) && ntExp(b) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Name) && b.Match(tokenizer.Equals) && ntExp(b) {
		// ... //
		return true
	}
	if ntExp(b) {
		// ... //
		return true
	}
	return false
}

func ntFieldsep(b *builder.Builder) (ok bool) {
	b.Enter("fieldsep")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Comma) {
		// ... //
		return true
	}
	if b.Match(tokenizer.SemiColon) {
		// ... //
		return true
	}
	return false
}

func ntBinop(b *builder.Builder) (ok bool) {
	b.Enter("binop")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Plus) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Hyphen) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Star) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Slash) {
		// ... //
		return true
	}
	if b.Match(tokenizer.SlashSlash) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Caret) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Percent) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Ampersand) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Tilde) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Bar) {
		// ... //
		return true
	}
	if b.Match(tokenizer.GreaterThanGreaterThan) {
		// ... //
		return true
	}
	if b.Match(tokenizer.LessThanLessThan) {
		// ... //
		return true
	}
	if b.Match(tokenizer.DotDot) {
		// ... //
		return true
	}
	if b.Match(tokenizer.LessThan) {
		// ... //
		return true
	}
	if b.Match(tokenizer.LessThanEquals) {
		// ... //
		return true
	}
	if b.Match(tokenizer.GreaterThan) {
		// ... //
		return true
	}
	if b.Match(tokenizer.GreaterThanEquals) {
		// ... //
		return true
	}
	if b.Match(tokenizer.EqualsEquals) {
		// ... //
		return true
	}
	if b.Match(tokenizer.TildeEquals) {
		// ... //
		return true
	}
	if b.Match(tokenizer.And) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Or) {
		// ... //
		return true
	}
	return false
}

func ntUnop(b *builder.Builder) (ok bool) {
	b.Enter("unop")
	defer b.Exit(&ok)

	if b.Match(tokenizer.Hyphen) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Not) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Hash) {
		// ... //
		return true
	}
	if b.Match(tokenizer.Tilde) {
		// ... //
		return true
	}
	return false
}
