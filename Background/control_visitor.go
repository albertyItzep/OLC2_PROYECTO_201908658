// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ControlParser.
type ControlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ControlParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ControlParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ControlParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by ControlParser#assing.
	VisitAssing(ctx *AssingContext) interface{}

	// Visit a parse tree produced by ControlParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ControlParser#LINT.
	VisitLINT(ctx *LINTContext) interface{}

	// Visit a parse tree produced by ControlParser#NEGACION.
	VisitNEGACION(ctx *NEGACIONContext) interface{}

	// Visit a parse tree produced by ControlParser#LDECIMAL.
	VisitLDECIMAL(ctx *LDECIMALContext) interface{}

	// Visit a parse tree produced by ControlParser#OPSUMRES.
	VisitOPSUMRES(ctx *OPSUMRESContext) interface{}

	// Visit a parse tree produced by ControlParser#OPMULDIV.
	VisitOPMULDIV(ctx *OPMULDIVContext) interface{}

	// Visit a parse tree produced by ControlParser#LSTRING.
	VisitLSTRING(ctx *LSTRINGContext) interface{}
}
