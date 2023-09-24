package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_InstSwitch struct {
	exp         interprete.AbstrExpr
	InstCase    []NT_InstCase
	InstDefault interprete.AbstrExpr
	linea       int
	columna     int
}

// Constructor for NT_InstSwitch
func NewNT_InstSwitch(exp interprete.AbstrExpr, linea int, columna int) *NT_InstSwitch {
	o := new(NT_InstSwitch)
	o.exp = exp
	o.InstCase = make([]NT_InstCase, 0)
	o.InstDefault = nil
	o.linea = linea
	o.columna = columna
	return o
}

func (ntSwitch *NT_InstSwitch) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := ntSwitch.exp.Interpretar(ctx)
	caseTrue := false
	for _, v := range ntSwitch.InstCase {
		tmp := v.exp.Interpretar(ctx)
		if expr.Tipo != tmp.Tipo {
			ctx.AddError("Error Switch: tipos incompatibles para comparacion", ntSwitch.linea, ntSwitch.columna)
		}
		switch expr.Tipo {
		case interprete.Bool:
			if expr.ValorB == tmp.ValorB {
				ctx.AgregarAmbito()
				v.block.Interpretar(ctx)
				caseTrue = true
				ctx.PopAmbito()
				return interprete.NewNil()
			}
		case interprete.Integer:
			if expr.Valor == tmp.Valor {
				ctx.AgregarAmbito()
				v.block.Interpretar(ctx)
				caseTrue = true
				ctx.PopAmbito()
				return interprete.NewNil()
			}
		case interprete.String:
			if expr.ValorS == tmp.ValorS {
				ctx.AgregarAmbito()
				v.block.Interpretar(ctx)
				caseTrue = true
				ctx.PopAmbito()
				return interprete.NewNil()
			}
		case interprete.Character:
			if expr.ValorC == tmp.ValorC {
				ctx.AgregarAmbito()
				v.block.Interpretar(ctx)
				caseTrue = true
				ctx.PopAmbito()
				return interprete.NewNil()
			}
		case interprete.Float:
			if expr.ValorF == tmp.ValorF {
				ctx.AgregarAmbito()
				v.block.Interpretar(ctx)
				caseTrue = true
				ctx.PopAmbito()
				return interprete.NewNil()
			}
		}
	}
	if ntSwitch.InstDefault != nil && !caseTrue {
		ctx.AgregarAmbito()
		ntSwitch.InstDefault.Interpretar(ctx)
		ctx.PopAmbito()
		return interprete.NewNil()
	}
	return interprete.NewNil()
}

func (ntSwitch *NT_InstSwitch) AddCase(cases NT_InstCase) {
	ntSwitch.InstCase = append(ntSwitch.InstCase, cases)
}

func (ntSwitch *NT_InstSwitch) AddDefault(exp interprete.AbstrExpr) {
	ntSwitch.InstDefault = exp
}

type NT_InstCase struct {
	exp   interprete.AbstrExpr
	block interprete.AbstrExpr
}

// Constructor for NT_InstCase
func NewNT_InstCase(exp interprete.AbstrExpr, block interprete.AbstrExpr) *NT_InstCase {
	o := new(NT_InstCase)
	o.exp = exp
	o.block = block
	return o
}

func (ntCases *NT_InstCase) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ntCases.block.Interpretar(ctx)
	return interprete.NewNil()
}

type NT_InstDefault struct {
	block interprete.AbstrExpr
}

// Constructor for NT_InstDefault
func NewNT_InstDefault(block interprete.AbstrExpr) *NT_InstDefault {
	o := new(NT_InstDefault)
	o.block = block
	return o
}

func (ntDefault *NT_InstDefault) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return ntDefault.block.Interpretar(ctx)
}
