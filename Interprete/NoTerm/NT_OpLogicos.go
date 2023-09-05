package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_LogAnd struct {
	expIzq interprete.AbstrExpr
	expDer interprete.AbstrExpr
}

// Constructor for NT_LogAnd
func NewNT_LogAnd(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr) *NT_LogAnd {
	o := new(NT_LogAnd)
	o.expIzq = expIzq
	o.expDer = expDer
	return o
}

func (ntAnd *NT_LogAnd) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntAnd.expIzq.Interpretar(ctx)
	expDer := ntAnd.expDer.Interpretar(ctx)

	if expIzq.Tipo != interprete.Bool || expDer.Tipo != interprete.Bool {
		ctx.AddError("Error: Tipos incompatibles para operador logico")
		return interprete.NewNil()
	}

	if expIzq.ValorB && expDer.ValorB {
		return interprete.NewBoolLiteral(true)
	} else {
		return interprete.NewBoolLiteral(false)
	}
}

type NT_LogOR struct {
	expIzq interprete.AbstrExpr
	expDer interprete.AbstrExpr
}

// Constructor for NT_LogOR
func NewNT_LogOR(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr) *NT_LogOR {
	o := new(NT_LogOR)
	o.expIzq = expIzq
	o.expDer = expDer
	return o
}

func (ntOR *NT_LogOR) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntOR.expIzq.Interpretar(ctx)
	expDer := ntOR.expDer.Interpretar(ctx)

	if expIzq.Tipo != interprete.Bool || expDer.Tipo != interprete.Bool {
		ctx.AddError("Error: Tipos incompatibles para operador logico")
		return interprete.NewNil()
	}

	if expIzq.ValorB || expDer.ValorB {
		return interprete.NewBoolLiteral(true)
	} else {
		return interprete.NewBoolLiteral(false)
	}
}

type NT_LogNot struct {
	exp interprete.AbstrExpr
}

// Constructor for NT_LogNot
func NewNT_LogNot(exp interprete.AbstrExpr) *NT_LogNot {
	o := new(NT_LogNot)
	o.exp = exp
	return o
}

func (ntNot *NT_LogNot) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntNot.exp.Interpretar(ctx)
	if exp.Tipo != interprete.Bool {
		ctx.AddError("Error: Tipo incompatible para operador logico")
		return interprete.NewNil()
	}

	if exp.ValorB {
		return interprete.NewBoolLiteral(false)
	} else {
		return interprete.NewBoolLiteral(true)
	}
}
