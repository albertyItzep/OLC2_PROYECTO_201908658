package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_AsigVector struct {
	nombre         string
	tipoAsignacion string
	posVector      interprete.AbstrExpr
	valor          interprete.AbstrExpr
}

// Constructor for NT_AsigVector
func NewNT_AsigVector(nombre string, tipoAsignacion string, posVector interprete.AbstrExpr, valor interprete.AbstrExpr) *NT_AsigVector {
	o := new(NT_AsigVector)
	o.nombre = nombre
	o.tipoAsignacion = tipoAsignacion
	o.posVector = posVector
	o.valor = valor
	return o
}

func (ntAsigVector *NT_AsigVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//validamos tipo de asignacion
	if ntAsigVector.tipoAsignacion == "+=" {
		//validamos la posicion del vector
		expTmp := ntAsigVector.posVector.Interpretar(ctx)
		if expTmp.Tipo != interprete.Integer {
			ctx.AddError("Error Vector: Posicion de tipo invalida")
			return interprete.NewNil()
		}
		valorTmp := ntAsigVector.valor.Interpretar(ctx)
		stat, getVectorTipo, contenedorTmp, res := ctx.GetValorVector(ntAsigVector.nombre)
		if !res {
			ctx.AddError("Se encontraron errores al obtener el valor del arreglo")
			return interprete.NewNil()
		}
		if stat {
			ctx.AddError("El vector " + ntAsigVector.nombre + " es constante")
			return interprete.NewNil()
		}
		if len(contenedorTmp) <= expTmp.Valor {
			ctx.AddError("Indice fuera de rango")
			return interprete.NewNil()
		}
		if valorTmp.Tipo == interprete.Integer || valorTmp.Tipo == interprete.Float {
			switch getVectorTipo {
			case interprete.Integer:
				valorTmp = ctx.Conversor.Ampliar(valorTmp, getVectorTipo)
				contenedorTmp[expTmp.Valor].Valor = contenedorTmp[expTmp.Valor].Valor + valorTmp.Valor
				ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
				return interprete.NewNil()
			case interprete.Float:
				valorTmp = ctx.Conversor.Ampliar(valorTmp, getVectorTipo)
				contenedorTmp[expTmp.Valor].ValorF = contenedorTmp[expTmp.Valor].ValorF + valorTmp.ValorF
				ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
				return interprete.NewNil()
			default:
				ctx.AddError("Error incremento de tipos incompatibles")
			}
		} else {
			ctx.AddError("Incremeto de tipo incorrecto para el vector")
			return interprete.NewNil()
		}
	} else if ntAsigVector.tipoAsignacion == "-=" {
		expTmp := ntAsigVector.posVector.Interpretar(ctx)
		if expTmp.Tipo != interprete.Integer {
			ctx.AddError("Error Vector: Posicion de tipo invalida")
			return interprete.NewNil()
		}
		valorTmp := ntAsigVector.valor.Interpretar(ctx)
		stat, getVectorTipo, contenedorTmp, res := ctx.GetValorVector(ntAsigVector.nombre)
		if !res {
			ctx.AddError("Se encontraron errores al obtener el valor del arreglo")
			return interprete.NewNil()
		}
		if stat {
			ctx.AddError("El vector " + ntAsigVector.nombre + " es constante")
			return interprete.NewNil()
		}
		if len(contenedorTmp) <= expTmp.Valor {
			ctx.AddError("Indice fuera de rango")
			return interprete.NewNil()
		}
		if valorTmp.Tipo == interprete.Integer || valorTmp.Tipo == interprete.Float {
			switch getVectorTipo {
			case interprete.Integer:
				valorTmp = ctx.Conversor.Ampliar(valorTmp, getVectorTipo)
				contenedorTmp[expTmp.Valor].Valor = contenedorTmp[expTmp.Valor].Valor - valorTmp.Valor
				ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
				return interprete.NewNil()
			case interprete.Float:
				valorTmp = ctx.Conversor.Ampliar(valorTmp, getVectorTipo)
				contenedorTmp[expTmp.Valor].ValorF = contenedorTmp[expTmp.Valor].ValorF - valorTmp.ValorF
				ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
				return interprete.NewNil()
			default:
				ctx.AddError("Error Decremento de tipos incompatibles")
				return interprete.NewNil()
			}
		} else {
			ctx.AddError("Incremeto de tipo incorrecto para el vector")
			return interprete.NewNil()
		}
	} else if ntAsigVector.tipoAsignacion == "==" {
		expTmp := ntAsigVector.posVector.Interpretar(ctx)
		if expTmp.Tipo != interprete.Integer {
			ctx.AddError("Error Vector: Posicion de tipo invalida")
			return interprete.NewNil()
		}
		valorTmp := ntAsigVector.valor.Interpretar(ctx)
		stat, getVectorTipo, contenedorTmp, res := ctx.GetValorVector(ntAsigVector.nombre)
		if !res {
			ctx.AddError("Se encontraron errores al obtener el valor del arreglo")
			return interprete.NewNil()
		}
		if stat {
			ctx.AddError("El vector " + ntAsigVector.nombre + " es constante")
			return interprete.NewNil()
		}
		if len(contenedorTmp) <= expTmp.Valor {
			ctx.AddError("Indice fuera de rango")
			return interprete.NewNil()
		}
		switch getVectorTipo {
		case interprete.Bool:
			if valorTmp.Tipo != interprete.Bool {
				ctx.AddError("Error Asignacion de tipos incompatibles")
				return interprete.NewNil()
			}
			contenedorTmp[expTmp.Valor].ValorB = valorTmp.ValorB
			ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
			return interprete.NewNil()
		case interprete.Integer:
			if valorTmp.Tipo != interprete.Integer {
				ctx.AddError("Error Asignacion de tipos incompatibles")
				return interprete.NewNil()
			}
			contenedorTmp[expTmp.Valor].Valor = valorTmp.Valor
			ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
			return interprete.NewNil()
		case interprete.String:
			if valorTmp.Tipo != interprete.String {
				ctx.AddError("Error Asignacion de tipos incompatibles")
				return interprete.NewNil()
			}
			contenedorTmp[expTmp.Valor].ValorS = valorTmp.ValorS
			ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
			return interprete.NewNil()
		case interprete.Character:
			if valorTmp.Tipo != interprete.Character {
				ctx.AddError("Error Asignacion de tipos incompatibles")
				return interprete.NewNil()
			}
			contenedorTmp[expTmp.Valor].ValorC = valorTmp.ValorC
			ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
			return interprete.NewNil()
		case interprete.Float:
			if valorTmp.Tipo != interprete.Float {
				ctx.AddError("Error Asignacion de tipos incompatibles")
				return interprete.NewNil()
			}
			contenedorTmp[expTmp.Valor].ValorF = valorTmp.ValorF
			ctx.ReasignarValorVector(ntAsigVector.nombre, contenedorTmp)
			return interprete.NewNil()
		default:
			ctx.AddError("Error Asignacion de tipos incompatibles")
			return interprete.NewNil()
		}
	}
	return interprete.NewNil()
}
