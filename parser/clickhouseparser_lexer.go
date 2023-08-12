// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type ClickHouseParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ClickHouseParserLexerLexerStaticData struct {
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

func clickhouseparserlexerLexerInit() {
	staticData := &ClickHouseParserLexerLexerStaticData
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
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.RuleNames = []string{
		"CREATE", "DATABASE", "IF", "NOT", "EXISTS", "ON", "CLUSTER", "ENGINE",
		"COMMENT", "ID", "EQUAL", "LEFT_P", "RIGHT_P", "COMMA", "NUMBER", "STRING",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 228, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 5, 9, 146, 8, 9, 10, 9, 12, 9, 149, 9, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 4, 14, 160, 8, 14, 11, 14, 12, 14,
		161, 1, 15, 1, 15, 5, 15, 166, 8, 15, 10, 15, 12, 15, 169, 9, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 167, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 0, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22,
		47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 63, 31,
		65, 32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 79, 39, 81, 40,
		83, 41, 85, 42, 1, 0, 31, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 39, 39, 2, 0, 65, 65, 97,
		97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100,
		2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103,
		2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106,
		2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109,
		2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112,
		2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115,
		2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118,
		2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121,
		2, 0, 90, 90, 122, 122, 3, 0, 9, 10, 13, 13, 32, 32, 229, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3,
		94, 1, 0, 0, 0, 5, 103, 1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 110, 1, 0, 0,
		0, 11, 117, 1, 0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 135,
		1, 0, 0, 0, 19, 143, 1, 0, 0, 0, 21, 150, 1, 0, 0, 0, 23, 152, 1, 0, 0,
		0, 25, 154, 1, 0, 0, 0, 27, 156, 1, 0, 0, 0, 29, 159, 1, 0, 0, 0, 31, 163,
		1, 0, 0, 0, 33, 172, 1, 0, 0, 0, 35, 174, 1, 0, 0, 0, 37, 176, 1, 0, 0,
		0, 39, 178, 1, 0, 0, 0, 41, 180, 1, 0, 0, 0, 43, 182, 1, 0, 0, 0, 45, 184,
		1, 0, 0, 0, 47, 186, 1, 0, 0, 0, 49, 188, 1, 0, 0, 0, 51, 190, 1, 0, 0,
		0, 53, 192, 1, 0, 0, 0, 55, 194, 1, 0, 0, 0, 57, 196, 1, 0, 0, 0, 59, 198,
		1, 0, 0, 0, 61, 200, 1, 0, 0, 0, 63, 202, 1, 0, 0, 0, 65, 204, 1, 0, 0,
		0, 67, 206, 1, 0, 0, 0, 69, 208, 1, 0, 0, 0, 71, 210, 1, 0, 0, 0, 73, 212,
		1, 0, 0, 0, 75, 214, 1, 0, 0, 0, 77, 216, 1, 0, 0, 0, 79, 218, 1, 0, 0,
		0, 81, 220, 1, 0, 0, 0, 83, 222, 1, 0, 0, 0, 85, 224, 1, 0, 0, 0, 87, 88,
		3, 37, 18, 0, 88, 89, 3, 67, 33, 0, 89, 90, 3, 41, 20, 0, 90, 91, 3, 33,
		16, 0, 91, 92, 3, 71, 35, 0, 92, 93, 3, 41, 20, 0, 93, 2, 1, 0, 0, 0, 94,
		95, 3, 39, 19, 0, 95, 96, 3, 33, 16, 0, 96, 97, 3, 71, 35, 0, 97, 98, 3,
		33, 16, 0, 98, 99, 3, 35, 17, 0, 99, 100, 3, 33, 16, 0, 100, 101, 3, 69,
		34, 0, 101, 102, 3, 41, 20, 0, 102, 4, 1, 0, 0, 0, 103, 104, 3, 49, 24,
		0, 104, 105, 3, 43, 21, 0, 105, 6, 1, 0, 0, 0, 106, 107, 3, 59, 29, 0,
		107, 108, 3, 61, 30, 0, 108, 109, 3, 71, 35, 0, 109, 8, 1, 0, 0, 0, 110,
		111, 3, 41, 20, 0, 111, 112, 3, 79, 39, 0, 112, 113, 3, 49, 24, 0, 113,
		114, 3, 69, 34, 0, 114, 115, 3, 71, 35, 0, 115, 116, 3, 69, 34, 0, 116,
		10, 1, 0, 0, 0, 117, 118, 3, 61, 30, 0, 118, 119, 3, 59, 29, 0, 119, 12,
		1, 0, 0, 0, 120, 121, 3, 37, 18, 0, 121, 122, 3, 55, 27, 0, 122, 123, 3,
		73, 36, 0, 123, 124, 3, 69, 34, 0, 124, 125, 3, 71, 35, 0, 125, 126, 3,
		41, 20, 0, 126, 127, 3, 67, 33, 0, 127, 14, 1, 0, 0, 0, 128, 129, 3, 41,
		20, 0, 129, 130, 3, 59, 29, 0, 130, 131, 3, 45, 22, 0, 131, 132, 3, 49,
		24, 0, 132, 133, 3, 59, 29, 0, 133, 134, 3, 41, 20, 0, 134, 16, 1, 0, 0,
		0, 135, 136, 3, 37, 18, 0, 136, 137, 3, 61, 30, 0, 137, 138, 3, 57, 28,
		0, 138, 139, 3, 57, 28, 0, 139, 140, 3, 41, 20, 0, 140, 141, 3, 59, 29,
		0, 141, 142, 3, 71, 35, 0, 142, 18, 1, 0, 0, 0, 143, 147, 7, 0, 0, 0, 144,
		146, 7, 1, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 20, 1, 0, 0, 0, 149, 147, 1, 0,
		0, 0, 150, 151, 5, 61, 0, 0, 151, 22, 1, 0, 0, 0, 152, 153, 5, 40, 0, 0,
		153, 24, 1, 0, 0, 0, 154, 155, 5, 41, 0, 0, 155, 26, 1, 0, 0, 0, 156, 157,
		5, 44, 0, 0, 157, 28, 1, 0, 0, 0, 158, 160, 7, 2, 0, 0, 159, 158, 1, 0,
		0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0,
		162, 30, 1, 0, 0, 0, 163, 167, 5, 39, 0, 0, 164, 166, 8, 3, 0, 0, 165,
		164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 5, 39,
		0, 0, 171, 32, 1, 0, 0, 0, 172, 173, 7, 4, 0, 0, 173, 34, 1, 0, 0, 0, 174,
		175, 7, 5, 0, 0, 175, 36, 1, 0, 0, 0, 176, 177, 7, 6, 0, 0, 177, 38, 1,
		0, 0, 0, 178, 179, 7, 7, 0, 0, 179, 40, 1, 0, 0, 0, 180, 181, 7, 8, 0,
		0, 181, 42, 1, 0, 0, 0, 182, 183, 7, 9, 0, 0, 183, 44, 1, 0, 0, 0, 184,
		185, 7, 10, 0, 0, 185, 46, 1, 0, 0, 0, 186, 187, 7, 11, 0, 0, 187, 48,
		1, 0, 0, 0, 188, 189, 7, 12, 0, 0, 189, 50, 1, 0, 0, 0, 190, 191, 7, 13,
		0, 0, 191, 52, 1, 0, 0, 0, 192, 193, 7, 14, 0, 0, 193, 54, 1, 0, 0, 0,
		194, 195, 7, 15, 0, 0, 195, 56, 1, 0, 0, 0, 196, 197, 7, 16, 0, 0, 197,
		58, 1, 0, 0, 0, 198, 199, 7, 17, 0, 0, 199, 60, 1, 0, 0, 0, 200, 201, 7,
		18, 0, 0, 201, 62, 1, 0, 0, 0, 202, 203, 7, 19, 0, 0, 203, 64, 1, 0, 0,
		0, 204, 205, 7, 20, 0, 0, 205, 66, 1, 0, 0, 0, 206, 207, 7, 21, 0, 0, 207,
		68, 1, 0, 0, 0, 208, 209, 7, 22, 0, 0, 209, 70, 1, 0, 0, 0, 210, 211, 7,
		23, 0, 0, 211, 72, 1, 0, 0, 0, 212, 213, 7, 24, 0, 0, 213, 74, 1, 0, 0,
		0, 214, 215, 7, 25, 0, 0, 215, 76, 1, 0, 0, 0, 216, 217, 7, 26, 0, 0, 217,
		78, 1, 0, 0, 0, 218, 219, 7, 27, 0, 0, 219, 80, 1, 0, 0, 0, 220, 221, 7,
		28, 0, 0, 221, 82, 1, 0, 0, 0, 222, 223, 7, 29, 0, 0, 223, 84, 1, 0, 0,
		0, 224, 225, 7, 30, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 6, 42, 0, 0,
		227, 86, 1, 0, 0, 0, 5, 0, 145, 147, 161, 167, 1, 6, 0, 0,
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

// ClickHouseParserLexerInit initializes any static state used to implement ClickHouseParserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewClickHouseParserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClickHouseParserLexerInit() {
	staticData := &ClickHouseParserLexerLexerStaticData
	staticData.once.Do(clickhouseparserlexerLexerInit)
}

// NewClickHouseParserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewClickHouseParserLexer(input antlr.CharStream) *ClickHouseParserLexer {
	ClickHouseParserLexerInit()
	l := new(ClickHouseParserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ClickHouseParserLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ClickHouseParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ClickHouseParserLexer tokens.
const (
	ClickHouseParserLexerCREATE   = 1
	ClickHouseParserLexerDATABASE = 2
	ClickHouseParserLexerIF       = 3
	ClickHouseParserLexerNOT      = 4
	ClickHouseParserLexerEXISTS   = 5
	ClickHouseParserLexerON       = 6
	ClickHouseParserLexerCLUSTER  = 7
	ClickHouseParserLexerENGINE   = 8
	ClickHouseParserLexerCOMMENT  = 9
	ClickHouseParserLexerID       = 10
	ClickHouseParserLexerEQUAL    = 11
	ClickHouseParserLexerLEFT_P   = 12
	ClickHouseParserLexerRIGHT_P  = 13
	ClickHouseParserLexerCOMMA    = 14
	ClickHouseParserLexerNUMBER   = 15
	ClickHouseParserLexerSTRING   = 16
	ClickHouseParserLexerB        = 17
	ClickHouseParserLexerC        = 18
	ClickHouseParserLexerD        = 19
	ClickHouseParserLexerE        = 20
	ClickHouseParserLexerF        = 21
	ClickHouseParserLexerG        = 22
	ClickHouseParserLexerH        = 23
	ClickHouseParserLexerI        = 24
	ClickHouseParserLexerJ        = 25
	ClickHouseParserLexerK        = 26
	ClickHouseParserLexerL        = 27
	ClickHouseParserLexerM        = 28
	ClickHouseParserLexerN        = 29
	ClickHouseParserLexerO        = 30
	ClickHouseParserLexerP        = 31
	ClickHouseParserLexerQ        = 32
	ClickHouseParserLexerR        = 33
	ClickHouseParserLexerS        = 34
	ClickHouseParserLexerT        = 35
	ClickHouseParserLexerU        = 36
	ClickHouseParserLexerV        = 37
	ClickHouseParserLexerW        = 38
	ClickHouseParserLexerX        = 39
	ClickHouseParserLexerY        = 40
	ClickHouseParserLexerZ        = 41
	ClickHouseParserLexerWS       = 42
)
