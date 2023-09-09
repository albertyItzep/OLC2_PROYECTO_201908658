package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type T_Rango struct {
	posInicial interprete.AbstrExpr
	posFinal   interprete.AbstrExpr
	linea      int
	columna    int
}

// Constructor for T_Rango
func NewT_Rango(posInicial interprete.AbstrExpr, posFinal interprete.AbstrExpr, linea int, columna int) *T_Rango {
	o := new(T_Rango)
	o.posInicial = posInicial
	o.posFinal = posFinal
	o.linea = linea
	o.columna = columna
	return o
}

func (tRango *T_Rango) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	iniciaPos := tRango.posInicial.Interpretar(ctx)
	finalPos := tRango.posFinal.Interpretar(ctx)
	if iniciaPos.Tipo != interprete.Integer || finalPos.Tipo != interprete.Integer {
		ctx.AddError("Error Rango: Tipos invalidos")
	}
	rango := interprete.NewRangoLiteral()

	rango.Rango = append(rango.Rango, iniciaPos.Valor)
	rango.Rango = append(rango.Rango, finalPos.Valor)
	return rango
}
