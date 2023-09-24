package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
)

type NT_InsFor struct {
	id      string
	exp     interprete.AbstrExpr
	block   interprete.AbstrExpr
	linea   int
	columna int
}

// Constructor for NT_InsFor
func NewNT_InsFor(id string, exp interprete.AbstrExpr, block interprete.AbstrExpr, linea int, columna int) *NT_InsFor {
	o := new(NT_InsFor)
	o.id = id
	o.exp = exp
	o.block = block
	o.linea = linea
	o.columna = columna
	return o
}

func (ntFor *NT_InsFor) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expRecorrido := ntFor.exp.Interpretar(ctx)

	fmt.Println(expRecorrido)
	fmt.Println("Jello")
	if expRecorrido.Tipo == interprete.Nil {
		ctx.AddError("Error For: No se puede recorrer un Nil", ntFor.linea, ntFor.columna)
		return interprete.NewNil()
	}
	//creamos la variable con que vamos a recorrer el for
	if expRecorrido.Tipo == interprete.Rango {
		inicio := expRecorrido.Rango[0]
		ctx.AgregarConstanteExpresion(ntFor.id, interprete.Integer, interprete.Nativo, -1, nil, interprete.NewIntLiteral(inicio), nil, ntFor.linea, ntFor.columna)
	} else if expRecorrido.Tipo != interprete.Break && expRecorrido.Tipo != interprete.Continue && expRecorrido.Tipo != interprete.Vector {
		ctx.AgregarConstanteExpresion(ntFor.id, expRecorrido.Tipo, interprete.Nativo, -1, nil, expRecorrido, nil, ntFor.linea, ntFor.columna)
	}
	switch expRecorrido.Tipo {
	case interprete.Bool:
		ctx.AddError("Error For: No se puede recorrer un Bool", ntFor.linea, ntFor.columna)
		return interprete.NewNil()
	case interprete.Integer:
		ctx.AddError("Error For: No se puede recorrer un entero", ntFor.linea, ntFor.columna)
		return interprete.NewNil()
	case interprete.String:
		for i := 0; i < len(expRecorrido.ValorS); i++ {
			ctx.AgregarAmbito()
			tmp := interprete.NewStringLiteral(string(expRecorrido.ValorS[i]))
			if tmp.ValorS != "\"" {
				ctx.ReasignarValorNativoFor(ntFor.id, tmp, ntFor.linea, ntFor.columna)
				ntFor.block.Interpretar(ctx)
			} else if i < len(expRecorrido.ValorS)-1 && tmp.ValorS != "\\" && string(expRecorrido.ValorS[i+1]) == "\"" {
				tmp.ValorS = string(expRecorrido.ValorS[i+1])
				ctx.ReasignarValorNativoFor(ntFor.id, tmp, ntFor.linea, ntFor.columna)
				ntFor.block.Interpretar(ctx)
			}
			ctx.PopAmbito()
		}
		break
	case interprete.Character:
		ctx.AddError("Error For: No se puede recorrer un Caracter", ntFor.linea, ntFor.columna)
		return interprete.NewNil()
	case interprete.Float:
		ctx.AddError("Error For: No se puede recorrer un Float", ntFor.linea, ntFor.columna)
		return interprete.NewNil()
	case interprete.Rango:
		for i := expRecorrido.Rango[0]; i < expRecorrido.Rango[1]+1; i++ {
			ctx.AgregarAmbito()
			tmp := interprete.NewIntLiteral(i)
			ctx.ReasignarValorNativoFor(ntFor.id, tmp, ntFor.linea, ntFor.columna)
			ntFor.block.Interpretar(ctx)
			ctx.PopAmbito()
		}
		break
	case interprete.Vector:
		for i := 0; i < len(expRecorrido.Vector); i++ {
			ctx.AgregarAmbito()
			ctx.AgregarConstanteExpresion(ntFor.id, expRecorrido.TipoVector, interprete.Nativo, -1, nil, &expRecorrido.Vector[i], nil, ntFor.linea, ntFor.columna)
			ntFor.block.Interpretar(ctx)
			ctx.PopAmbito()
		}
		break
	}
	return interprete.NewNil()
}
