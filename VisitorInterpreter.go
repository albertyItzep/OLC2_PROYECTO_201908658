package main

import (
	"OLC2_PROYECTO_201908658/Background"

	"github.com/antlr4-go/antlr"
)

type Visitorinterprete struct {
}

func (vI *Visitorinterprete) Visit(tree antlr.ParseTree) interface{} {
	panic("not implemented") // TODO: Implement
}

func (vI *Visitorinterprete) VisitChildren(node antlr.RuleNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (vI *Visitorinterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (vI *Visitorinterprete) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#prog.
func (vI *Visitorinterprete) VisitProg(ctx *Background.ProgContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#block.
func (vI *Visitorinterprete) VisitBlock(ctx *Background.BlockContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#stmt.
func (vI *Visitorinterprete) VisitStmt(ctx *Background.StmtContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#assing.
func (vI *Visitorinterprete) VisitAssing(ctx *Background.AssingContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#declaration.
func (vI *Visitorinterprete) VisitDeclaration(ctx *Background.DeclarationContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#LINT.
func (vI *Visitorinterprete) VisitLINT(ctx *Background.LINTContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#NEGACION.
func (vI *Visitorinterprete) VisitNEGACION(ctx *Background.NEGACIONContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#LDECIMAL.
func (vI *Visitorinterprete) VisitLDECIMAL(ctx *Background.LDECIMALContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#OPSUMRES.
func (vI *Visitorinterprete) VisitOPSUMRES(ctx *Background.OPSUMRESContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#OPMULDIV.
func (vI *Visitorinterprete) VisitOPMULDIV(ctx *Background.OPMULDIVContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#LSTRING.
func (vI *Visitorinterprete) VisitLSTRING(ctx *Background.LSTRINGContext) interface{} {
	panic("not implemented") // TODO: Implement
}
