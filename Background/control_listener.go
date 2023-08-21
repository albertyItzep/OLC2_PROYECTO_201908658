// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// ControlListener is a complete listener for a parse tree produced by ControlParser.
type ControlListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssing is called when entering the assing production.
	EnterAssing(c *AssingContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterLINT is called when entering the LINT production.
	EnterLINT(c *LINTContext)

	// EnterNEGACION is called when entering the NEGACION production.
	EnterNEGACION(c *NEGACIONContext)

	// EnterLDECIMAL is called when entering the LDECIMAL production.
	EnterLDECIMAL(c *LDECIMALContext)

	// EnterOPSUMRES is called when entering the OPSUMRES production.
	EnterOPSUMRES(c *OPSUMRESContext)

	// EnterOPMULDIV is called when entering the OPMULDIV production.
	EnterOPMULDIV(c *OPMULDIVContext)

	// EnterLSTRING is called when entering the LSTRING production.
	EnterLSTRING(c *LSTRINGContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssing is called when exiting the assing production.
	ExitAssing(c *AssingContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitLINT is called when exiting the LINT production.
	ExitLINT(c *LINTContext)

	// ExitNEGACION is called when exiting the NEGACION production.
	ExitNEGACION(c *NEGACIONContext)

	// ExitLDECIMAL is called when exiting the LDECIMAL production.
	ExitLDECIMAL(c *LDECIMALContext)

	// ExitOPSUMRES is called when exiting the OPSUMRES production.
	ExitOPSUMRES(c *OPSUMRESContext)

	// ExitOPMULDIV is called when exiting the OPMULDIV production.
	ExitOPMULDIV(c *OPMULDIVContext)

	// ExitLSTRING is called when exiting the LSTRING production.
	ExitLSTRING(c *LSTRINGContext)
}
