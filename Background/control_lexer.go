// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ControlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ControlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func controllexerLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'var'", "':'", "'?'", "'let'", "'!'", "'*'", "'/'", "'+'",
		"'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "ID", "INT", "DECIMAL",
		"STRING", "WS", "LINEALCOMENT", "DATATYPE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "ID", "INT", "DECIMAL", "STRING", "WS", "LINEALCOMENT", "DATATYPE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 138, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 62, 8, 10, 10, 10, 12, 10, 65,
		9, 10, 1, 11, 4, 11, 68, 8, 11, 11, 11, 12, 11, 69, 1, 12, 4, 12, 73, 8,
		12, 11, 12, 12, 12, 74, 1, 12, 1, 12, 4, 12, 79, 8, 12, 11, 12, 12, 12,
		80, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 87, 8, 13, 10, 13, 12, 13, 90, 9,
		13, 1, 13, 1, 13, 1, 14, 4, 14, 95, 8, 14, 11, 14, 12, 14, 96, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 105, 8, 15, 10, 15, 12, 15, 108,
		9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 137, 8, 16, 0,
		0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 5, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48,
		57, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 9, 10, 32, 32, 149, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 41, 1, 0, 0, 0,
		7, 43, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13, 51, 1, 0,
		0, 0, 15, 53, 1, 0, 0, 0, 17, 55, 1, 0, 0, 0, 19, 57, 1, 0, 0, 0, 21, 59,
		1, 0, 0, 0, 23, 67, 1, 0, 0, 0, 25, 72, 1, 0, 0, 0, 27, 82, 1, 0, 0, 0,
		29, 94, 1, 0, 0, 0, 31, 100, 1, 0, 0, 0, 33, 136, 1, 0, 0, 0, 35, 36, 5,
		61, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 118, 0, 0, 38, 39, 5, 97, 0, 0,
		39, 40, 5, 114, 0, 0, 40, 4, 1, 0, 0, 0, 41, 42, 5, 58, 0, 0, 42, 6, 1,
		0, 0, 0, 43, 44, 5, 63, 0, 0, 44, 8, 1, 0, 0, 0, 45, 46, 5, 108, 0, 0,
		46, 47, 5, 101, 0, 0, 47, 48, 5, 116, 0, 0, 48, 10, 1, 0, 0, 0, 49, 50,
		5, 33, 0, 0, 50, 12, 1, 0, 0, 0, 51, 52, 5, 42, 0, 0, 52, 14, 1, 0, 0,
		0, 53, 54, 5, 47, 0, 0, 54, 16, 1, 0, 0, 0, 55, 56, 5, 43, 0, 0, 56, 18,
		1, 0, 0, 0, 57, 58, 5, 45, 0, 0, 58, 20, 1, 0, 0, 0, 59, 63, 7, 0, 0, 0,
		60, 62, 7, 1, 0, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 22, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66,
		68, 7, 2, 0, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0,
		0, 69, 70, 1, 0, 0, 0, 70, 24, 1, 0, 0, 0, 71, 73, 7, 2, 0, 0, 72, 71,
		1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 78, 5, 46, 0, 0, 77, 79, 7, 2, 0, 0, 78, 77, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81,
		26, 1, 0, 0, 0, 82, 88, 5, 34, 0, 0, 83, 87, 8, 3, 0, 0, 84, 85, 5, 34,
		0, 0, 85, 87, 5, 34, 0, 0, 86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87,
		90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0,
		0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 34, 0, 0, 92, 28, 1, 0, 0, 0, 93, 95,
		7, 4, 0, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 6, 14, 0, 0, 99, 30, 1,
		0, 0, 0, 100, 101, 5, 47, 0, 0, 101, 102, 5, 47, 0, 0, 102, 106, 1, 0,
		0, 0, 103, 105, 9, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0,
		106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 32, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 110, 5, 83, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5,
		114, 0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 110, 0, 0, 114, 137, 5,
		103, 0, 0, 115, 116, 5, 70, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5,
		111, 0, 0, 118, 119, 5, 97, 0, 0, 119, 137, 5, 116, 0, 0, 120, 121, 5,
		73, 0, 0, 121, 122, 5, 110, 0, 0, 122, 137, 5, 116, 0, 0, 123, 124, 5,
		66, 0, 0, 124, 125, 5, 111, 0, 0, 125, 126, 5, 111, 0, 0, 126, 137, 5,
		108, 0, 0, 127, 128, 5, 67, 0, 0, 128, 129, 5, 104, 0, 0, 129, 130, 5,
		97, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 99,
		0, 0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 101, 0, 0, 135, 137, 5, 114,
		0, 0, 136, 109, 1, 0, 0, 0, 136, 115, 1, 0, 0, 0, 136, 120, 1, 0, 0, 0,
		136, 123, 1, 0, 0, 0, 136, 127, 1, 0, 0, 0, 137, 34, 1, 0, 0, 0, 10, 0,
		63, 69, 74, 80, 86, 88, 96, 106, 136, 1, 6, 0, 0,
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

// ControlLexerInit initializes any static state used to implement ControlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewControlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlLexerInit() {
	staticData := &ControlLexerLexerStaticData
	staticData.once.Do(controllexerLexerInit)
}

// NewControlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewControlLexer(input antlr.CharStream) *ControlLexer {
	ControlLexerInit()
	l := new(ControlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ControlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Control.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ControlLexer tokens.
const (
	ControlLexerT__0         = 1
	ControlLexerT__1         = 2
	ControlLexerT__2         = 3
	ControlLexerT__3         = 4
	ControlLexerT__4         = 5
	ControlLexerT__5         = 6
	ControlLexerT__6         = 7
	ControlLexerT__7         = 8
	ControlLexerT__8         = 9
	ControlLexerT__9         = 10
	ControlLexerID           = 11
	ControlLexerINT          = 12
	ControlLexerDECIMAL      = 13
	ControlLexerSTRING       = 14
	ControlLexerWS           = 15
	ControlLexerLINEALCOMENT = 16
	ControlLexerDATATYPE     = 17
)
