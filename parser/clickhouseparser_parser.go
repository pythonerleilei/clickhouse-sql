// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
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

type ClickHouseParserParser struct {
	*antlr.BaseParser
}

var ClickHouseParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clickhouseparserParserInit() {
	staticData := &ClickHouseParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "'='", "'('", "')'", "','", "'.'", "'+'",
	}
	staticData.SymbolicNames = []string{
		"", "CREATE", "ALTER", "DATABASE", "TABLE", "IF", "NOT", "EXISTS", "ON",
		"CLUSTER", "ENGINE", "COMMENT", "ORDER", "PARTITION", "PRIMARY", "SAMPLE",
		"TTL", "SETTINGS", "KEY", "BY", "DEFAULT", "CODEC", "INDEX", "TYPE",
		"GRANULARITY", "INTERVAL", "SECOND", "MINUTE", "HOUR", "DAY", "WEEK",
		"MONTH", "QUARTER", "YEAR", "ADD", "COLUMN", "ID", "EQUAL", "LEFT_P",
		"RIGHT_P", "COMMA", "DOT", "PLUS", "NUMBER", "STRING", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "createStatement", "createDatabaseStatement", "createTableStatement",
		"column_defs", "column_def", "index_defs", "index_def", "ttl_def", "key_values",
		"key_value", "alterStatement", "alterTableColumn", "talbeIdentifier",
		"columnType", "expression", "parameters", "parameter",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 240, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 3, 0, 39, 8, 0, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 50, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 56, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 61, 8, 2, 1, 2, 1,
		2, 3, 2, 65, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 72, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 78, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 84, 8,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 94, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 102, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		107, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 3, 3, 3, 115, 8, 3, 1,
		3, 1, 3, 3, 3, 119, 8, 3, 1, 3, 1, 3, 3, 3, 123, 8, 3, 1, 4, 1, 4, 1, 4,
		5, 4, 128, 8, 4, 10, 4, 12, 4, 131, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		137, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 143, 8, 5, 1, 5, 1, 5, 3, 5, 147,
		8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 152, 8, 6, 10, 6, 12, 6, 155, 9, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 164, 8, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 176, 8, 9, 10, 9, 12, 9,
		179, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 193, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 200, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 206, 8, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 215, 8, 15, 1, 15,
		3, 15, 218, 8, 15, 1, 15, 1, 15, 3, 15, 222, 8, 15, 1, 15, 3, 15, 225,
		8, 15, 1, 16, 1, 16, 1, 16, 5, 16, 230, 8, 16, 10, 16, 12, 16, 233, 9,
		16, 1, 17, 1, 17, 1, 17, 3, 17, 238, 8, 17, 1, 17, 0, 0, 18, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 1, 1, 0, 26,
		33, 254, 0, 38, 1, 0, 0, 0, 2, 42, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6, 66,
		1, 0, 0, 0, 8, 124, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 148, 1, 0, 0,
		0, 14, 156, 1, 0, 0, 0, 16, 165, 1, 0, 0, 0, 18, 172, 1, 0, 0, 0, 20, 180,
		1, 0, 0, 0, 22, 184, 1, 0, 0, 0, 24, 186, 1, 0, 0, 0, 26, 205, 1, 0, 0,
		0, 28, 209, 1, 0, 0, 0, 30, 224, 1, 0, 0, 0, 32, 226, 1, 0, 0, 0, 34, 237,
		1, 0, 0, 0, 36, 39, 3, 2, 1, 0, 37, 39, 3, 22, 11, 0, 38, 36, 1, 0, 0,
		0, 38, 37, 1, 0, 0, 0, 39, 1, 1, 0, 0, 0, 40, 43, 3, 4, 2, 0, 41, 43, 3,
		6, 3, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 3, 1, 0, 0, 0, 44,
		45, 5, 1, 0, 0, 45, 49, 5, 3, 0, 0, 46, 47, 5, 5, 0, 0, 47, 48, 5, 6, 0,
		0, 48, 50, 5, 7, 0, 0, 49, 46, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 55, 5, 36, 0, 0, 52, 53, 5, 8, 0, 0, 53, 54, 5, 9, 0, 0,
		54, 56, 5, 36, 0, 0, 55, 52, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 60, 1,
		0, 0, 0, 57, 58, 5, 10, 0, 0, 58, 59, 5, 37, 0, 0, 59, 61, 3, 30, 15, 0,
		60, 57, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 63, 5,
		11, 0, 0, 63, 65, 5, 44, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 5, 1, 0, 0, 0, 66, 67, 5, 1, 0, 0, 67, 71, 5, 4, 0, 0, 68, 69, 5, 5,
		0, 0, 69, 70, 5, 6, 0, 0, 70, 72, 5, 7, 0, 0, 71, 68, 1, 0, 0, 0, 71, 72,
		1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 77, 3, 26, 13, 0, 74, 75, 5, 8, 0,
		0, 75, 76, 5, 9, 0, 0, 76, 78, 5, 36, 0, 0, 77, 74, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 38, 0, 0, 80, 83, 3, 8, 4, 0,
		81, 82, 5, 40, 0, 0, 82, 84, 3, 12, 6, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 5, 39, 0, 0, 86, 87, 5, 10, 0, 0,
		87, 88, 5, 37, 0, 0, 88, 93, 5, 36, 0, 0, 89, 90, 5, 38, 0, 0, 90, 91,
		3, 32, 16, 0, 91, 92, 5, 39, 0, 0, 92, 94, 1, 0, 0, 0, 93, 89, 1, 0, 0,
		0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 5, 12, 0, 0, 96, 97,
		5, 19, 0, 0, 97, 101, 3, 30, 15, 0, 98, 99, 5, 13, 0, 0, 99, 100, 5, 19,
		0, 0, 100, 102, 3, 30, 15, 0, 101, 98, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0,
		102, 106, 1, 0, 0, 0, 103, 104, 5, 14, 0, 0, 104, 105, 5, 18, 0, 0, 105,
		107, 3, 30, 15, 0, 106, 103, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 111,
		1, 0, 0, 0, 108, 109, 5, 15, 0, 0, 109, 110, 5, 19, 0, 0, 110, 112, 3,
		30, 15, 0, 111, 108, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0,
		0, 0, 113, 115, 3, 16, 8, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 118, 1, 0, 0, 0, 116, 117, 5, 17, 0, 0, 117, 119, 3, 18, 9, 0, 118,
		116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 121,
		5, 11, 0, 0, 121, 123, 5, 44, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1,
		0, 0, 0, 123, 7, 1, 0, 0, 0, 124, 129, 3, 10, 5, 0, 125, 126, 5, 40, 0,
		0, 126, 128, 3, 10, 5, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 9, 1, 0, 0, 0, 131, 129, 1,
		0, 0, 0, 132, 133, 5, 36, 0, 0, 133, 136, 3, 30, 15, 0, 134, 135, 5, 20,
		0, 0, 135, 137, 3, 30, 15, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0,
		0, 137, 142, 1, 0, 0, 0, 138, 139, 5, 21, 0, 0, 139, 140, 5, 38, 0, 0,
		140, 141, 5, 36, 0, 0, 141, 143, 5, 39, 0, 0, 142, 138, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 145, 5, 11, 0, 0, 145, 147,
		5, 44, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 11, 1, 0,
		0, 0, 148, 153, 3, 14, 7, 0, 149, 150, 5, 40, 0, 0, 150, 152, 3, 14, 7,
		0, 151, 149, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 13, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5,
		22, 0, 0, 157, 158, 5, 36, 0, 0, 158, 159, 3, 30, 15, 0, 159, 160, 5, 23,
		0, 0, 160, 163, 3, 30, 15, 0, 161, 162, 5, 24, 0, 0, 162, 164, 5, 43, 0,
		0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 15, 1, 0, 0, 0, 165,
		166, 5, 16, 0, 0, 166, 167, 5, 36, 0, 0, 167, 168, 5, 42, 0, 0, 168, 169,
		5, 25, 0, 0, 169, 170, 5, 43, 0, 0, 170, 171, 7, 0, 0, 0, 171, 17, 1, 0,
		0, 0, 172, 177, 3, 20, 10, 0, 173, 174, 5, 40, 0, 0, 174, 176, 3, 20, 10,
		0, 175, 173, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 19, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 181, 5,
		36, 0, 0, 181, 182, 5, 37, 0, 0, 182, 183, 3, 34, 17, 0, 183, 21, 1, 0,
		0, 0, 184, 185, 3, 24, 12, 0, 185, 23, 1, 0, 0, 0, 186, 187, 5, 2, 0, 0,
		187, 188, 5, 4, 0, 0, 188, 192, 3, 26, 13, 0, 189, 190, 5, 8, 0, 0, 190,
		191, 5, 9, 0, 0, 191, 193, 5, 36, 0, 0, 192, 189, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 5, 34, 0, 0, 195, 199, 5, 35,
		0, 0, 196, 197, 5, 5, 0, 0, 197, 198, 5, 6, 0, 0, 198, 200, 5, 7, 0, 0,
		199, 196, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		202, 3, 28, 14, 0, 202, 25, 1, 0, 0, 0, 203, 204, 5, 36, 0, 0, 204, 206,
		5, 41, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1, 0,
		0, 0, 207, 208, 5, 36, 0, 0, 208, 27, 1, 0, 0, 0, 209, 210, 3, 30, 15,
		0, 210, 29, 1, 0, 0, 0, 211, 217, 5, 36, 0, 0, 212, 214, 5, 38, 0, 0, 213,
		215, 3, 32, 16, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216,
		1, 0, 0, 0, 216, 218, 5, 39, 0, 0, 217, 212, 1, 0, 0, 0, 217, 218, 1, 0,
		0, 0, 218, 225, 1, 0, 0, 0, 219, 221, 5, 38, 0, 0, 220, 222, 3, 32, 16,
		0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223,
		225, 5, 39, 0, 0, 224, 211, 1, 0, 0, 0, 224, 219, 1, 0, 0, 0, 225, 31,
		1, 0, 0, 0, 226, 231, 3, 34, 17, 0, 227, 228, 5, 40, 0, 0, 228, 230, 3,
		34, 17, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 33, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0,
		234, 238, 5, 43, 0, 0, 235, 238, 5, 44, 0, 0, 236, 238, 3, 30, 15, 0, 237,
		234, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1, 0, 0, 0, 238, 35, 1,
		0, 0, 0, 32, 38, 42, 49, 55, 60, 64, 71, 77, 83, 93, 101, 106, 111, 114,
		118, 122, 129, 136, 142, 146, 153, 163, 177, 192, 199, 205, 214, 217, 221,
		224, 231, 237,
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

// ClickHouseParserParserInit initializes any static state used to implement ClickHouseParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClickHouseParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClickHouseParserParserInit() {
	staticData := &ClickHouseParserParserStaticData
	staticData.once.Do(clickhouseparserParserInit)
}

// NewClickHouseParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClickHouseParserParser(input antlr.TokenStream) *ClickHouseParserParser {
	ClickHouseParserParserInit()
	this := new(ClickHouseParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ClickHouseParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ClickHouseParser.g4"

	return this
}

// ClickHouseParserParser tokens.
const (
	ClickHouseParserParserEOF         = antlr.TokenEOF
	ClickHouseParserParserCREATE      = 1
	ClickHouseParserParserALTER       = 2
	ClickHouseParserParserDATABASE    = 3
	ClickHouseParserParserTABLE       = 4
	ClickHouseParserParserIF          = 5
	ClickHouseParserParserNOT         = 6
	ClickHouseParserParserEXISTS      = 7
	ClickHouseParserParserON          = 8
	ClickHouseParserParserCLUSTER     = 9
	ClickHouseParserParserENGINE      = 10
	ClickHouseParserParserCOMMENT     = 11
	ClickHouseParserParserORDER       = 12
	ClickHouseParserParserPARTITION   = 13
	ClickHouseParserParserPRIMARY     = 14
	ClickHouseParserParserSAMPLE      = 15
	ClickHouseParserParserTTL         = 16
	ClickHouseParserParserSETTINGS    = 17
	ClickHouseParserParserKEY         = 18
	ClickHouseParserParserBY          = 19
	ClickHouseParserParserDEFAULT     = 20
	ClickHouseParserParserCODEC       = 21
	ClickHouseParserParserINDEX       = 22
	ClickHouseParserParserTYPE        = 23
	ClickHouseParserParserGRANULARITY = 24
	ClickHouseParserParserINTERVAL    = 25
	ClickHouseParserParserSECOND      = 26
	ClickHouseParserParserMINUTE      = 27
	ClickHouseParserParserHOUR        = 28
	ClickHouseParserParserDAY         = 29
	ClickHouseParserParserWEEK        = 30
	ClickHouseParserParserMONTH       = 31
	ClickHouseParserParserQUARTER     = 32
	ClickHouseParserParserYEAR        = 33
	ClickHouseParserParserADD         = 34
	ClickHouseParserParserCOLUMN      = 35
	ClickHouseParserParserID          = 36
	ClickHouseParserParserEQUAL       = 37
	ClickHouseParserParserLEFT_P      = 38
	ClickHouseParserParserRIGHT_P     = 39
	ClickHouseParserParserCOMMA       = 40
	ClickHouseParserParserDOT         = 41
	ClickHouseParserParserPLUS        = 42
	ClickHouseParserParserNUMBER      = 43
	ClickHouseParserParserSTRING      = 44
	ClickHouseParserParserB           = 45
	ClickHouseParserParserC           = 46
	ClickHouseParserParserD           = 47
	ClickHouseParserParserE           = 48
	ClickHouseParserParserF           = 49
	ClickHouseParserParserG           = 50
	ClickHouseParserParserH           = 51
	ClickHouseParserParserI           = 52
	ClickHouseParserParserJ           = 53
	ClickHouseParserParserK           = 54
	ClickHouseParserParserL           = 55
	ClickHouseParserParserM           = 56
	ClickHouseParserParserN           = 57
	ClickHouseParserParserO           = 58
	ClickHouseParserParserP           = 59
	ClickHouseParserParserQ           = 60
	ClickHouseParserParserR           = 61
	ClickHouseParserParserS           = 62
	ClickHouseParserParserT           = 63
	ClickHouseParserParserU           = 64
	ClickHouseParserParserV           = 65
	ClickHouseParserParserW           = 66
	ClickHouseParserParserX           = 67
	ClickHouseParserParserY           = 68
	ClickHouseParserParserZ           = 69
	ClickHouseParserParserWS          = 70
)

