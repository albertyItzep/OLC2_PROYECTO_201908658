package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type NT_ExpSum struct {
	ExpIzquierda interprete.AbstrExpr
	ExpDerecha   interprete.AbstrExpr
	linea        int
	columna      int
}

// Constructor for NT_ExpSum
func NewNT_ExpSum(expIzquierda interprete.AbstrExpr, expDerecha interprete.AbstrExpr, linea int, columna int) *NT_ExpSum {
	o := new(NT_ExpSum)
	o.ExpIzquierda = expIzquierda
	o.ExpDerecha = expDerecha
	o.linea = linea
	o.columna = columna
	return o
}

func (ntSuma *NT_ExpSum) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntSuma.ExpIzquierda.Interpretar(ctx)
	expDer := ntSuma.ExpDerecha.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		if expIzq.Tipo > expDer.Tipo {
			expDer = ctx.Conversor.Ampliar(expDer, expIzq.Tipo, ntSuma.linea, ntSuma.columna)
		} else {
			expIzq = ctx.Conversor.Ampliar(expIzq, expDer.Tipo, ntSuma.linea, ntSuma.columna)
		}
	}

	if expIzq.Nil || expDer.Nil {
		ctx.AddError("No se puede realizar una suma de tipos incompatibles", ntSuma.linea, ntSuma.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("No se pueden sumar expresiones de tipo Nil", ntSuma.linea, ntSuma.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("No se pueden sumar expresiones de tipo Bool", ntSuma.linea, ntSuma.columna)
		return interprete.NewNil()
	case interprete.Integer:
		return interprete.NewIntLiteral(expIzq.Valor + expDer.Valor)
	case interprete.String:
		res := expIzq.ValorS + expDer.ValorS
		return interprete.NewStringLiteral(res)
	case interprete.Character:
		ctx.AddError("No se pueden sumar expresiones de tipo Character", ntSuma.linea, ntSuma.columna)
		return interprete.NewNil()
	case interprete.Float:
		return interprete.NewFloatLiteral(expIzq.ValorF + expDer.ValorF)
	default:
		ctx.AddError("Error de tipos no valido"+expIzq.Tipo.String()+
			" y "+expDer.Tipo.String(), ntSuma.linea, ntSuma.columna)
	}

	return interprete.NewNil()
}

type NT_ExpRes struct {
	ExpIzq  interprete.AbstrExpr
	ExpDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_ExpRes
func NewNT_ExpRes(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_ExpRes {
	o := new(NT_ExpRes)
	o.ExpIzq = expIzq
	o.ExpDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntRes *NT_ExpRes) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntRes.ExpIzq.Interpretar(ctx)
	expDer := ntRes.ExpDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		if expIzq.Tipo > expDer.Tipo {
			expDer = ctx.Conversor.Ampliar(expDer, expIzq.Tipo, ntRes.linea, ntRes.columna)
		} else {
			expIzq = ctx.Conversor.Ampliar(expIzq, expDer.Tipo, ntRes.linea, ntRes.columna)
		}
	}

	if expIzq.Nil || expDer.Nil {
		ctx.AddError("No se puede realizar una resta de tipos incompatibles", ntRes.linea, ntRes.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("No se pueden restar expresiones de tipo Nil", ntRes.linea, ntRes.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("No se pueden restar expresiones de tipo Bool", ntRes.linea, ntRes.columna)
		return interprete.NewNil()
	case interprete.Integer:
		return interprete.NewIntLiteral(expIzq.Valor - expDer.Valor)
	case interprete.String:
		ctx.AddError("No se pueden sumar expresiones de tipo String", ntRes.linea, ntRes.columna)
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("No se pueden restar expresiones de tipo Character", ntRes.linea, ntRes.columna)
		return interprete.NewNil()
	case interprete.Float:
		return interprete.NewFloatLiteral(expIzq.ValorF - expDer.ValorF)
	default:
		ctx.AddError("Error de tipos no valido"+expIzq.Tipo.String()+
			" y "+expDer.Tipo.String(), ntRes.linea, ntRes.columna)
	}

	return interprete.NewNil()
}

type NT_ExpModulo struct {
	ExpIzq  interprete.AbstrExpr
	ExpDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_ExpModulo
func NewNT_ExpModulo(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_ExpModulo {
	o := new(NT_ExpModulo)
	o.ExpIzq = expIzq
	o.ExpDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntModulo *NT_ExpModulo) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntModulo.ExpIzq.Interpretar(ctx)
	expDer := ntModulo.ExpDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		if expIzq.Tipo > expDer.Tipo {
			expDer = ctx.Conversor.Ampliar(expDer, expIzq.Tipo, ntModulo.linea, ntModulo.columna)
		} else {
			expIzq = ctx.Conversor.Ampliar(expIzq, expDer.Tipo, ntModulo.linea, ntModulo.columna)
		}
	}

	if expIzq.Nil || expDer.Nil {
		ctx.AddError("No se puede realizar un Modulo de tipos incompatibles", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("No se puede realizar un Modulo de expresiones de tipo Nil", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("No se puede realizar un Modulo de tipo Bool", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	case interprete.Integer:
		if expDer.Valor == 0 {
			ctx.AddError("No se puede realizar divisiones por 0", ntModulo.linea, ntModulo.columna)
			return interprete.NewNil()
		}
		return interprete.NewIntLiteral(expIzq.Valor % expDer.Valor)
	case interprete.String:
		ctx.AddError("No se puede realizar un Modulo detipo String", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("No se puede realizar un Modulo detipo Character", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	case interprete.Float:
		ctx.AddError("No se puede realizar un Modulo detipo Float", ntModulo.linea, ntModulo.columna)
		return interprete.NewNil()
	default:
		ctx.AddError("Error de tipos no valido"+expIzq.Tipo.String()+
			" y "+expDer.Tipo.String(), ntModulo.linea, ntModulo.columna)
	}

	return interprete.NewNil()
}
