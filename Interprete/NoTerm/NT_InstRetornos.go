package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_InstBreak struct {
}

// Constructor for NT_InstBreak
func NewNT_InstBreak() *NT_InstBreak {
	o := new(NT_InstBreak)
	return o
}

func (ntBreak *NT_InstBreak) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewBreakLiteral()
}

type NT_InstContinue struct {
}

// Constructor for NT_InstContinue
func NewNT_InstContinue() *NT_InstContinue {
	o := new(NT_InstContinue)
	return o
}

func (ntContinue *NT_InstContinue) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewContinueLiteral()
}

type NT_InstReturn struct {
	exp interprete.AbstrExpr
}

// Constructor for NT_InstReturn
func NewNT_InstReturn(exp interprete.AbstrExpr) *NT_InstReturn {
	o := new(NT_InstReturn)
	o.exp = exp
	return o
}
func (ntReturn *NT_InstReturn) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	if ntReturn != nil {
		exp := ntReturn.exp.Interpretar(ctx)
		if exp.Tipo == interprete.Integer {
			tmp := interprete.NewIntLiteral(exp.Valor)
			tmp.Return = true
			return tmp
		} else if exp.Tipo == interprete.Character {
			tmp := interprete.NewCharLiteral(exp.ValorC)
			tmp.Return = true
			return tmp
		} else if exp.Tipo == interprete.Bool {
			tmp := interprete.NewBoolLiteral(exp.ValorB)
			tmp.Return = true
			return tmp
		} else if exp.Tipo == interprete.Float {
			tmp := interprete.NewFloatLiteral(exp.ValorF)
			tmp.Return = true
			return tmp
		} else if exp.Tipo == interprete.String {
			tmp := interprete.NewStringLiteral(exp.ValorS)
			tmp.Return = true
			return tmp
		}
		return interprete.NewNil()
	} else {
		tmp := interprete.NewNil()
		tmp.Return = true
		return tmp
	}
}
