// Code generated from ClickhouseLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

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

type ClickhouseLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ClickhouseLexerLexerStaticData struct {
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

func clickhouselexerLexerInit() {
	staticData := &ClickhouseLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "'='", "'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "CREATE", "DATABASE", "IF", "NOT", "EXISTS", "ON", "CLUSTER", "ENGINE",
		"COMMENT", "ID", "EQUAL", "LEFT_P", "RIGHT_P", "COMMA", "NUMBER", "STRING",
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	staticData.RuleNames = []string{
		"CREATE", "DATABASE", "IF", "NOT", "EXISTS", "ON", "CLUSTER", "ENGINE",
		"COMMENT", "ID", "EQUAL", "LEFT_P", "RIGHT_P", "COMMA", "NUMBER", "STRING",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 222, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 144,
		8, 9, 10, 9, 12, 9, 147, 9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 4, 14, 158, 8, 14, 11, 14, 12, 14, 159, 1, 15, 1,
		15, 5, 15, 164, 8, 15, 10, 15, 12, 15, 167, 9, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 165,
		0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0, 35, 17, 37, 18,
		39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27,
		57, 28, 59, 29, 61, 30, 63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73, 36,
		75, 37, 77, 38, 79, 39, 81, 40, 83, 41, 1, 0, 30, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 39,
		39, 2, 0, 65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2,
		0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2,
		0, 71, 71, 103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2,
		0, 74, 74, 106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2,
		0, 77, 77, 109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2,
		0, 80, 80, 112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2,
		0, 83, 83, 115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2,
		0, 86, 86, 118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2,
		0, 89, 89, 121, 121, 2, 0, 90, 90, 122, 122, 223, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0,
		0, 0, 0, 83, 1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3, 92, 1, 0, 0, 0, 5, 101,
		1, 0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 108, 1, 0, 0, 0, 11, 115, 1, 0, 0, 0,
		13, 118, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17, 133, 1, 0, 0, 0, 19, 141,
		1, 0, 0, 0, 21, 148, 1, 0, 0, 0, 23, 150, 1, 0, 0, 0, 25, 152, 1, 0, 0,
		0, 27, 154, 1, 0, 0, 0, 29, 157, 1, 0, 0, 0, 31, 161, 1, 0, 0, 0, 33, 170,
		1, 0, 0, 0, 35, 172, 1, 0, 0, 0, 37, 174, 1, 0, 0, 0, 39, 176, 1, 0, 0,
		0, 41, 178, 1, 0, 0, 0, 43, 180, 1, 0, 0, 0, 45, 182, 1, 0, 0, 0, 47, 184,
		1, 0, 0, 0, 49, 186, 1, 0, 0, 0, 51, 188, 1, 0, 0, 0, 53, 190, 1, 0, 0,
		0, 55, 192, 1, 0, 0, 0, 57, 194, 1, 0, 0, 0, 59, 196, 1, 0, 0, 0, 61, 198,
		1, 0, 0, 0, 63, 200, 1, 0, 0, 0, 65, 202, 1, 0, 0, 0, 67, 204, 1, 0, 0,
		0, 69, 206, 1, 0, 0, 0, 71, 208, 1, 0, 0, 0, 73, 210, 1, 0, 0, 0, 75, 212,
		1, 0, 0, 0, 77, 214, 1, 0, 0, 0, 79, 216, 1, 0, 0, 0, 81, 218, 1, 0, 0,
		0, 83, 220, 1, 0, 0, 0, 85, 86, 3, 37, 18, 0, 86, 87, 3, 67, 33, 0, 87,
		88, 3, 41, 20, 0, 88, 89, 3, 33, 16, 0, 89, 90, 3, 71, 35, 0, 90, 91, 3,
		41, 20, 0, 91, 2, 1, 0, 0, 0, 92, 93, 3, 39, 19, 0, 93, 94, 3, 33, 16,
		0, 94, 95, 3, 71, 35, 0, 95, 96, 3, 33, 16, 0, 96, 97, 3, 35, 17, 0, 97,
		98, 3, 33, 16, 0, 98, 99, 3, 69, 34, 0, 99, 100, 3, 41, 20, 0, 100, 4,
		1, 0, 0, 0, 101, 102, 3, 49, 24, 0, 102, 103, 3, 43, 21, 0, 103, 6, 1,
		0, 0, 0, 104, 105, 3, 59, 29, 0, 105, 106, 3, 61, 30, 0, 106, 107, 3, 71,
		35, 0, 107, 8, 1, 0, 0, 0, 108, 109, 3, 41, 20, 0, 109, 110, 3, 79, 39,
		0, 110, 111, 3, 49, 24, 0, 111, 112, 3, 69, 34, 0, 112, 113, 3, 71, 35,
		0, 113, 114, 3, 69, 34, 0, 114, 10, 1, 0, 0, 0, 115, 116, 3, 61, 30, 0,
		116, 117, 3, 59, 29, 0, 117, 12, 1, 0, 0, 0, 118, 119, 3, 37, 18, 0, 119,
		120, 3, 55, 27, 0, 120, 121, 3, 73, 36, 0, 121, 122, 3, 69, 34, 0, 122,
		123, 3, 71, 35, 0, 123, 124, 3, 41, 20, 0, 124, 125, 3, 67, 33, 0, 125,
		14, 1, 0, 0, 0, 126, 127, 3, 41, 20, 0, 127, 128, 3, 59, 29, 0, 128, 129,
		3, 45, 22, 0, 129, 130, 3, 49, 24, 0, 130, 131, 3, 59, 29, 0, 131, 132,
		3, 41, 20, 0, 132, 16, 1, 0, 0, 0, 133, 134, 3, 37, 18, 0, 134, 135, 3,
		61, 30, 0, 135, 136, 3, 57, 28, 0, 136, 137, 3, 57, 28, 0, 137, 138, 3,
		41, 20, 0, 138, 139, 3, 59, 29, 0, 139, 140, 3, 71, 35, 0, 140, 18, 1,
		0, 0, 0, 141, 145, 7, 0, 0, 0, 142, 144, 7, 1, 0, 0, 143, 142, 1, 0, 0,
		0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		20, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 61, 0, 0, 149, 22, 1,
		0, 0, 0, 150, 151, 5, 40, 0, 0, 151, 24, 1, 0, 0, 0, 152, 153, 5, 41, 0,
		0, 153, 26, 1, 0, 0, 0, 154, 155, 5, 44, 0, 0, 155, 28, 1, 0, 0, 0, 156,
		158, 7, 2, 0, 0, 157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 30, 1, 0, 0, 0, 161, 165, 5, 39,
		0, 0, 162, 164, 8, 3, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 169, 5, 39, 0, 0, 169, 32, 1, 0, 0, 0, 170, 171,
		7, 4, 0, 0, 171, 34, 1, 0, 0, 0, 172, 173, 7, 5, 0, 0, 173, 36, 1, 0, 0,
		0, 174, 175, 7, 6, 0, 0, 175, 38, 1, 0, 0, 0, 176, 177, 7, 7, 0, 0, 177,
		40, 1, 0, 0, 0, 178, 179, 7, 8, 0, 0, 179, 42, 1, 0, 0, 0, 180, 181, 7,
		9, 0, 0, 181, 44, 1, 0, 0, 0, 182, 183, 7, 10, 0, 0, 183, 46, 1, 0, 0,
		0, 184, 185, 7, 11, 0, 0, 185, 48, 1, 0, 0, 0, 186, 187, 7, 12, 0, 0, 187,
		50, 1, 0, 0, 0, 188, 189, 7, 13, 0, 0, 189, 52, 1, 0, 0, 0, 190, 191, 7,
		14, 0, 0, 191, 54, 1, 0, 0, 0, 192, 193, 7, 15, 0, 0, 193, 56, 1, 0, 0,
		0, 194, 195, 7, 16, 0, 0, 195, 58, 1, 0, 0, 0, 196, 197, 7, 17, 0, 0, 197,
		60, 1, 0, 0, 0, 198, 199, 7, 18, 0, 0, 199, 62, 1, 0, 0, 0, 200, 201, 7,
		19, 0, 0, 201, 64, 1, 0, 0, 0, 202, 203, 7, 20, 0, 0, 203, 66, 1, 0, 0,
		0, 204, 205, 7, 21, 0, 0, 205, 68, 1, 0, 0, 0, 206, 207, 7, 22, 0, 0, 207,
		70, 1, 0, 0, 0, 208, 209, 7, 23, 0, 0, 209, 72, 1, 0, 0, 0, 210, 211, 7,
		24, 0, 0, 211, 74, 1, 0, 0, 0, 212, 213, 7, 25, 0, 0, 213, 76, 1, 0, 0,
		0, 214, 215, 7, 26, 0, 0, 215, 78, 1, 0, 0, 0, 216, 217, 7, 27, 0, 0, 217,
		80, 1, 0, 0, 0, 218, 219, 7, 28, 0, 0, 219, 82, 1, 0, 0, 0, 220, 221, 7,
		29, 0, 0, 221, 84, 1, 0, 0, 0, 5, 0, 143, 145, 159, 165, 0,
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

// ClickhouseLexerInit initializes any static state used to implement ClickhouseLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewClickhouseLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClickhouseLexerInit() {
	staticData := &ClickhouseLexerLexerStaticData
	staticData.once.Do(clickhouselexerLexerInit)
}

// NewClickhouseLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewClickhouseLexer(input antlr.CharStream) *ClickhouseLexer {
	ClickhouseLexerInit()
	l := new(ClickhouseLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ClickhouseLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ClickhouseLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ClickhouseLexer tokens.
const (
	ClickhouseLexerCREATE   = 1
	ClickhouseLexerDATABASE = 2
	ClickhouseLexerIF       = 3
	ClickhouseLexerNOT      = 4
	ClickhouseLexerEXISTS   = 5
	ClickhouseLexerON       = 6
	ClickhouseLexerCLUSTER  = 7
	ClickhouseLexerENGINE   = 8
	ClickhouseLexerCOMMENT  = 9
	ClickhouseLexerID       = 10
	ClickhouseLexerEQUAL    = 11
	ClickhouseLexerLEFT_P   = 12
	ClickhouseLexerRIGHT_P  = 13
	ClickhouseLexerCOMMA    = 14
	ClickhouseLexerNUMBER   = 15
	ClickhouseLexerSTRING   = 16
	ClickhouseLexerB        = 17
	ClickhouseLexerC        = 18
	ClickhouseLexerD        = 19
	ClickhouseLexerE        = 20
	ClickhouseLexerF        = 21
	ClickhouseLexerG        = 22
	ClickhouseLexerH        = 23
	ClickhouseLexerI        = 24
	ClickhouseLexerJ        = 25
	ClickhouseLexerK        = 26
	ClickhouseLexerL        = 27
	ClickhouseLexerM        = 28
	ClickhouseLexerN        = 29
	ClickhouseLexerO        = 30
	ClickhouseLexerP        = 31
	ClickhouseLexerQ        = 32
	ClickhouseLexerR        = 33
	ClickhouseLexerS        = 34
	ClickhouseLexerT        = 35
	ClickhouseLexerU        = 36
	ClickhouseLexerV        = 37
	ClickhouseLexerW        = 38
	ClickhouseLexerX        = 39
	ClickhouseLexerY        = 40
	ClickhouseLexerZ        = 41
)