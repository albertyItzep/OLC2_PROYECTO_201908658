package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_CompIgualdad struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_OpComparacion
func NewNT_CompIgualdad(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_CompIgualdad {
	o := new(NT_CompIgualdad)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntComparacion *NT_CompIgualdad) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntComparacion.expIzq.Interpretar(ctx)
	expDer := ntComparacion.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Comparacion de tipos incompatibles", ntComparacion.linea, ntComparacion.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar comparacion de Nil", ntComparacion.linea, ntComparacion.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar comparacion de Bool", ntComparacion.linea, ntComparacion.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor == expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS == expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC == expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF == expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}

type NT_CompDif struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_CompDif
func NewNT_CompDif(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_CompDif {
	o := new(NT_CompDif)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntCompDif *NT_CompDif) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntCompDif.expIzq.Interpretar(ctx)
	expDer := ntCompDif.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Comparacion de tipos incompatibles", ntCompDif.linea, ntCompDif.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar comparacion de Nil", ntCompDif.linea, ntCompDif.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar comparacion de Bool", ntCompDif.linea, ntCompDif.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor != expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS != expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC != expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF != expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}
