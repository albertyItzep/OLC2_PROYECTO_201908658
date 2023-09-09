package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"strconv"
)

type NT_DecVector struct {
	nombre      string
	tipoVar     interprete.TipoD
	tipoId      interprete.TipoD
	Expresiones []interprete.AbstrExpr
	linea       int
	columna     int
}

// Constructor for NT_DecVector
func NewNT_DecVector(nombre string, tipoVar interprete.TipoD, tipoId interprete.TipoD, linea int, columna int) *NT_DecVector {
	o := new(NT_DecVector)
	o.nombre = nombre
	o.tipoVar = tipoVar
	o.tipoId = tipoId
	o.linea = linea
	o.columna = columna
	return o
}

func (ntVector *NT_DecVector) AddExpresion(exp interprete.AbstrExpr) {
	ntVector.Expresiones = append(ntVector.Expresiones, exp)
}

func (ntVector *NT_DecVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {

	//primero sabemos que nombre tiene eso siempre viene
	//debemos saber que tipo es en este caso el valor de tipo que contendra si seran int o algun objeto en especifico
	if ntVector.tipoVar != interprete.Nil && ntVector.tipoId == interprete.Nil {
		//vector de tipos nativos
		tmp := make([]interprete.Resultado, 0)
		for _, v := range ntVector.Expresiones {
			tmp1 := *v.Interpretar(ctx)
			if tmp1.Tipo == ntVector.tipoVar {
				tmp = append(tmp, tmp1)
			} else {
				ctx.AddError("Error Vector: Tipos no compatibles Linea: " + strconv.Itoa(ntVector.linea) + " Columna: " + strconv.Itoa(ntVector.columna))
				return interprete.NewNil()
			}
		}
		res, strRes := ctx.AgregarVector(false, ntVector.nombre, interprete.Vector, ntVector.tipoVar, interprete.Nativo, -1, nil, nil, tmp, ntVector.linea, ntVector.columna)
		if !res {
			ctx.AddError(strRes)
			return interprete.NewNil()
		}
		return interprete.NewNil()
	} else if ntVector.tipoVar == interprete.Nil && ntVector.tipoId != interprete.Nil {
		//verctor de tipo struct

	}
	panic("not implemented") // TODO: Implement
}