// ClickHouseParserParser rules.
const (
	ClickHouseParserParserRULE_statement               = 0
	ClickHouseParserParserRULE_createStatement         = 1
	ClickHouseParserParserRULE_createDatabaseStatement = 2
	ClickHouseParserParserRULE_createTableStatement    = 3
	ClickHouseParserParserRULE_column_defs             = 4
	ClickHouseParserParserRULE_column_def              = 5
	ClickHouseParserParserRULE_index_defs              = 6
	ClickHouseParserParserRULE_index_def               = 7
	ClickHouseParserParserRULE_ttl_def                 = 8
	ClickHouseParserParserRULE_key_values              = 9
	ClickHouseParserParserRULE_key_value               = 10
	ClickHouseParserParserRULE_alterStatement          = 11
	ClickHouseParserParserRULE_alterTableColumn        = 12
	ClickHouseParserParserRULE_talbeIdentifier         = 13
	ClickHouseParserParserRULE_columnType              = 14
	ClickHouseParserParserRULE_expression              = 15
	ClickHouseParserParserRULE_parameters              = 16
	ClickHouseParserParserRULE_parameter               = 17
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateStatement() ICreateStatementContext
	AlterStatement() IAlterStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CreateStatement() ICreateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateStatementContext)
}

func (s *StatementContext) AlterStatement() IAlterStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClickHouseParserParserRULE_statement)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClickHouseParserParserCREATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.CreateStatement()
		}

	case ClickHouseParserParserALTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.AlterStatement()
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

// ICreateStatementContext is an interface to support dynamic dispatch.
type ICreateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateDatabaseStatement() ICreateDatabaseStatementContext
	CreateTableStatement() ICreateTableStatementContext

	// IsCreateStatementContext differentiates from other interfaces.
	IsCreateStatementContext()
}

type CreateStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateStatementContext() *CreateStatementContext {
	var p = new(CreateStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createStatement
	return p
}

func InitEmptyCreateStatementContext(p *CreateStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createStatement
}

func (*CreateStatementContext) IsCreateStatementContext() {}

func NewCreateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateStatementContext {
	var p = new(CreateStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_createStatement

	return p
}

func (s *CreateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateStatementContext) CreateDatabaseStatement() ICreateDatabaseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateDatabaseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateDatabaseStatementContext)
}

func (s *CreateStatementContext) CreateTableStatement() ICreateTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableStatementContext)
}

func (s *CreateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterCreateStatement(s)
	}
}

func (s *CreateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitCreateStatement(s)
	}
}

func (s *CreateStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitCreateStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) CreateStatement() (localctx ICreateStatementContext) {
	localctx = NewCreateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClickHouseParserParserRULE_createStatement)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.CreateDatabaseStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.CreateTableStatement()
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

// ICreateDatabaseStatementContext is an interface to support dynamic dispatch.
type ICreateDatabaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDb_name returns the db_name token.
	GetDb_name() antlr.Token

	// GetCluster_name returns the cluster_name token.
	GetCluster_name() antlr.Token

	// GetComment returns the comment token.
	GetComment() antlr.Token

	// SetDb_name sets the db_name token.
	SetDb_name(antlr.Token)

	// SetCluster_name sets the cluster_name token.
	SetCluster_name(antlr.Token)

	// SetComment sets the comment token.
	SetComment(antlr.Token)

	// GetEngine_name returns the engine_name rule contexts.
	GetEngine_name() IExpressionContext

	// SetEngine_name sets the engine_name rule contexts.
	SetEngine_name(IExpressionContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	ON() antlr.TerminalNode
	CLUSTER() antlr.TerminalNode
	ENGINE() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	Expression() IExpressionContext
	STRING() antlr.TerminalNode

	// IsCreateDatabaseStatementContext differentiates from other interfaces.
	IsCreateDatabaseStatementContext()
}

type CreateDatabaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	db_name      antlr.Token
	cluster_name antlr.Token
	engine_name  IExpressionContext
	comment      antlr.Token
}

func NewEmptyCreateDatabaseStatementContext() *CreateDatabaseStatementContext {
	var p = new(CreateDatabaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createDatabaseStatement
	return p
}

func InitEmptyCreateDatabaseStatementContext(p *CreateDatabaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createDatabaseStatement
}

func (*CreateDatabaseStatementContext) IsCreateDatabaseStatementContext() {}

func NewCreateDatabaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDatabaseStatementContext {
	var p = new(CreateDatabaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_createDatabaseStatement

	return p
}

func (s *CreateDatabaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDatabaseStatementContext) GetDb_name() antlr.Token { return s.db_name }

func (s *CreateDatabaseStatementContext) GetCluster_name() antlr.Token { return s.cluster_name }

func (s *CreateDatabaseStatementContext) GetComment() antlr.Token { return s.comment }

func (s *CreateDatabaseStatementContext) SetDb_name(v antlr.Token) { s.db_name = v }

func (s *CreateDatabaseStatementContext) SetCluster_name(v antlr.Token) { s.cluster_name = v }

func (s *CreateDatabaseStatementContext) SetComment(v antlr.Token) { s.comment = v }

func (s *CreateDatabaseStatementContext) GetEngine_name() IExpressionContext { return s.engine_name }

func (s *CreateDatabaseStatementContext) SetEngine_name(v IExpressionContext) { s.engine_name = v }

func (s *CreateDatabaseStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCREATE, 0)
}

func (s *CreateDatabaseStatementContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserDATABASE, 0)
}

func (s *CreateDatabaseStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserID)
}

func (s *CreateDatabaseStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, i)
}

func (s *CreateDatabaseStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserIF, 0)
}

func (s *CreateDatabaseStatementContext) NOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNOT, 0)
}

func (s *CreateDatabaseStatementContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEXISTS, 0)
}

func (s *CreateDatabaseStatementContext) ON() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserON, 0)
}

func (s *CreateDatabaseStatementContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCLUSTER, 0)
}

func (s *CreateDatabaseStatementContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserENGINE, 0)
}

func (s *CreateDatabaseStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEQUAL, 0)
}

func (s *CreateDatabaseStatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMENT, 0)
}

