package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_ExpMult struct {
	ExpIzq interprete.AbstrExpr
	ExpDer interprete.AbstrExpr
}

// Constructor for NT_ExpMult
func NewNT_ExpMult(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr) *NT_ExpMult {
	o := new(NT_ExpMult)
	o.ExpIzq = expIzq
	o.ExpDer = expDer
	return o
}

func (ntMult *NT_ExpMult) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntMult.ExpIzq.Interpretar(ctx)
	expDer := ntMult.ExpDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		if expIzq.Tipo > expDer.Tipo {
			expDer = ctx.Conversor.Ampliar(expDer, expIzq.Tipo)
		} else {
			expIzq = ctx.Conversor.Ampliar(expIzq, expDer.Tipo)
		}
	}

	if expIzq.Nil || expDer.Nil {
		ctx.AddError("No se puede realizar una Multiplicacion de tipos incompatibles")
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("No se pueden multiplicar expresiones de tipo Nil")
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("No se pueden multiplicar expresiones de tipo Bool")
		return interprete.NewNil()
	case interprete.Integer:
		return interprete.NewIntLiteral(expIzq.Valor * expDer.Valor)
	case interprete.String:
		ctx.AddError("No se pueden multiplicar expresiones de tipo String")
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("No se pueden multiplicar expresiones de tipo Character")
		return interprete.NewNil()
	case interprete.Float:
		return interprete.NewFloatLiteral(expIzq.ValorF * expDer.ValorF)
	default:
		ctx.AddError("Error de tipos no valido" + expIzq.Tipo.String() +
			" y " + expDer.Tipo.String())
	}

	return interprete.NewNil()
}

type NT_ExpDiv struct {
	ExpIzq interprete.AbstrExpr
	ExpDer interprete.AbstrExpr
}

// Constructor for NT_ExpDiv
func NewNT_ExpDiv(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr) *NT_ExpDiv {
	o := new(NT_ExpDiv)
	o.ExpIzq = expIzq
	o.ExpDer = expDer
	return o
}

func (ntDiv *NT_ExpDiv) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntDiv.ExpIzq.Interpretar(ctx)
	expDer := ntDiv.ExpDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		if expIzq.Tipo > expDer.Tipo {
			expDer = ctx.Conversor.Ampliar(expDer, expIzq.Tipo)
		} else {
			expIzq = ctx.Conversor.Ampliar(expIzq, expDer.Tipo)
		}
	}

	if expIzq.Nil || expDer.Nil {
		ctx.AddError("No se puede realizar una Division de tipos incompatibles")
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("No se pueden dividir expresiones de tipo Nil")
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("No se pueden dividir expresiones de tipo Bool")
		return interprete.NewNil()
	case interprete.Integer:
		if expDer.Valor == 0 {
			ctx.AddError("No se puede realizar divisiones por 0")
			return interprete.NewNil()
		}
		return interprete.NewIntLiteral(expIzq.Valor / expDer.Valor)
	case interprete.String:
		ctx.AddError("No se pueden dividir expresiones de tipo String")
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("No se pueden dividir expresiones de tipo Character")
		return interprete.NewNil()
	case interprete.Float:
		if expDer.ValorF == 0 {
			ctx.AddError("No se puede realizar divisiones por 0")
			return interprete.NewNil()
		}
		return interprete.NewFloatLiteral(expIzq.ValorF / expDer.ValorF)
	default:
		ctx.AddError("Error de tipos no valido" + expIzq.Tipo.String() +
			" y " + expDer.Tipo.String())
	}

	return interprete.NewNil()
}
