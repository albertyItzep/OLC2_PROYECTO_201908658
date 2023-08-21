// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// BaseControlListener is a complete listener for a parse tree produced by ControlParser.
type BaseControlListener struct{}

var _ ControlListener = &BaseControlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseControlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseControlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseControlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseControlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseControlListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseControlListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseControlListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseControlListener) ExitBlock(ctx *BlockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseControlListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseControlListener) ExitStmt(ctx *StmtContext) {}

// EnterAssing is called when production assing is entered.
func (s *BaseControlListener) EnterAssing(ctx *AssingContext) {}

// ExitAssing is called when production assing is exited.
func (s *BaseControlListener) ExitAssing(ctx *AssingContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseControlListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseControlListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterLINT is called when production LINT is entered.
func (s *BaseControlListener) EnterLINT(ctx *LINTContext) {}

// ExitLINT is called when production LINT is exited.
func (s *BaseControlListener) ExitLINT(ctx *LINTContext) {}

// EnterNEGACION is called when production NEGACION is entered.
func (s *BaseControlListener) EnterNEGACION(ctx *NEGACIONContext) {}

// ExitNEGACION is called when production NEGACION is exited.
func (s *BaseControlListener) ExitNEGACION(ctx *NEGACIONContext) {}

// EnterLDECIMAL is called when production LDECIMAL is entered.
func (s *BaseControlListener) EnterLDECIMAL(ctx *LDECIMALContext) {}

// ExitLDECIMAL is called when production LDECIMAL is exited.
func (s *BaseControlListener) ExitLDECIMAL(ctx *LDECIMALContext) {}

// EnterOPSUMRES is called when production OPSUMRES is entered.
func (s *BaseControlListener) EnterOPSUMRES(ctx *OPSUMRESContext) {}

// ExitOPSUMRES is called when production OPSUMRES is exited.
func (s *BaseControlListener) ExitOPSUMRES(ctx *OPSUMRESContext) {}

// EnterOPMULDIV is called when production OPMULDIV is entered.
func (s *BaseControlListener) EnterOPMULDIV(ctx *OPMULDIVContext) {}

// ExitOPMULDIV is called when production OPMULDIV is exited.
func (s *BaseControlListener) ExitOPMULDIV(ctx *OPMULDIVContext) {}

// EnterLSTRING is called when production LSTRING is entered.
func (s *BaseControlListener) EnterLSTRING(ctx *LSTRINGContext) {}

// ExitLSTRING is called when production LSTRING is exited.
func (s *BaseControlListener) ExitLSTRING(ctx *LSTRINGContext) {}