func (s *CreateDatabaseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CreateDatabaseStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSTRING, 0)
}

func (s *CreateDatabaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDatabaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDatabaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterCreateDatabaseStatement(s)
	}
}

func (s *CreateDatabaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitCreateDatabaseStatement(s)
	}
}

func (s *CreateDatabaseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitCreateDatabaseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) CreateDatabaseStatement() (localctx ICreateDatabaseStatementContext) {
	localctx = NewCreateDatabaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ClickHouseParserParserRULE_createDatabaseStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(ClickHouseParserParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(ClickHouseParserParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(46)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(51)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*CreateDatabaseStatementContext).db_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(52)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*CreateDatabaseStatementContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserENGINE {
		{
			p.SetState(57)
			p.Match(ClickHouseParserParserENGINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(ClickHouseParserParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _x = p.Expression()

			localctx.(*CreateDatabaseStatementContext).engine_name = _x
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMENT {
		{
			p.SetState(62)
			p.Match(ClickHouseParserParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)

			var _m = p.Match(ClickHouseParserParserSTRING)

			localctx.(*CreateDatabaseStatementContext).comment = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ICreateTableStatementContext is an interface to support dynamic dispatch.
type ICreateTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCluster_name returns the cluster_name token.
	GetCluster_name() antlr.Token

	// GetEngine returns the engine token.
	GetEngine() antlr.Token

	// SetCluster_name sets the cluster_name token.
	SetCluster_name(antlr.Token)

	// SetEngine sets the engine token.
	SetEngine(antlr.Token)

	// GetTable_name returns the table_name rule contexts.
	GetTable_name() ITalbeIdentifierContext

	// GetEngine_paras returns the engine_paras rule contexts.
	GetEngine_paras() IParametersContext

	// GetOder_by returns the oder_by rule contexts.
	GetOder_by() IExpressionContext

	// GetPartition_by returns the partition_by rule contexts.
	GetPartition_by() IExpressionContext

	// GetPrimary_key returns the primary_key rule contexts.
	GetPrimary_key() IExpressionContext

	// GetSample_by returns the sample_by rule contexts.
	GetSample_by() IExpressionContext

	// GetKvs returns the kvs rule contexts.
	GetKvs() IKey_valuesContext

	// SetTable_name sets the table_name rule contexts.
	SetTable_name(ITalbeIdentifierContext)

	// SetEngine_paras sets the engine_paras rule contexts.
	SetEngine_paras(IParametersContext)

	// SetOder_by sets the oder_by rule contexts.
	SetOder_by(IExpressionContext)

	// SetPartition_by sets the partition_by rule contexts.
	SetPartition_by(IExpressionContext)

	// SetPrimary_key sets the primary_key rule contexts.
	SetPrimary_key(IExpressionContext)

	// SetSample_by sets the sample_by rule contexts.
	SetSample_by(IExpressionContext)

	// SetKvs sets the kvs rule contexts.
	SetKvs(IKey_valuesContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	AllLEFT_P() []antlr.TerminalNode
	LEFT_P(i int) antlr.TerminalNode
	Column_defs() IColumn_defsContext
	AllRIGHT_P() []antlr.TerminalNode
	RIGHT_P(i int) antlr.TerminalNode
	ENGINE() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	ORDER() antlr.TerminalNode
	AllBY() []antlr.TerminalNode
	BY(i int) antlr.TerminalNode
	TalbeIdentifier() ITalbeIdentifierContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	ON() antlr.TerminalNode
	CLUSTER() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Index_defs() IIndex_defsContext
	PARTITION() antlr.TerminalNode
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode
	SAMPLE() antlr.TerminalNode
	Ttl_def() ITtl_defContext
	SETTINGS() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Parameters() IParametersContext
	Key_values() IKey_valuesContext

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	table_name   ITalbeIdentifierContext
	cluster_name antlr.Token
	engine       antlr.Token
	engine_paras IParametersContext
	oder_by      IExpressionContext
	partition_by IExpressionContext
	primary_key  IExpressionContext
	sample_by    IExpressionContext
	kvs          IKey_valuesContext
}

func NewEmptyCreateTableStatementContext() *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createTableStatement
	return p
}

func InitEmptyCreateTableStatementContext(p *CreateTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_createTableStatement
}

func (*CreateTableStatementContext) IsCreateTableStatementContext() {}

func NewCreateTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_createTableStatement

	return p
}

func (s *CreateTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableStatementContext) GetCluster_name() antlr.Token { return s.cluster_name }

func (s *CreateTableStatementContext) GetEngine() antlr.Token { return s.engine }

func (s *CreateTableStatementContext) SetCluster_name(v antlr.Token) { s.cluster_name = v }

func (s *CreateTableStatementContext) SetEngine(v antlr.Token) { s.engine = v }

func (s *CreateTableStatementContext) GetTable_name() ITalbeIdentifierContext { return s.table_name }

func (s *CreateTableStatementContext) GetEngine_paras() IParametersContext { return s.engine_paras }

func (s *CreateTableStatementContext) GetOder_by() IExpressionContext { return s.oder_by }

func (s *CreateTableStatementContext) GetPartition_by() IExpressionContext { return s.partition_by }

func (s *CreateTableStatementContext) GetPrimary_key() IExpressionContext { return s.primary_key }

func (s *CreateTableStatementContext) GetSample_by() IExpressionContext { return s.sample_by }

func (s *CreateTableStatementContext) GetKvs() IKey_valuesContext { return s.kvs }

func (s *CreateTableStatementContext) SetTable_name(v ITalbeIdentifierContext) { s.table_name = v }

func (s *CreateTableStatementContext) SetEngine_paras(v IParametersContext) { s.engine_paras = v }

func (s *CreateTableStatementContext) SetOder_by(v IExpressionContext) { s.oder_by = v }

func (s *CreateTableStatementContext) SetPartition_by(v IExpressionContext) { s.partition_by = v }

func (s *CreateTableStatementContext) SetPrimary_key(v IExpressionContext) { s.primary_key = v }

func (s *CreateTableStatementContext) SetSample_by(v IExpressionContext) { s.sample_by = v }

func (s *CreateTableStatementContext) SetKvs(v IKey_valuesContext) { s.kvs = v }

func (s *CreateTableStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCREATE, 0)
}

func (s *CreateTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserTABLE, 0)
}

func (s *CreateTableStatementContext) AllLEFT_P() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserLEFT_P)
}

func (s *CreateTableStatementContext) LEFT_P(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserLEFT_P, i)
}

func (s *CreateTableStatementContext) Column_defs() IColumn_defsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_defsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_defsContext)
}

func (s *CreateTableStatementContext) AllRIGHT_P() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserRIGHT_P)
}

func (s *CreateTableStatementContext) RIGHT_P(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserRIGHT_P, i)
}

func (s *CreateTableStatementContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserENGINE, 0)
}

func (s *CreateTableStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEQUAL, 0)
}

