// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

type BaseControlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseControlVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAssing(ctx *AssingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLINT(ctx *LINTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNEGACION(ctx *NEGACIONContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLDECIMAL(ctx *LDECIMALContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOPSUMRES(ctx *OPSUMRESContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOPMULDIV(ctx *OPMULDIVContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLSTRING(ctx *LSTRINGContext) interface{} {
	return v.VisitChildren(ctx)
}
