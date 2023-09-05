package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
)

// variables
type NT_Variables struct {
	Id        string
	Tipo      interprete.TipoD
	expresion interprete.AbstrExpr
	linea     int
	columna   int
}

func NewNT_Variables(Id string, tipo interprete.TipoD, expresion interprete.AbstrExpr, linea int, columna int) *NT_Variables {
	return &NT_Variables{
		Id:        Id,
		Tipo:      tipo,
		expresion: expresion,
		linea:     linea,
		columna:   columna,
	}
}

func (ntVar *NT_Variables) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if ntVar.expresion != nil {
		expr = ntVar.expresion.Interpretar(ctx)
		res, resString := ctx.AgregarVariable(ntVar.Id, ntVar.Tipo, interprete.Nativo, -1, nil, expr, nil, ntVar.linea, ntVar.columna)
		if !res {
			ctx.AddError(resString)
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}

type NT_VariablesTipo struct {
	Id      string
	Tipo    interprete.TipoD
	linea   int
	columna int
}

// Constructor for NT_VariablesTipo
func NewNT_VariablesTipo(id string, tipo interprete.TipoD, linea int, columna int) *NT_VariablesTipo {
	o := new(NT_VariablesTipo)
	o.Id = id
	o.Tipo = tipo
	o.linea = linea
	o.columna = columna
	return o
}

func (ntVarTipo NT_VariablesTipo) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	res, resString := ctx.AgregarVariableTipo(ntVarTipo.Id, ntVarTipo.Tipo, interprete.Nativo, -1, nil, expr, nil, ntVarTipo.linea, ntVarTipo.columna)
	if !res {
		ctx.AddError(resString)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_VariableExpresion struct {
	Id        string
	Expresion interprete.AbstrExpr
	linea     int
	columna   int
}

// Constructor for NT_VariableExpresion
func NewNT_VariableExpresion(id string, expresion interprete.AbstrExpr, linea int, columna int) *NT_VariableExpresion {
	o := new(NT_VariableExpresion)
	o.Id = id
	o.Expresion = expresion
	o.linea = linea
	o.columna = columna
	return o
}

func (ntVarExpresion *NT_VariableExpresion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if ntVarExpresion.Expresion != nil {
		expr = ntVarExpresion.Expresion.Interpretar(ctx)
		fmt.Println(expr.Tipo)
		res, resString := ctx.AgregarVariable(ntVarExpresion.Id, expr.Tipo, interprete.Nativo, -1, nil, expr, nil, ntVarExpresion.linea, ntVarExpresion.columna)
		if !res {
			ctx.AddError(resString)
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}