func (s *CreateTableStatementContext) ORDER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserORDER, 0)
}

func (s *CreateTableStatementContext) AllBY() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserBY)
}

func (s *CreateTableStatementContext) BY(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserBY, i)
}

func (s *CreateTableStatementContext) TalbeIdentifier() ITalbeIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITalbeIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITalbeIdentifierContext)
}

func (s *CreateTableStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserID)
}

func (s *CreateTableStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, i)
}

func (s *CreateTableStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CreateTableStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *CreateTableStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserIF, 0)
}

func (s *CreateTableStatementContext) NOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNOT, 0)
}

func (s *CreateTableStatementContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEXISTS, 0)
}

func (s *CreateTableStatementContext) ON() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserON, 0)
}

func (s *CreateTableStatementContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCLUSTER, 0)
}

func (s *CreateTableStatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, 0)
}

func (s *CreateTableStatementContext) Index_defs() IIndex_defsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_defsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_defsContext)
}

func (s *CreateTableStatementContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserPARTITION, 0)
}

func (s *CreateTableStatementContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserPRIMARY, 0)
}

func (s *CreateTableStatementContext) KEY() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserKEY, 0)
}

func (s *CreateTableStatementContext) SAMPLE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSAMPLE, 0)
}

func (s *CreateTableStatementContext) Ttl_def() ITtl_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITtl_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITtl_defContext)
}

func (s *CreateTableStatementContext) SETTINGS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSETTINGS, 0)
}

func (s *CreateTableStatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMENT, 0)
}

func (s *CreateTableStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSTRING, 0)
}

func (s *CreateTableStatementContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *CreateTableStatementContext) Key_values() IKey_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKey_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKey_valuesContext)
}

func (s *CreateTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterCreateTableStatement(s)
	}
}

func (s *CreateTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitCreateTableStatement(s)
	}
}

