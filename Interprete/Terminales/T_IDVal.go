package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type T_IdExp struct {
	nombre string
	line   int
	column int
}

func (tID *T_IdExp) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//cuando venga solo el id tengo que ir a verificar a la memoria el nombre de esta y su valor
	_, tmp, resStatus := ctx.GetValorNativo(tID.nombre)
	if !resStatus {
		ctx.AddError("Variable: " + tID.nombre + " No declarada")
	}
	return tmp
}

// Constructor for T_IdExp
func NewT_IdExp(nombre string, line int, column int) *T_IdExp {
	o := new(T_IdExp)
	o.nombre = nombre
	o.line = line
	o.column = column
	return o
}

type T_IdVector struct {
	nombre    string
	posVector interprete.AbstrExpr
}

// Constructor for T_IdVector
func NewT_IdVector(nombre string, posVector interprete.AbstrExpr) *T_IdVector {
	o := new(T_IdVector)
	o.nombre = nombre
	o.posVector = posVector
	return o
}

func (tIdVector *T_IdVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//ejecutamos la posicion
	posTmp := tIdVector.posVector.Interpretar(ctx)
	if posTmp.Tipo != interprete.Integer {
		ctx.AddError("Error Vector, indice de tipo Incorrecto")
		return interprete.NewNil()
	}
	_, _, contenedorTmp, res := ctx.GetValorVector(tIdVector.nombre)
	if !res {
		ctx.AddError("Error Vector, No se pudo obtener el valor")
		return interprete.NewNil()
	}
	if len(contenedorTmp) <= posTmp.Valor {
		ctx.AddError("Error Vector, Indice Fuera de rango")
		return interprete.NewNil()
	}
	tmp := contenedorTmp[posTmp.Valor]
	if tmp.Tipo != interprete.Nil {
		return &tmp
	}
	return interprete.NewNil()
}
