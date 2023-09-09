package noterm

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
)

type NT_InstIF struct {
	exp        interprete.AbstrExpr
	bloqueInst interprete.AbstrExpr
	elseIfIns  []interprete.AbstrExpr
	elseIns    interprete.AbstrExpr
}

// Constructor for NT_InstIF
func NewNT_InstIF(exp interprete.AbstrExpr, bloqueInst interprete.AbstrExpr) *NT_InstIF {
	o := new(NT_InstIF)
	o.exp = exp
	o.bloqueInst = bloqueInst
	o.elseIfIns = make([]interprete.AbstrExpr, 0)
	o.elseIns = nil
	return o
}

func (ntIFG *NT_InstIF) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntIFG.exp.Interpretar(ctx)

	if exp.Tipo != interprete.Bool {
		ctx.AddError("Error IF: Expresion tipo incompatible")
		return interprete.NewNil()
	}
	if exp.ValorB {
		ctx.AgregarAmbito()
		ntIFG.bloqueInst.Interpretar(ctx)
		ctx.PopAmbito()
	} else if len(ntIFG.elseIfIns) > 0 {
		for _, v := range ntIFG.elseIfIns {
			ctx.AgregarAmbito()
			res := v.Interpretar(ctx)
			if res.Tipo == interprete.Bool {
				return interprete.NewNil()
			}
			ctx.PopAmbito()
		}
		if ntIFG.elseIns != nil {
			ctx.AgregarAmbito()
			ntIFG.elseIns.Interpretar(ctx)
			ctx.PopAmbito()
		}
	} else if ntIFG.elseIns != nil {
		ctx.AgregarAmbito()
		ntIFG.elseIns.Interpretar(ctx)
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}

func (ntIG *NT_InstIF) AddElseIF(exp interprete.AbstrExpr) {
	ntIG.elseIfIns = append(ntIG.elseIfIns, exp)
}

func (ntIG *NT_InstIF) AddElse(exp interprete.AbstrExpr) {
	ntIG.elseIns = exp
}

type NT_InsElseIf struct {
	exp        interprete.AbstrExpr
	bloqueInst interprete.AbstrExpr
}

// Constructor for NT_InsElseIf
func NewNT_InsElseIf(exp interprete.AbstrExpr, bloqueInst interprete.AbstrExpr) *NT_InsElseIf {
	o := new(NT_InsElseIf)
	o.exp = exp
	o.bloqueInst = bloqueInst
	return o
}

func (ntElseIf *NT_InsElseIf) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := ntElseIf.exp.Interpretar(ctx)

	if exp.Tipo != interprete.Bool {
		ctx.AddError("Error IF: Expresion tipo incompatible")
		return interprete.NewNil()
	}

	if exp.ValorB {
		ntElseIf.bloqueInst.Interpretar(ctx)
		return interprete.NewBoolLiteral(true)
	}
	return interprete.NewNil()
}

type NT_InsElse struct {
	bloqInst interprete.AbstrExpr
}

// Constructor for NT_InsElse
func NewNT_InsElse(bloqInst interprete.AbstrExpr) *NT_InsElse {
	o := new(NT_InsElse)
	o.bloqInst = bloqInst
	return o
}

func (ntElse *NT_InsElse) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ntElse.bloqInst.Interpretar(ctx)
	return interprete.NewNil()
}
