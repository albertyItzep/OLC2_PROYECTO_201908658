package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type T_NegExp struct {
	expOpe  interprete.AbstrExpr
	Linea   int
	Columna int
}

// Constructor for T_NegExp
func NewT_NegExp(expOpe interprete.AbstrExpr, linea int, columna int) *T_NegExp {
	o := new(T_NegExp)
	o.expOpe = expOpe
	o.Linea = linea
	o.Columna = columna
	return o
}

func (tNegExp *T_NegExp) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expre := tNegExp.expOpe.Interpretar(ctx)
	switch expre.Tipo {
	case interprete.Integer:
		expre.Valor = expre.Valor * -1
		return interprete.NewIntLiteral(expre.Valor)
	case interprete.Float:
		expre.ValorF = expre.ValorF * -1
		return interprete.NewFloatLiteral(expre.ValorF)
	default:
		ctx.AddError("Error: Negacion unaria incorrecta", tNegExp.Linea, tNegExp.Columna)
		return interprete.NewNil()
	}
}
