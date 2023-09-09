package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_FunVectors struct {
	nombre   string
	tipoFunc string
	expre    interprete.AbstrExpr
}

// Constructor for NT_FunVectors
func NewNT_FunVectors(nombre string, tipoFunc string, expre interprete.AbstrExpr) *NT_FunVectors {
	o := new(NT_FunVectors)
	o.nombre = nombre
	o.tipoFunc = tipoFunc
	o.expre = expre
	return o
}

func NewNT_FunVectorsN() *NT_FunVectors {
	o := new(NT_FunVectors)
	o.nombre = ""
	o.tipoFunc = ""
	o.expre = nil
	return o
}

// Setter method for the field expre of type interprete.AbstrExpr in the object NT_FunVectors
func (n *NT_FunVectors) SetExpre(expre interprete.AbstrExpr) {
	n.expre = expre
}

// Setter method for the field nombre of type string in the object NT_FunVectors
func (n *NT_FunVectors) SetNombre(nombre string) {
	n.nombre = nombre
}

// Setter method for the field tipoFunc of type string in the object NT_FunVectors
func (n *NT_FunVectors) SetTipoFunc(tipoFunc string) {
	n.tipoFunc = tipoFunc
}

func (ntFuncVectors *NT_FunVectors) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {

	//primero obtenemos el valor del vector

	stat, getVectorTipo, contenedorTmp, res := ctx.GetValorVector(ntFuncVectors.nombre)
	if !res {
		ctx.AddError("Se encontraron errores al obtener el valor del arreglo")
		return interprete.NewNil()
	}
	if stat {
		ctx.AddError("El vector " + ntFuncVectors.nombre + " es constante")
		return interprete.NewNil()
	}
	if ntFuncVectors.tipoFunc == "append" {
		valor := ntFuncVectors.expre.Interpretar(ctx)
		if getVectorTipo != valor.Tipo {
			ctx.AddError("Error Append: Asignacion de tipos incompatibles")
			return interprete.NewNil()
		}
		contenedorTmp = append(contenedorTmp, *valor)
		ctx.ReasignarValorVector(ntFuncVectors.nombre, contenedorTmp)
		return interprete.NewNil()
	} else if ntFuncVectors.tipoFunc == "removeLast" {
		if len(contenedorTmp) <= 0 {
			ctx.AddError("Error Remove: No existen elementos en el vector")
			return interprete.NewNil()
		}
		contenedorTmp = contenedorTmp[:len(contenedorTmp)-1]
		ctx.ReasignarValorVector(ntFuncVectors.nombre, contenedorTmp)
		return interprete.NewNil()
	} else if ntFuncVectors.tipoFunc == "removeAt" {
		valor := ntFuncVectors.expre.Interpretar(ctx)
		if valor.Tipo != interprete.Integer {
			ctx.AddError("Error RemoveAt: Indice de tipo Incorrecto")
			return interprete.NewNil()
		}
		if len(contenedorTmp) <= valor.Valor {
			ctx.AddError("Error Remove:Indice fuera de rango")
			return interprete.NewNil()
		}
		contenedorTmp = append(contenedorTmp[:valor.Valor], contenedorTmp[valor.Valor+1:]...)
		ctx.ReasignarValorVector(ntFuncVectors.nombre, contenedorTmp)
		return interprete.NewNil()
	} else if ntFuncVectors.tipoFunc == "empty" {
		if len(contenedorTmp) != 0 {
			return interprete.NewBoolLiteral(false)
		} else {
			return interprete.NewBoolLiteral(true)
		}
	} else if ntFuncVectors.tipoFunc == "count" {
		tmp := len(contenedorTmp)
		return interprete.NewIntLiteral(tmp)
	}
	return interprete.NewNil()
}
