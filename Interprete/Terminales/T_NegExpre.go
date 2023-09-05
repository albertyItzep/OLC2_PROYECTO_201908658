package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
)

type T_NegExp struct {
	expOpe interprete.AbstrExpr
}

// Constructor for T_NegExp
func NewT_NegExp(expOpe interprete.AbstrExpr) *T_NegExp {
	o := new(T_NegExp)
	o.expOpe = expOpe
	return o
}

func (tNegExp *T_NegExp) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expre := tNegExp.expOpe.Interpretar(ctx)
	fmt.Println(expre.Tipo)
	switch expre.Tipo {
	case interprete.Integer:
		expre.Valor = expre.Valor * -1
		return interprete.NewIntLiteral(expre.Valor)
	case interprete.Float:
		expre.ValorF = expre.ValorF * -1
		return interprete.NewFloatLiteral(expre.ValorF)
	default:
		ctx.AddError("Error: Negacion unaria incorrecta")
		return interprete.NewNil()
	}
}
