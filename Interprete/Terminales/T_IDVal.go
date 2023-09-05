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
	_, tmp, resStatus := ctx.Memoria.GetSimbolo(tID.nombre)
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
