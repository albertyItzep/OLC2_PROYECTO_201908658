// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ControlParser struct {
	*antlr.BaseParser
}

var ControlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func controlParserInit() {
	staticData := &ControlParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'var'", "':'", "'?'", "'let'", "'!'", "'*'", "'/'", "'+'",
		"'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "ID", "INT", "DECIMAL",
		"STRING", "WS", "LINEALCOMENT", "DATATYPE",
	}
	staticData.RuleNames = []string{
		"prog", "block", "stmt", "assing", "declaration", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 17, 8, 1, 10, 1, 12, 1, 20,
		9, 1, 1, 2, 1, 2, 3, 2, 24, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 55,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 63, 8, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 71, 8, 5, 10, 5, 12, 5, 74, 9, 5, 1, 5, 0, 1,
		10, 6, 0, 2, 4, 6, 8, 10, 0, 2, 1, 0, 7, 8, 1, 0, 9, 10, 80, 0, 12, 1,
		0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 25, 1, 0, 0, 0, 8, 54,
		1, 0, 0, 0, 10, 62, 1, 0, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 0, 0, 1,
		14, 1, 1, 0, 0, 0, 15, 17, 3, 4, 2, 0, 16, 15, 1, 0, 0, 0, 17, 20, 1, 0,
		0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 3, 1, 0, 0, 0, 20, 18,
		1, 0, 0, 0, 21, 24, 3, 8, 4, 0, 22, 24, 3, 6, 3, 0, 23, 21, 1, 0, 0, 0,
		23, 22, 1, 0, 0, 0, 24, 5, 1, 0, 0, 0, 25, 26, 5, 11, 0, 0, 26, 27, 5,
		1, 0, 0, 27, 28, 3, 10, 5, 0, 28, 7, 1, 0, 0, 0, 29, 30, 5, 2, 0, 0, 30,
		31, 5, 11, 0, 0, 31, 32, 5, 1, 0, 0, 32, 55, 3, 10, 5, 0, 33, 34, 5, 2,
		0, 0, 34, 35, 5, 11, 0, 0, 35, 36, 5, 3, 0, 0, 36, 37, 5, 17, 0, 0, 37,
		55, 5, 4, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 5, 11, 0, 0, 40, 41, 5, 3,
		0, 0, 41, 42, 5, 17, 0, 0, 42, 43, 5, 1, 0, 0, 43, 55, 3, 10, 5, 0, 44,
		45, 5, 5, 0, 0, 45, 46, 5, 11, 0, 0, 46, 47, 5, 1, 0, 0, 47, 55, 3, 10,
		5, 0, 48, 49, 5, 5, 0, 0, 49, 50, 5, 11, 0, 0, 50, 51, 5, 3, 0, 0, 51,
		52, 5, 17, 0, 0, 52, 53, 5, 1, 0, 0, 53, 55, 3, 10, 5, 0, 54, 29, 1, 0,
		0, 0, 54, 33, 1, 0, 0, 0, 54, 38, 1, 0, 0, 0, 54, 44, 1, 0, 0, 0, 54, 48,
		1, 0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 57, 6, 5, -1, 0, 57, 58, 5, 6, 0, 0,
		58, 63, 3, 10, 5, 6, 59, 63, 5, 12, 0, 0, 60, 63, 5, 13, 0, 0, 61, 63,
		5, 14, 0, 0, 62, 56, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0,
		62, 61, 1, 0, 0, 0, 63, 72, 1, 0, 0, 0, 64, 65, 10, 5, 0, 0, 65, 66, 7,
		0, 0, 0, 66, 71, 3, 10, 5, 6, 67, 68, 10, 4, 0, 0, 68, 69, 7, 1, 0, 0,
		69, 71, 3, 10, 5, 5, 70, 64, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 71, 74, 1,
		0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 11, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 6, 18, 23, 54, 62, 70, 72,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ControlParserInit initializes any static state used to implement ControlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewControlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlParserInit() {
	staticData := &ControlParserStaticData
	staticData.once.Do(controlParserInit)
}

// NewControlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewControlParser(input antlr.TokenStream) *ControlParser {
	ControlParserInit()
	this := new(ControlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ControlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Control.g4"

	return this
}

// ControlParser tokens.
const (
	ControlParserEOF          = antlr.TokenEOF
	ControlParserT__0         = 1
	ControlParserT__1         = 2
	ControlParserT__2         = 3
	ControlParserT__3         = 4
	ControlParserT__4         = 5
	ControlParserT__5         = 6
	ControlParserT__6         = 7
	ControlParserT__7         = 8
	ControlParserT__8         = 9
	ControlParserT__9         = 10
	ControlParserID           = 11
	ControlParserINT          = 12
	ControlParserDECIMAL      = 13
	ControlParserSTRING       = 14
	ControlParserWS           = 15
	ControlParserLINEALCOMENT = 16
	ControlParserDATATYPE     = 17
)

// ControlParser rules.
const (
	ControlParserRULE_prog        = 0
	ControlParserRULE_block       = 1
	ControlParserRULE_stmt        = 2
	ControlParserRULE_assing      = 3
	ControlParserRULE_declaration = 4
	ControlParserRULE_expr        = 5
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ControlParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ControlParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Block()
	}
	{
		p.SetState(13)
		p.Match(ControlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ControlParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2084) != 0 {
		{
			p.SetState(15)
			p.Stmt()
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	Assing() IAssingContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StmtContext) Assing() IAssingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssingContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ControlParserRULE_stmt)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__1, ControlParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Declaration()
		}

	case ControlParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Assing()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssingContext is an interface to support dynamic dispatch.
type IAssingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAssingContext differentiates from other interfaces.
	IsAssingContext()
}

type AssingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssingContext() *AssingContext {
	var p = new(AssingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assing
	return p
}

func InitEmptyAssingContext(p *AssingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assing
}

func (*AssingContext) IsAssingContext() {}

func NewAssingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssingContext {
	var p = new(AssingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_assing

	return p
}

func (s *AssingContext) GetParser() antlr.Parser { return s.parser }

func (s *AssingContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AssingContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAssing(s)
	}
}

func (s *AssingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAssing(s)
	}
}

func (s *AssingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAssing(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Assing() (localctx IAssingContext) {
	localctx = NewAssingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ControlParserRULE_assing)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(ControlParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(26)
		p.Match(ControlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext
	DATATYPE() antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *DeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationContext) DATATYPE() antlr.TerminalNode {
	return s.GetToken(ControlParserDATATYPE, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ControlParserRULE_declaration)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(ControlParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(ControlParserDATATYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(ControlParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Match(ControlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(ControlParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Match(ControlParserDATATYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.expr(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.Match(ControlParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.expr(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(ControlParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(ControlParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(ControlParserDATATYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(ControlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LINTContext struct {
	ExprContext
}

func NewLINTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LINTContext {
	var p = new(LINTContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LINTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LINTContext) INT() antlr.TerminalNode {
	return s.GetToken(ControlParserINT, 0)
}

func (s *LINTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterLINT(s)
	}
}

func (s *LINTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitLINT(s)
	}
}

func (s *LINTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitLINT(s)

	default:
		return t.VisitChildren(s)
	}
}

type NEGACIONContext struct {
	ExprContext
	rigth IExprContext
}

func NewNEGACIONContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NEGACIONContext {
	var p = new(NEGACIONContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NEGACIONContext) GetRigth() IExprContext { return s.rigth }

func (s *NEGACIONContext) SetRigth(v IExprContext) { s.rigth = v }

func (s *NEGACIONContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NEGACIONContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NEGACIONContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNEGACION(s)
	}
}

func (s *NEGACIONContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNEGACION(s)
	}
}

func (s *NEGACIONContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNEGACION(s)

	default:
		return t.VisitChildren(s)
	}
}

type LDECIMALContext struct {
	ExprContext
}

func NewLDECIMALContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LDECIMALContext {
	var p = new(LDECIMALContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LDECIMALContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LDECIMALContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ControlParserDECIMAL, 0)
}

func (s *LDECIMALContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterLDECIMAL(s)
	}
}

func (s *LDECIMALContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitLDECIMAL(s)
	}
}

func (s *LDECIMALContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitLDECIMAL(s)

	default:
		return t.VisitChildren(s)
	}
}

type OPSUMRESContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	rigth IExprContext
}

func NewOPSUMRESContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OPSUMRESContext {
	var p = new(OPSUMRESContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OPSUMRESContext) GetOp() antlr.Token { return s.op }

func (s *OPSUMRESContext) SetOp(v antlr.Token) { s.op = v }

func (s *OPSUMRESContext) GetLeft() IExprContext { return s.left }

func (s *OPSUMRESContext) GetRigth() IExprContext { return s.rigth }

func (s *OPSUMRESContext) SetLeft(v IExprContext) { s.left = v }

func (s *OPSUMRESContext) SetRigth(v IExprContext) { s.rigth = v }

func (s *OPSUMRESContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OPSUMRESContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OPSUMRESContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OPSUMRESContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOPSUMRES(s)
	}
}

func (s *OPSUMRESContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOPSUMRES(s)
	}
}

func (s *OPSUMRESContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOPSUMRES(s)

	default:
		return t.VisitChildren(s)
	}
}

type OPMULDIVContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	rigth IExprContext
}

func NewOPMULDIVContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OPMULDIVContext {
	var p = new(OPMULDIVContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OPMULDIVContext) GetOp() antlr.Token { return s.op }

func (s *OPMULDIVContext) SetOp(v antlr.Token) { s.op = v }

func (s *OPMULDIVContext) GetLeft() IExprContext { return s.left }

func (s *OPMULDIVContext) GetRigth() IExprContext { return s.rigth }

func (s *OPMULDIVContext) SetLeft(v IExprContext) { s.left = v }

func (s *OPMULDIVContext) SetRigth(v IExprContext) { s.rigth = v }

func (s *OPMULDIVContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OPMULDIVContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OPMULDIVContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OPMULDIVContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOPMULDIV(s)
	}
}

func (s *OPMULDIVContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOPMULDIV(s)
	}
}

func (s *OPMULDIVContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOPMULDIV(s)

	default:
		return t.VisitChildren(s)
	}
}

type LSTRINGContext struct {
	ExprContext
}

func NewLSTRINGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LSTRINGContext {
	var p = new(LSTRINGContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LSTRINGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LSTRINGContext) STRING() antlr.TerminalNode {
	return s.GetToken(ControlParserSTRING, 0)
}

func (s *LSTRINGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterLSTRING(s)
	}
}

func (s *LSTRINGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitLSTRING(s)
	}
}

func (s *LSTRINGContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitLSTRING(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ControlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ControlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__5:
		localctx = NewNEGACIONContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(57)
			p.Match(ControlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)

			var _x = p.expr(6)

			localctx.(*NEGACIONContext).rigth = _x
		}

	case ControlParserINT:
		localctx = NewLINTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Match(ControlParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserDECIMAL:
		localctx = NewLDECIMALContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(ControlParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserSTRING:
		localctx = NewLSTRINGContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(ControlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOPMULDIVContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OPMULDIVContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(65)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OPMULDIVContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__6 || _la == ControlParserT__7) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OPMULDIVContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)

					var _x = p.expr(6)

					localctx.(*OPMULDIVContext).rigth = _x
				}

			case 2:
				localctx = NewOPSUMRESContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OPSUMRESContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(68)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OPSUMRESContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__8 || _la == ControlParserT__9) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OPSUMRESContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)

					var _x = p.expr(5)

					localctx.(*OPSUMRESContext).rigth = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ControlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ControlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
