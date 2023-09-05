package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
	"strconv"
)

type NT_Print struct {
	listaExpresiones []interprete.AbstrExpr
}

func NewNT_Print() *NT_Print {
	return &NT_Print{
		listaExpresiones: make([]interprete.AbstrExpr, 0),
	}
}

func (ntPrint *NT_Print) AddExpresion(sentencia interprete.AbstrExpr) {
	ntPrint.listaExpresiones = append(ntPrint.listaExpresiones, sentencia)
}

func (ntP *NT_Print) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, expr := range ntP.listaExpresiones {
		res := expr.Interpretar(ctx)
		switch res.Tipo {
		case interprete.Nil:
			ctx.Print("Nil")
		case interprete.Integer:
			ctx.Print(strconv.Itoa(res.Valor))
		case interprete.Float:
			ctx.Print(fmt.Sprintf("%f", res.ValorF))
		case interprete.String:
			ctx.Print(res.ValorS)
		case interprete.Character:
			ctx.Print(res.ValorC)
		case interprete.Bool:
			if res.ValorB {
				ctx.Print("true")
			} else {
				ctx.Print("false")
			}
		}
	}
	ctx.Println("")
	return interprete.NewNil()
}
