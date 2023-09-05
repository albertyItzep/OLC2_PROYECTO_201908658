package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_AsIncremento struct {
	id  string
	exp interprete.AbstrExpr
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
		ctx.AddError("Error Asignacion: Variable Inexistente")
		return interprete.NewNil()
	}
	if constante {
		ctx.AddError("Error Asignacion: Variable " + ntIncremento.id + " es constante")
		return interprete.NewNil()
	}
	switch exp.Tipo {
	case interprete.Nil:
		ctx.AddError("Error Asignacion: No puede incrementarse nil")
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error Asignacion: No puede incrementarse bool ")
		return interprete.NewNil()
	case interprete.Integer:
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntIncremento.id, tmp)
			return interprete.NewNil()
		}
		tmp.Valor = tmp.Valor + exp.Valor
		ctx.ReasignarValorNativo(ntIncremento.id, tmp)
		return interprete.NewNil()
	case interprete.String:
		ctx.AddError("Error Asignacion: No puede incrementarse string ")
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("Error Asignacion: No puede incrementarse character ")
		return interprete.NewNil()
	case interprete.Float:
		ctx.Conversor.Ampliar(exp, interprete.Float)
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntIncremento.id, tmp)
			return interprete.NewNil()
		}
		tmp.ValorF = tmp.ValorF + exp.ValorF
		ctx.ReasignarValorNativo(ntIncremento.id, tmp)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_AsDecremento struct {
	id  string
	exp interprete.AbstrExpr
}

// Constructor for NT_AsDecremento
func NewNT_AsDecremento(id string, exp interprete.AbstrExpr) *NT_AsDecremento {
	o := new(NT_AsDecremento)
	o.id = id
	o.exp = exp
	return o
}

func (ntDecremento *NT_AsDecremento) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntDecremento.exp.Interpretar(ctx)

	constante, tmp, resStatus := ctx.GetValorNativo(ntDecremento.id)
	if !resStatus {
		ctx.AddError("Error Asignacion: Variable Inexistente")
		return interprete.NewNil()
	}
	if constante {
		ctx.AddError("Error Asignacion: Variable " + ntDecremento.id + " es constante")
		return interprete.NewNil()
	}
	switch exp.Tipo {
	case interprete.Nil:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento nil")
		return interprete.NewNil()
	case interprete.Bool:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento de bool ")
		return interprete.NewNil()
	case interprete.Integer:
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntDecremento.id, tmp)
			return interprete.NewNil()
		}
		tmp.Valor = tmp.Valor - exp.Valor
		ctx.ReasignarValorNativo(ntDecremento.id, tmp)
		return interprete.NewNil()
	case interprete.String:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento de string ")
		return interprete.NewNil()
	case interprete.Character:
		ctx.AddError("Error Asignacion: No se puede realizar un decremento character ")
		return interprete.NewNil()
	case interprete.Float:
		ctx.Conversor.Ampliar(exp, interprete.Float)
		if tmp.Tipo == interprete.Nil {
			tmp = exp
			ctx.ReasignarValorNativo(ntDecremento.id, tmp)
			return interprete.NewNil()
		}
		tmp.ValorF = tmp.ValorF - exp.ValorF
		ctx.ReasignarValorNativo(ntDecremento.id, tmp)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}
