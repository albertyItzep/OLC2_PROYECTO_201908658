package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_RelMayor struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_RelMayor
func NewNT_RelMayor(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_RelMayor {
	o := new(NT_RelMayor)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntMayor *NT_RelMayor) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntMayor.expIzq.Interpretar(ctx)
	expDer := ntMayor.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Relacion de tipos incompatibles", ntMayor.linea, ntMayor.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar la relacion de Nil", ntMayor.linea, ntMayor.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar la relacion de Bool", ntMayor.linea, ntMayor.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor > expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS > expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC > expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF > expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}

type NT_RelMayIgual struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_RelMayIgual
func NewNT_RelMayIgual(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_RelMayIgual {
	o := new(NT_RelMayIgual)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntMayIgual *NT_RelMayIgual) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntMayIgual.expIzq.Interpretar(ctx)
	expDer := ntMayIgual.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Relacion de tipos incompatibles", ntMayIgual.linea, ntMayIgual.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar la relacion de Nil", ntMayIgual.linea, ntMayIgual.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar la relacion de Bool", ntMayIgual.linea, ntMayIgual.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor >= expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS >= expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC >= expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF >= expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}

type NT_RelMenor struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_RelMenor
func NewNT_RelMenor(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_RelMenor {
	o := new(NT_RelMenor)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntMenor *NT_RelMenor) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntMenor.expIzq.Interpretar(ctx)
	expDer := ntMenor.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Relacion de tipos incompatibles", ntMenor.linea, ntMenor.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar la relacion de Nil", ntMenor.linea, ntMenor.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar la relacion de Bool", ntMenor.linea, ntMenor.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor < expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS < expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC < expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF < expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}

type NT_RelMenIgual struct {
	expIzq  interprete.AbstrExpr
	expDer  interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_RelMenIgual
func NewNT_RelMenIgual(expIzq interprete.AbstrExpr, expDer interprete.AbstrExpr, linea int, columna int) *NT_RelMenIgual {
	o := new(NT_RelMenIgual)
	o.expIzq = expIzq
	o.expDer = expDer
	o.linea = linea
	o.columna = columna
	return o
}

func (ntMenIgual *NT_RelMenIgual) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expIzq := ntMenIgual.expIzq.Interpretar(ctx)
	expDer := ntMenIgual.expDer.Interpretar(ctx)

	if expIzq.Tipo != expDer.Tipo {
		ctx.AddError("Error: Relacion de tipos incompatibles", ntMenIgual.linea, ntMenIgual.columna)
		return interprete.NewNil()
	}

	switch expIzq.Tipo {
	case interprete.Nil:
		ctx.AddError("Error: No se puede realizar la relacion de Nil", ntMenIgual.linea, ntMenIgual.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar la relacion de Bool", ntMenIgual.linea, ntMenIgual.columna)
		return interprete.NewNil()
	case interprete.Integer:
		valRes := expIzq.Valor <= expDer.Valor
		return interprete.NewBoolLiteral(valRes)
	case interprete.String:
		valRes := expIzq.ValorS <= expDer.ValorS
		return interprete.NewBoolLiteral(valRes)
	case interprete.Character:
		valRes := expIzq.ValorC <= expDer.ValorC
		return interprete.NewBoolLiteral(valRes)
	case interprete.Float:
		valRes := expIzq.ValorF <= expDer.ValorF
		return interprete.NewBoolLiteral(valRes)
	}
	return interprete.NewNil()
}
