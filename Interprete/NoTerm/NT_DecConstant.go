package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type NT_DecConstanteExpresion struct {
	Id        string
	Expresion interprete.AbstrExpr
	linea     int
	columna   int
}

// Constructor for NT_DecConstanteExpresion
func NewNT_DecConstanteExpresion(id string, expresion interprete.AbstrExpr, linea int, columna int) *NT_DecConstanteExpresion {
	o := new(NT_DecConstanteExpresion)
	o.Id = id
	o.Expresion = expresion
	o.linea = linea
	o.columna = columna
	return o
}

func (ntDecConstEx *NT_DecConstanteExpresion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if ntDecConstEx.Expresion != nil {
		expr = ntDecConstEx.Expresion.Interpretar(ctx)
		res, resString := ctx.AgregarConstanteExpresion(ntDecConstEx.Id, expr.Tipo, interprete.Nativo, -1, nil, expr, nil, ntDecConstEx.linea, ntDecConstEx.columna)
		if !res {
			ctx.AddError(resString, ntDecConstEx.linea, ntDecConstEx.columna)
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}

type NT_DecConsComp struct {
	Id        string
	Tipo      interprete.TipoD
	expresion interprete.AbstrExpr
	linea     int
	columna   int
}

// Constructor for NT_DecConsComp
func NewNT_DecConsComp(id string, tipo interprete.TipoD, expresion interprete.AbstrExpr, linea int, columna int) *NT_DecConsComp {
	o := new(NT_DecConsComp)
	o.Id = id
	o.Tipo = tipo
	o.expresion = expresion
	o.linea = linea
	o.columna = columna
	return o
}

func (ntConsComplet *NT_DecConsComp) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if ntConsComplet.expresion != nil {
		expr = ntConsComplet.expresion.Interpretar(ctx)
		res, resString := ctx.AgregarConstanteExpresion(ntConsComplet.Id, ntConsComplet.Tipo, interprete.Nativo, -1, nil, expr, nil, ntConsComplet.linea, ntConsComplet.columna)
		if !res {
			ctx.AddError(resString, ntConsComplet.linea, ntConsComplet.columna)
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}
