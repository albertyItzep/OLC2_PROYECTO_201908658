package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_InstWhile struct {
	expr  interprete.AbstrExpr
	block interprete.AbstrExpr
}

// Constructor for NT_InstWhile
func NewNT_InstWhile(expr interprete.AbstrExpr, block interprete.AbstrExpr) *NT_InstWhile {
	o := new(NT_InstWhile)
	o.expr = expr
	o.block = block
	return o
}

func (ntWhile *NT_InstWhile) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for true {
		exp := ntWhile.expr.Interpretar(ctx)
		if exp.Tipo != interprete.Bool {
			ctx.AddError("Erro While: tipo de verificacion incorrecto")
			return interprete.NewNil()
		}
		if !exp.ValorB {
			return interprete.NewNil()
		}
		ctx.AgregarAmbito()
		ntWhile.block.Interpretar(ctx)
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