func (s *CreateTableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitCreateTableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) CreateTableStatement() (localctx ICreateTableStatementContext) {
	localctx = NewCreateTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ClickHouseParserParserRULE_createTableStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ClickHouseParserParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(ClickHouseParserParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(68)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(73)

		var _x = p.TalbeIdentifier()

		localctx.(*CreateTableStatementContext).table_name = _x
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(74)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*CreateTableStatementContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(79)
		p.Match(ClickHouseParserParserLEFT_P)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Column_defs()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMA {
		{
			p.SetState(81)
			p.Match(ClickHouseParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Index_defs()
		}

	}
	{
		p.SetState(85)
		p.Match(ClickHouseParserParserRIGHT_P)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(ClickHouseParserParserENGINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(ClickHouseParserParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*CreateTableStatementContext).engine = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserLEFT_P {
		{
			p.SetState(89)
			p.Match(ClickHouseParserParserLEFT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)

			var _x = p.Parameters()

			localctx.(*CreateTableStatementContext).engine_paras = _x
		}
		{
			p.SetState(91)
			p.Match(ClickHouseParserParserRIGHT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(95)
		p.Match(ClickHouseParserParserORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(ClickHouseParserParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)

		var _x = p.Expression()

		localctx.(*CreateTableStatementContext).oder_by = _x
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserPARTITION {
		{
			p.SetState(98)
			p.Match(ClickHouseParserParserPARTITION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(ClickHouseParserParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)

			var _x = p.Expression()

			localctx.(*CreateTableStatementContext).partition_by = _x
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserPRIMARY {
		{
			p.SetState(103)
			p.Match(ClickHouseParserParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(ClickHouseParserParserKEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _x = p.Expression()

			localctx.(*CreateTableStatementContext).primary_key = _x
		}

	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserSAMPLE {
		{
			p.SetState(108)
			p.Match(ClickHouseParserParserSAMPLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(ClickHouseParserParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)

			var _x = p.Expression()

			localctx.(*CreateTableStatementContext).sample_by = _x
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserTTL {
		{
			p.SetState(113)
			p.Ttl_def()
		}

	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserSETTINGS {
		{
			p.SetState(116)
			p.Match(ClickHouseParserParserSETTINGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)

			var _x = p.Key_values()

			localctx.(*CreateTableStatementContext).kvs = _x
		}

	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMENT {
		{
			p.SetState(120)
			p.Match(ClickHouseParserParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(ClickHouseParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IColumn_defsContext is an interface to support dynamic dispatch.
type IColumn_defsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_def() []IColumn_defContext
	Column_def(i int) IColumn_defContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_defsContext differentiates from other interfaces.
	IsColumn_defsContext()
}

type Column_defsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_defsContext() *Column_defsContext {
	var p = new(Column_defsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_column_defs
	return p
}

func InitEmptyColumn_defsContext(p *Column_defsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_column_defs
}

func (*Column_defsContext) IsColumn_defsContext() {}

func NewColumn_defsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_defsContext {
	var p = new(Column_defsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_column_defs

	return p
}

func (s *Column_defsContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_defsContext) AllColumn_def() []IColumn_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_defContext); ok {
			len++
		}
	}

	tst := make([]IColumn_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_defContext); ok {
			tst[i] = t.(IColumn_defContext)
			i++
		}
	}

	return tst
}

func (s *Column_defsContext) Column_def(i int) IColumn_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_defContext); ok {
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

	return t.(IColumn_defContext)
}

func (s *Column_defsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserCOMMA)
}

func (s *Column_defsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, i)
}

func (s *Column_defsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_defsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_defsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterColumn_defs(s)
	}
}

func (s *Column_defsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitColumn_defs(s)
	}
}

func (s *Column_defsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitColumn_defs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Column_defs() (localctx IColumn_defsContext) {
	localctx = NewColumn_defsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ClickHouseParserParserRULE_column_defs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Column_def()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(125)
				p.Match(ClickHouseParserParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(126)
				p.Column_def()
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumn_defContext is an interface to support dynamic dispatch.
type IColumn_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetCodec returns the codec token.
	GetCodec() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetCodec sets the codec token.
	SetCodec(antlr.Token)

	// GetData_type returns the data_type rule contexts.
	GetData_type() IExpressionContext

	// GetDefault_value returns the default_value rule contexts.
	GetDefault_value() IExpressionContext

	// SetData_type sets the data_type rule contexts.
	SetData_type(IExpressionContext)

	// SetDefault_value sets the default_value rule contexts.
	SetDefault_value(IExpressionContext)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	DEFAULT() antlr.TerminalNode
	CODEC() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsColumn_defContext differentiates from other interfaces.
	IsColumn_defContext()
}

type Column_defContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	name          antlr.Token
	data_type     IExpressionContext
	default_value IExpressionContext
	codec         antlr.Token
}

func NewEmptyColumn_defContext() *Column_defContext {
	var p = new(Column_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_column_def
	return p
}

func InitEmptyColumn_defContext(p *Column_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_column_def
}

func (*Column_defContext) IsColumn_defContext() {}

func NewColumn_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_defContext {
	var p = new(Column_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_column_def

	return p
}

func (s *Column_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_defContext) GetName() antlr.Token { return s.name }

func (s *Column_defContext) GetCodec() antlr.Token { return s.codec }

func (s *Column_defContext) SetName(v antlr.Token) { s.name = v }

func (s *Column_defContext) SetCodec(v antlr.Token) { s.codec = v }

func (s *Column_defContext) GetData_type() IExpressionContext { return s.data_type }

func (s *Column_defContext) GetDefault_value() IExpressionContext { return s.default_value }

func (s *Column_defContext) SetData_type(v IExpressionContext) { s.data_type = v }

func (s *Column_defContext) SetDefault_value(v IExpressionContext) { s.default_value = v }

func (s *Column_defContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserID)
}

func (s *Column_defContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, i)
}

func (s *Column_defContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Column_defContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Column_defContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserDEFAULT, 0)
}

func (s *Column_defContext) CODEC() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCODEC, 0)
}

func (s *Column_defContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserLEFT_P, 0)
}

func (s *Column_defContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserRIGHT_P, 0)
}

func (s *Column_defContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMENT, 0)
}

func (s *Column_defContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSTRING, 0)
}

func (s *Column_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterColumn_def(s)
	}
}

func (s *Column_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitColumn_def(s)
	}
}

func (s *Column_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitColumn_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Column_def() (localctx IColumn_defContext) {
	localctx = NewColumn_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ClickHouseParserParserRULE_column_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*Column_defContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)

		var _x = p.Expression()

		localctx.(*Column_defContext).data_type = _x
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserDEFAULT {
		{
			p.SetState(134)
			p.Match(ClickHouseParserParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _x = p.Expression()

			localctx.(*Column_defContext).default_value = _x
		}

	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCODEC {
		{
			p.SetState(138)
			p.Match(ClickHouseParserParserCODEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(ClickHouseParserParserLEFT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*Column_defContext).codec = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(ClickHouseParserParserRIGHT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMENT {
		{
			p.SetState(144)
			p.Match(ClickHouseParserParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(ClickHouseParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IIndex_defsContext is an interface to support dynamic dispatch.
type IIndex_defsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIndex_def() []IIndex_defContext
	Index_def(i int) IIndex_defContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIndex_defsContext differentiates from other interfaces.
	IsIndex_defsContext()
}

type Index_defsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_defsContext() *Index_defsContext {
	var p = new(Index_defsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_index_defs
	return p
}

func InitEmptyIndex_defsContext(p *Index_defsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_index_defs
}

func (*Index_defsContext) IsIndex_defsContext() {}

func NewIndex_defsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_defsContext {
	var p = new(Index_defsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_index_defs

	return p
}

func (s *Index_defsContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_defsContext) AllIndex_def() []IIndex_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_defContext); ok {
			len++
		}
	}

	tst := make([]IIndex_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_defContext); ok {
			tst[i] = t.(IIndex_defContext)
			i++
		}
	}

	return tst
}

func (s *Index_defsContext) Index_def(i int) IIndex_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_defContext); ok {
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

	return t.(IIndex_defContext)
}

func (s *Index_defsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserCOMMA)
}

func (s *Index_defsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, i)
}

func (s *Index_defsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_defsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_defsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterIndex_defs(s)
	}
}

func (s *Index_defsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitIndex_defs(s)
	}
}

func (s *Index_defsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitIndex_defs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Index_defs() (localctx IIndex_defsContext) {
	localctx = NewIndex_defsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ClickHouseParserParserRULE_index_defs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Index_def()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ClickHouseParserParserCOMMA {
		{
			p.SetState(149)
			p.Match(ClickHouseParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Index_def()
		}

		p.SetState(155)
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

// IIndex_defContext is an interface to support dynamic dispatch.
type IIndex_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEval returns the eval rule contexts.
	GetEval() IExpressionContext

	// GetType_exp returns the type_exp rule contexts.
	GetType_exp() IExpressionContext

	// SetEval sets the eval rule contexts.
	SetEval(IExpressionContext)

	// SetType_exp sets the type_exp rule contexts.
	SetType_exp(IExpressionContext)

	// Getter signatures
	INDEX() antlr.TerminalNode
	ID() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	GRANULARITY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsIndex_defContext differentiates from other interfaces.
	IsIndex_defContext()
}

type Index_defContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	eval     IExpressionContext
	type_exp IExpressionContext
}

func NewEmptyIndex_defContext() *Index_defContext {
	var p = new(Index_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_index_def
	return p
}

func InitEmptyIndex_defContext(p *Index_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_index_def
}

func (*Index_defContext) IsIndex_defContext() {}

func NewIndex_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_defContext {
	var p = new(Index_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_index_def

	return p
}

func (s *Index_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_defContext) GetEval() IExpressionContext { return s.eval }

func (s *Index_defContext) GetType_exp() IExpressionContext { return s.type_exp }

func (s *Index_defContext) SetEval(v IExpressionContext) { s.eval = v }

func (s *Index_defContext) SetType_exp(v IExpressionContext) { s.type_exp = v }

func (s *Index_defContext) INDEX() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserINDEX, 0)
}

func (s *Index_defContext) ID() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, 0)
}

func (s *Index_defContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserTYPE, 0)
}

func (s *Index_defContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Index_defContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Index_defContext) GRANULARITY() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserGRANULARITY, 0)
}

func (s *Index_defContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNUMBER, 0)
}

func (s *Index_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterIndex_def(s)
	}
}

func (s *Index_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitIndex_def(s)
	}
}

func (s *Index_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitIndex_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Index_def() (localctx IIndex_defContext) {
	localctx = NewIndex_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ClickHouseParserParserRULE_index_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(ClickHouseParserParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(ClickHouseParserParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)

		var _x = p.Expression()

		localctx.(*Index_defContext).eval = _x
	}
	{
		p.SetState(159)
		p.Match(ClickHouseParserParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)

		var _x = p.Expression()

		localctx.(*Index_defContext).type_exp = _x
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserGRANULARITY {
		{
			p.SetState(161)
			p.Match(ClickHouseParserParserGRANULARITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(ClickHouseParserParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ITtl_defContext is an interface to support dynamic dispatch.
type ITtl_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetColumn returns the column token.
	GetColumn() antlr.Token

	// GetUnit returns the unit token.
	GetUnit() antlr.Token

	// SetColumn sets the column token.
	SetColumn(antlr.Token)

	// SetUnit sets the unit token.
	SetUnit(antlr.Token)

	// Getter signatures
	TTL() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	INTERVAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	ID() antlr.TerminalNode
	SECOND() antlr.TerminalNode
	MINUTE() antlr.TerminalNode
	HOUR() antlr.TerminalNode
	DAY() antlr.TerminalNode
	WEEK() antlr.TerminalNode
	MONTH() antlr.TerminalNode
	QUARTER() antlr.TerminalNode
	YEAR() antlr.TerminalNode

	// IsTtl_defContext differentiates from other interfaces.
	IsTtl_defContext()
}

type Ttl_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	column antlr.Token
	unit   antlr.Token
}

func NewEmptyTtl_defContext() *Ttl_defContext {
	var p = new(Ttl_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_ttl_def
	return p
}

func InitEmptyTtl_defContext(p *Ttl_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_ttl_def
}

func (*Ttl_defContext) IsTtl_defContext() {}

func NewTtl_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ttl_defContext {
	var p = new(Ttl_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_ttl_def

	return p
}

func (s *Ttl_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Ttl_defContext) GetColumn() antlr.Token { return s.column }

func (s *Ttl_defContext) GetUnit() antlr.Token { return s.unit }

func (s *Ttl_defContext) SetColumn(v antlr.Token) { s.column = v }

func (s *Ttl_defContext) SetUnit(v antlr.Token) { s.unit = v }

func (s *Ttl_defContext) TTL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserTTL, 0)
}

func (s *Ttl_defContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserPLUS, 0)
}

func (s *Ttl_defContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserINTERVAL, 0)
}

func (s *Ttl_defContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNUMBER, 0)
}

func (s *Ttl_defContext) ID() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, 0)
}

func (s *Ttl_defContext) SECOND() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSECOND, 0)
}

func (s *Ttl_defContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserMINUTE, 0)
}

func (s *Ttl_defContext) HOUR() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserHOUR, 0)
}

