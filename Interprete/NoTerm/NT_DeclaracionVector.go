package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
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
	o.Expresiones = make([]interprete.AbstrExpr, 0)
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
				ctx.AddError("Error Vector: Tipos no compatibles Linea: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
				return interprete.NewNil()
			}
		}
		res, strRes := ctx.AgregarVector(false, ntVector.nombre, interprete.Vector, ntVector.tipoVar, interprete.Nativo, -1, nil, nil, tmp, ntVector.linea, ntVector.columna)
		if !res {
			ctx.AddError(strRes, ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		return interprete.NewNil()
	} else if ntVector.tipoVar == interprete.Nil && ntVector.tipoId == interprete.Nil {
		if len(ntVector.Expresiones) == 0 {
			ctx.AddError("Error Vector: Vector sin tipo, vacio en: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		tmp := make([]interprete.Resultado, 0)
		for _, v := range ntVector.Expresiones {
			tmp1 := *v.Interpretar(ctx)
			tmp = append(tmp, tmp1)
		}
		tmpTipo := tmp[0].Tipo
		for _, v := range tmp {
			if v.Tipo != tmpTipo {
				ctx.AddError("Error Vector: Tipos no compatibles Linea: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
				return interprete.NewNil()
			}
		}
		res, strRes := ctx.AgregarVector(false, ntVector.nombre, interprete.Vector, tmp[0].Tipo, interprete.Nativo, -1, nil, nil, tmp, ntVector.linea, ntVector.columna)
		if !res {
			ctx.AddError(strRes, ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_DecVectorConst struct {
	nombre      string
	tipoVar     interprete.TipoD
	tipoId      interprete.TipoD
	Expresiones []interprete.AbstrExpr
	linea       int
	columna     int
}

// Constructor for NT_DecVectorConst
func NewNT_DecVectorConst(nombre string, tipoVar interprete.TipoD, tipoId interprete.TipoD, linea int, columna int) *NT_DecVectorConst {
	o := new(NT_DecVectorConst)
	o.nombre = nombre
	o.tipoVar = tipoVar
	o.tipoId = tipoId
	o.linea = linea
	o.Expresiones = make([]interprete.AbstrExpr, 0)
	o.columna = columna
	return o
}
func (ntVector *NT_DecVectorConst) AddExpresion(exp interprete.AbstrExpr) {
	ntVector.Expresiones = append(ntVector.Expresiones, exp)
}

func (ntVector *NT_DecVectorConst) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {

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
				ctx.AddError("Error Vector: Tipos no compatibles Linea: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
				return interprete.NewNil()
			}
		}
		res, strRes := ctx.AgregarVector(true, ntVector.nombre, interprete.Vector, ntVector.tipoVar, interprete.Nativo, -1, nil, nil, tmp, ntVector.linea, ntVector.columna)
		if !res {
			ctx.AddError(strRes, ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		return interprete.NewNil()
	} else if ntVector.tipoVar == interprete.Nil && ntVector.tipoId != interprete.Nil {
		//verctor de tipo struct

	} else if ntVector.tipoVar == interprete.Nil && ntVector.tipoId == interprete.Nil {
		if len(ntVector.Expresiones) == 0 {
			ctx.AddError("Error Vector: Vector sin tipo, vacio en: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		tmp := make([]interprete.Resultado, 0)
		for _, v := range ntVector.Expresiones {
			tmp1 := *v.Interpretar(ctx)
			tmp = append(tmp, tmp1)
		}
		tmpTipo := tmp[0].Tipo
		for _, v := range tmp {
			if v.Tipo != tmpTipo {
				ctx.AddError("Error Vector: Tipos no compatibles Linea: "+strconv.Itoa(ntVector.linea)+" Columna: "+strconv.Itoa(ntVector.columna), ntVector.linea, ntVector.columna)
				return interprete.NewNil()
			}
		}
		res, strRes := ctx.AgregarVector(true, ntVector.nombre, interprete.Vector, tmp[0].Tipo, interprete.Nativo, -1, nil, nil, tmp, ntVector.linea, ntVector.columna)
		if !res {
			ctx.AddError(strRes, ntVector.linea, ntVector.columna)
			return interprete.NewNil()
		}
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

type NT_DecVectorID struct {
	nombre   string
	varConst bool
	tipoVar  interprete.TipoD
	tipoId   string
	idExp    interprete.AbstrExpr
	linea    int
	columna  int
}

// Constructor for NT_DecVectorID
func NewNT_DecVectorID(nombre string, varConst bool, tipoVar interprete.TipoD, tipoId string, idExp interprete.AbstrExpr, linea int, columna int) *NT_DecVectorID {
	o := new(NT_DecVectorID)
	o.nombre = nombre
	o.varConst = varConst
	o.tipoVar = tipoVar
	o.tipoId = tipoId
	o.idExp = idExp
	o.linea = linea
	o.columna = columna
	return o
}

func (ntVecID *NT_DecVectorID) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	if ntVecID.tipoVar != interprete.Nil && ntVecID.tipoId == "" {
		expvar := ntVecID.idExp.Interpretar(ctx)
		fmt.Println(expvar)
		if expvar.Tipo == interprete.Vector && ntVecID.tipoVar == expvar.TipoVector {
			if ntVecID.varConst { //variable
				res, strRes := ctx.AgregarVector(false, ntVecID.nombre, interprete.Vector, ntVecID.tipoVar, interprete.Nativo, -1, nil, nil, expvar.Vector, ntVecID.linea, ntVecID.columna)
				if !res {
					ctx.AddError(strRes, ntVecID.linea, ntVecID.columna)
					return interprete.NewNil()
				}
			} else { // constante
				res, strRes := ctx.AgregarVector(true, ntVecID.nombre, interprete.Vector, ntVecID.tipoVar, interprete.Nativo, -1, nil, nil, expvar.Vector, ntVecID.linea, ntVecID.columna)
				if !res {
					ctx.AddError(strRes, ntVecID.linea, ntVecID.columna)
					return interprete.NewNil()
				}
			}
			return interprete.NewNil()
		} else {
			ctx.AddError("tipos incompatibles vector", ntVecID.linea, ntVecID.columna)
			return interprete.NewNil()
		}
		//es nativa
	} else if ntVecID.tipoVar == interprete.Nil && ntVecID.tipoId != "" {
		// es struct
	} else if ntVecID.tipoVar == interprete.Nil && ntVecID.tipoId == "" {
		// es del tipo que posee el id
		expvar := ntVecID.idExp.Interpretar(ctx)

		if expvar.Tipo == interprete.Vector && expvar.TipoVector != interprete.Nil {
			if ntVecID.varConst {
				res, strRes := ctx.AgregarVector(false, ntVecID.nombre, interprete.Vector, ntVecID.tipoVar, interprete.Nativo, -1, nil, nil, expvar.Vector, ntVecID.linea, ntVecID.columna)
				if !res {
					ctx.AddError(strRes, ntVecID.linea, ntVecID.columna)
					return interprete.NewNil()
				}
			} else {
				res, strRes := ctx.AgregarVector(true, ntVecID.nombre, interprete.Vector, ntVecID.tipoVar, interprete.Nativo, -1, nil, nil, expvar.Vector, ntVecID.linea, ntVecID.columna)
				if !res {
					ctx.AddError(strRes, ntVecID.linea, ntVecID.columna)
					return interprete.NewNil()
				}
			}
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}
