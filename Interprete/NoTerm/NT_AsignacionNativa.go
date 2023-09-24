package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type NT_AsIncremento struct {
	id      string
	exp     interprete.AbstrExpr
	Linea   int
	Columna int
}

// Constructor for NT_AsIncremento
func NewNT_AsIncremento(id string, exp interprete.AbstrExpr) *NT_AsIncremento {
	o := new(NT_AsIncremento)
	o.id = id
	o.exp = exp
	return o
}

func (ntIncremento *NT_AsIncremento) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntIncremento.exp.Interpretar(ctx)

	constante, tmp, resStatus := ctx.GetValorNativo(ntIncremento.id)
	if !resStatus {
		ctx.AddError("Error Asignacion: Variable Inexistente", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	}
	if constante {
		ctx.AddError("Error Asignacion: Variable "+ntIncremento.id+" es constante", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	}
	switch exp.Tipo {
	case interprete.Nil:
		ctx.AddError("Error Asignacion: No puede incrementarse nil", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error Asignacion: No puede incrementarse bool ", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	case interprete.Integer:
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntIncremento.id, tmp, ntIncremento.Linea, ntIncremento.Columna)
			return interprete.NewNil()
		}
		tmp.Valor = tmp.Valor + exp.Valor
		ctx.ReasignarValorNativo(ntIncremento.id, tmp, ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	case interprete.String:
		ctx.AddError("Error Asignacion: No puede incrementarse string ", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("Error Asignacion: No puede incrementarse character ", ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	case interprete.Float:
		ctx.Conversor.Ampliar(exp, interprete.Float, ntIncremento.Linea, ntIncremento.Columna)
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntIncremento.id, tmp, ntIncremento.Linea, ntIncremento.Columna)
			return interprete.NewNil()
		}
		tmp.ValorF = tmp.ValorF + exp.ValorF
		ctx.ReasignarValorNativo(ntIncremento.id, tmp, ntIncremento.Linea, ntIncremento.Columna)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_AsDecremento struct {
	id      string
	exp     interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_AsDecremento
func NewNT_AsDecremento(id string, exp interprete.AbstrExpr, linea int, columna int) *NT_AsDecremento {
	o := new(NT_AsDecremento)
	o.id = id
	o.exp = exp
	o.linea = linea
	o.columna = columna
	return o
}

func (ntDecremento *NT_AsDecremento) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntDecremento.exp.Interpretar(ctx)

	constante, tmp, resStatus := ctx.GetValorNativo(ntDecremento.id)
	if !resStatus {
		ctx.AddError("Error Asignacion: Variable Inexistente", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	}
	if constante {
		ctx.AddError("Error Asignacion: Variable "+ntDecremento.id+" es constante", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	}
	switch exp.Tipo {
	case interprete.Nil:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento nil", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento de bool ", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	case interprete.Integer:
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntDecremento.id, tmp, ntDecremento.linea, ntDecremento.columna)
			return interprete.NewNil()
		}
		tmp.Valor = tmp.Valor - exp.Valor
		ctx.ReasignarValorNativo(ntDecremento.id, tmp, ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	case interprete.String:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento de string ", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento character ", ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	case interprete.Float:
		ctx.Conversor.Ampliar(exp, interprete.Float, ntDecremento.linea, ntDecremento.columna)
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntDecremento.id, tmp, ntDecremento.linea, ntDecremento.columna)
			return interprete.NewNil()
		}
		tmp.ValorF = tmp.ValorF - exp.ValorF
		ctx.ReasignarValorNativo(ntDecremento.id, tmp, ntDecremento.linea, ntDecremento.columna)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_AsGeneral struct {
	id      string
	exp     interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_AsGeneral
func NewNT_AsGeneral(id string, exp interprete.AbstrExpr, linea int, columna int) *NT_AsGeneral {
	o := new(NT_AsGeneral)
	o.id = id
	o.exp = exp
	o.linea = linea
	o.columna = columna
	return o
}

func (ntGeneral *NT_AsGeneral) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := ntGeneral.exp.Interpretar(ctx)

	if expr.Tipo == interprete.Nil {
		ctx.AddError("Error Asignacion: No se puede asignar un nil", ntGeneral.linea, ntGeneral.columna)
	}
	ctx.ReasignarValorNativo(ntGeneral.id, expr, ntGeneral.linea, ntGeneral.columna)
	return interprete.NewNil()
}