func (s *Ttl_defContext) DAY() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserDAY, 0)
}

func (s *Ttl_defContext) WEEK() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserWEEK, 0)
}

func (s *Ttl_defContext) MONTH() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserMONTH, 0)
}

func (s *Ttl_defContext) QUARTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserQUARTER, 0)
}

func (s *Ttl_defContext) YEAR() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserYEAR, 0)
}

func (s *Ttl_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ttl_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ttl_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterTtl_def(s)
	}
}

func (s *Ttl_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitTtl_def(s)
	}
}

func (s *Ttl_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitTtl_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Ttl_def() (localctx ITtl_defContext) {
	localctx = NewTtl_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ClickHouseParserParserRULE_ttl_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(ClickHouseParserParserTTL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*Ttl_defContext).column = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(ClickHouseParserParserPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(ClickHouseParserParserINTERVAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(ClickHouseParserParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Ttl_defContext).unit = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17112760320) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Ttl_defContext).unit = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IKey_valuesContext is an interface to support dynamic dispatch.
type IKey_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKey_value() []IKey_valueContext
	Key_value(i int) IKey_valueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsKey_valuesContext differentiates from other interfaces.
	IsKey_valuesContext()
}

type Key_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_valuesContext() *Key_valuesContext {
	var p = new(Key_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_key_values
	return p
}

func InitEmptyKey_valuesContext(p *Key_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_key_values
}

func (*Key_valuesContext) IsKey_valuesContext() {}

func NewKey_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_valuesContext {
	var p = new(Key_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_key_values

	return p
}

func (s *Key_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_valuesContext) AllKey_value() []IKey_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKey_valueContext); ok {
			len++
		}
	}

	tst := make([]IKey_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKey_valueContext); ok {
			tst[i] = t.(IKey_valueContext)
			i++
		}
	}

	return tst
}

func (s *Key_valuesContext) Key_value(i int) IKey_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKey_valueContext); ok {
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

	return t.(IKey_valueContext)
}

func (s *Key_valuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserCOMMA)
}

func (s *Key_valuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, i)
}

func (s *Key_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterKey_values(s)
	}
}

func (s *Key_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitKey_values(s)
	}
}

func (s *Key_valuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitKey_values(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Key_values() (localctx IKey_valuesContext) {
	localctx = NewKey_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ClickHouseParserParserRULE_key_values)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Key_value()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ClickHouseParserParserCOMMA {
		{
			p.SetState(173)
			p.Match(ClickHouseParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Key_value()
		}

		p.SetState(179)
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

// IKey_valueContext is an interface to support dynamic dispatch.
type IKey_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Parameter() IParameterContext

	// IsKey_valueContext differentiates from other interfaces.
	IsKey_valueContext()
}

type Key_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_valueContext() *Key_valueContext {
	var p = new(Key_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_key_value
	return p
}

func InitEmptyKey_valueContext(p *Key_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_key_value
}

func (*Key_valueContext) IsKey_valueContext() {}

func NewKey_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_valueContext {
	var p = new(Key_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_key_value

	return p
}

func (s *Key_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_valueContext) ID() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, 0)
}

func (s *Key_valueContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEQUAL, 0)
}

func (s *Key_valueContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Key_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterKey_value(s)
	}
}

func (s *Key_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitKey_value(s)
	}
}

func (s *Key_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitKey_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Key_value() (localctx IKey_valueContext) {
	localctx = NewKey_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ClickHouseParserParserRULE_key_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(ClickHouseParserParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(ClickHouseParserParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Parameter()
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

// IAlterStatementContext is an interface to support dynamic dispatch.
type IAlterStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AlterTableColumn() IAlterTableColumnContext

	// IsAlterStatementContext differentiates from other interfaces.
	IsAlterStatementContext()
}

type AlterStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterStatementContext() *AlterStatementContext {
	var p = new(AlterStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_alterStatement
	return p
}

func InitEmptyAlterStatementContext(p *AlterStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_alterStatement
}

func (*AlterStatementContext) IsAlterStatementContext() {}

func NewAlterStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterStatementContext {
	var p = new(AlterStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_alterStatement

	return p
}

func (s *AlterStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterStatementContext) AlterTableColumn() IAlterTableColumnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterTableColumnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterTableColumnContext)
}

func (s *AlterStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterAlterStatement(s)
	}
}

func (s *AlterStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitAlterStatement(s)
	}
}

func (s *AlterStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitAlterStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) AlterStatement() (localctx IAlterStatementContext) {
	localctx = NewAlterStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ClickHouseParserParserRULE_alterStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.AlterTableColumn()
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

// IAlterTableColumnContext is an interface to support dynamic dispatch.
type IAlterTableColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCluster_name returns the cluster_name token.
	GetCluster_name() antlr.Token

	// SetCluster_name sets the cluster_name token.
	SetCluster_name(antlr.Token)

	// GetTable_name returns the table_name rule contexts.
	GetTable_name() ITalbeIdentifierContext

	// GetColumn_type returns the column_type rule contexts.
	GetColumn_type() IColumnTypeContext

	// SetTable_name sets the table_name rule contexts.
	SetTable_name(ITalbeIdentifierContext)

	// SetColumn_type sets the column_type rule contexts.
	SetColumn_type(IColumnTypeContext)

	// Getter signatures
	ALTER() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	ADD() antlr.TerminalNode
	COLUMN() antlr.TerminalNode
	TalbeIdentifier() ITalbeIdentifierContext
	ColumnType() IColumnTypeContext
	ON() antlr.TerminalNode
	CLUSTER() antlr.TerminalNode
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsAlterTableColumnContext differentiates from other interfaces.
	IsAlterTableColumnContext()
}

type AlterTableColumnContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	table_name   ITalbeIdentifierContext
	cluster_name antlr.Token
	column_type  IColumnTypeContext
}

