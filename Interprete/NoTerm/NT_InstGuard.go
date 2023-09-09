package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_InstGuard struct {
	exp   interprete.AbstrExpr
	block interprete.AbstrExpr
}

// Constructor for NT_InstGuard
func NewNT_InstGuard(exp interprete.AbstrExpr, block interprete.AbstrExpr) *NT_InstGuard {
	o := new(NT_InstGuard)
	o.exp = exp
	o.block = block
	return o
}

func (ntGuard *NT_InstGuard) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntGuard.exp.Interpretar(ctx)

	if exp.Tipo != interprete.Bool {
		ctx.AddError("Error Guard: Tipo de expresion invalida para comparacion")
		return interprete.NewNil()
	}

	if !exp.ValorB {
		ctx.AgregarAmbito()
		ntGuard.block.Interpretar(ctx)
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