func NewEmptyAlterTableColumnContext() *AlterTableColumnContext {
	var p = new(AlterTableColumnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_alterTableColumn
	return p
}

func InitEmptyAlterTableColumnContext(p *AlterTableColumnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_alterTableColumn
}

func (*AlterTableColumnContext) IsAlterTableColumnContext() {}

func NewAlterTableColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableColumnContext {
	var p = new(AlterTableColumnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_alterTableColumn

	return p
}

func (s *AlterTableColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableColumnContext) GetCluster_name() antlr.Token { return s.cluster_name }

func (s *AlterTableColumnContext) SetCluster_name(v antlr.Token) { s.cluster_name = v }

func (s *AlterTableColumnContext) GetTable_name() ITalbeIdentifierContext { return s.table_name }

func (s *AlterTableColumnContext) GetColumn_type() IColumnTypeContext { return s.column_type }

func (s *AlterTableColumnContext) SetTable_name(v ITalbeIdentifierContext) { s.table_name = v }

func (s *AlterTableColumnContext) SetColumn_type(v IColumnTypeContext) { s.column_type = v }

func (s *AlterTableColumnContext) ALTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserALTER, 0)
}

func (s *AlterTableColumnContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserTABLE, 0)
}

func (s *AlterTableColumnContext) ADD() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserADD, 0)
}

func (s *AlterTableColumnContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOLUMN, 0)
}

func (s *AlterTableColumnContext) TalbeIdentifier() ITalbeIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITalbeIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITalbeIdentifierContext)
}

func (s *AlterTableColumnContext) ColumnType() IColumnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnTypeContext)
}

func (s *AlterTableColumnContext) ON() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserON, 0)
}

func (s *AlterTableColumnContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCLUSTER, 0)
}

func (s *AlterTableColumnContext) IF() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserIF, 0)
}

func (s *AlterTableColumnContext) NOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNOT, 0)
}

func (s *AlterTableColumnContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEXISTS, 0)
}

func (s *AlterTableColumnContext) ID() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, 0)
}

func (s *AlterTableColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterTableColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterAlterTableColumn(s)
	}
}

func (s *AlterTableColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitAlterTableColumn(s)
	}
}

func (s *AlterTableColumnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitAlterTableColumn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) AlterTableColumn() (localctx IAlterTableColumnContext) {
	localctx = NewAlterTableColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ClickHouseParserParserRULE_alterTableColumn)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(ClickHouseParserParserALTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(ClickHouseParserParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)

		var _x = p.TalbeIdentifier()

		localctx.(*AlterTableColumnContext).table_name = _x
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(189)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*AlterTableColumnContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(194)
		p.Match(ClickHouseParserParserADD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(ClickHouseParserParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(196)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(201)

		var _x = p.ColumnType()

		localctx.(*AlterTableColumnContext).column_type = _x
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

// ITalbeIdentifierContext is an interface to support dynamic dispatch.
type ITalbeIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsTalbeIdentifierContext differentiates from other interfaces.
	IsTalbeIdentifierContext()
}

type TalbeIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTalbeIdentifierContext() *TalbeIdentifierContext {
	var p = new(TalbeIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_talbeIdentifier
	return p
}

func InitEmptyTalbeIdentifierContext(p *TalbeIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_talbeIdentifier
}

func (*TalbeIdentifierContext) IsTalbeIdentifierContext() {}

func NewTalbeIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TalbeIdentifierContext {
	var p = new(TalbeIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_talbeIdentifier

	return p
}

func (s *TalbeIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TalbeIdentifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserID)
}

func (s *TalbeIdentifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, i)
}

func (s *TalbeIdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserDOT, 0)
}

func (s *TalbeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TalbeIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TalbeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterTalbeIdentifier(s)
	}
}

func (s *TalbeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitTalbeIdentifier(s)
	}
}

func (s *TalbeIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitTalbeIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) TalbeIdentifier() (localctx ITalbeIdentifierContext) {
	localctx = NewTalbeIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ClickHouseParserParserRULE_talbeIdentifier)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(203)
			p.Match(ClickHouseParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(ClickHouseParserParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(207)
		p.Match(ClickHouseParserParserID)
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

// IColumnTypeContext is an interface to support dynamic dispatch.
type IColumnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsColumnTypeContext differentiates from other interfaces.
	IsColumnTypeContext()
}

type ColumnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnTypeContext() *ColumnTypeContext {
	var p = new(ColumnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_columnType
	return p
}

func InitEmptyColumnTypeContext(p *ColumnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_columnType
}

func (*ColumnTypeContext) IsColumnTypeContext() {}

func NewColumnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnTypeContext {
	var p = new(ColumnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_columnType

	return p
}

func (s *ColumnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnTypeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ColumnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterColumnType(s)
	}
}

func (s *ColumnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitColumnType(s)
	}
}

func (s *ColumnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitColumnType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) ColumnType() (localctx IColumnTypeContext) {
	localctx = NewColumnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ClickHouseParserParserRULE_columnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Expression()
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	Parameters() IParametersContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserID, 0)
}

func (s *ExpressionContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserLEFT_P, 0)
}

func (s *ExpressionContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserRIGHT_P, 0)
}

func (s *ExpressionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ClickHouseParserParserRULE_expression)
	var _la int

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClickHouseParserParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(ClickHouseParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ClickHouseParserParserLEFT_P {
			{
				p.SetState(212)
				p.Match(ClickHouseParserParserLEFT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26731876450304) != 0 {
				{
					p.SetState(213)
					p.Parameters()
				}

			}
			{
				p.SetState(216)
				p.Match(ClickHouseParserParserRIGHT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ClickHouseParserParserLEFT_P:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(ClickHouseParserParserLEFT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26731876450304) != 0 {
			{
				p.SetState(220)
				p.Parameters()
			}

		}
		{
			p.SetState(223)
			p.Match(ClickHouseParserParserRIGHT_P)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ClickHouseParserParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Parameter()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ClickHouseParserParserCOMMA {
		{
			p.SetState(227)
			p.Match(ClickHouseParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Parameter()
		}

		p.SetState(233)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Expression() IExpressionContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNUMBER, 0)
}

func (s *ParameterContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSTRING, 0)
}

func (s *ParameterContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClickHouseParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ClickHouseParserVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ClickHouseParserParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ClickHouseParserParserRULE_parameter)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClickHouseParserParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(ClickHouseParserParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ClickHouseParserParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(ClickHouseParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ClickHouseParserParserID, ClickHouseParserParserLEFT_P:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Expression()
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
