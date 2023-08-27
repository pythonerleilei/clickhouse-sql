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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'='", "'('",
		"')'", "','", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "CREATE", "ALTER", "DATABASE", "TABLE", "IF", "NOT", "EXISTS", "ON",
		"CLUSTER", "ENGINE", "COMMENT", "ADD", "COLUMN", "ID", "EQUAL", "LEFT_P",
		"RIGHT_P", "COMMA", "DOT", "NUMBER", "STRING", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.RuleNames = []string{
		"CREATE", "ALTER", "DATABASE", "TABLE", "IF", "NOT", "EXISTS", "ON",
		"CLUSTER", "ENGINE", "COMMENT", "ADD", "COLUMN", "ID", "EQUAL", "LEFT_P",
		"RIGHT_P", "COMMA", "DOT", "NUMBER", "STRING", "A", "B", "C", "D", "E",
		"F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 263, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 5, 13, 179, 8, 13, 10, 13, 12, 13, 182, 9, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4,
		19, 195, 8, 19, 11, 19, 12, 19, 196, 1, 20, 1, 20, 5, 20, 201, 8, 20, 10,
		20, 12, 20, 204, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 202,
		0, 48, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 0, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27,
		57, 28, 59, 29, 61, 30, 63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73, 36,
		75, 37, 77, 38, 79, 39, 81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 91, 45,
		93, 46, 95, 47, 1, 0, 31, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 39, 39, 2, 0, 65, 65, 97,
		97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68, 100, 100,
		2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71, 103, 103,
		2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74, 106, 106,
		2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77, 109, 109,
		2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80, 112, 112,
		2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83, 115, 115,
		2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86, 118, 118,
		2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89, 121, 121,
		2, 0, 90, 90, 122, 122, 3, 0, 9, 10, 13, 13, 32, 32, 264, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		1, 97, 1, 0, 0, 0, 3, 104, 1, 0, 0, 0, 5, 110, 1, 0, 0, 0, 7, 119, 1, 0,
		0, 0, 9, 125, 1, 0, 0, 0, 11, 128, 1, 0, 0, 0, 13, 132, 1, 0, 0, 0, 15,
		139, 1, 0, 0, 0, 17, 142, 1, 0, 0, 0, 19, 150, 1, 0, 0, 0, 21, 157, 1,
		0, 0, 0, 23, 165, 1, 0, 0, 0, 25, 169, 1, 0, 0, 0, 27, 176, 1, 0, 0, 0,
		29, 183, 1, 0, 0, 0, 31, 185, 1, 0, 0, 0, 33, 187, 1, 0, 0, 0, 35, 189,
		1, 0, 0, 0, 37, 191, 1, 0, 0, 0, 39, 194, 1, 0, 0, 0, 41, 198, 1, 0, 0,
		0, 43, 207, 1, 0, 0, 0, 45, 209, 1, 0, 0, 0, 47, 211, 1, 0, 0, 0, 49, 213,
		1, 0, 0, 0, 51, 215, 1, 0, 0, 0, 53, 217, 1, 0, 0, 0, 55, 219, 1, 0, 0,
		0, 57, 221, 1, 0, 0, 0, 59, 223, 1, 0, 0, 0, 61, 225, 1, 0, 0, 0, 63, 227,
		1, 0, 0, 0, 65, 229, 1, 0, 0, 0, 67, 231, 1, 0, 0, 0, 69, 233, 1, 0, 0,
		0, 71, 235, 1, 0, 0, 0, 73, 237, 1, 0, 0, 0, 75, 239, 1, 0, 0, 0, 77, 241,
		1, 0, 0, 0, 79, 243, 1, 0, 0, 0, 81, 245, 1, 0, 0, 0, 83, 247, 1, 0, 0,
		0, 85, 249, 1, 0, 0, 0, 87, 251, 1, 0, 0, 0, 89, 253, 1, 0, 0, 0, 91, 255,
		1, 0, 0, 0, 93, 257, 1, 0, 0, 0, 95, 259, 1, 0, 0, 0, 97, 98, 3, 47, 23,
		0, 98, 99, 3, 77, 38, 0, 99, 100, 3, 51, 25, 0, 100, 101, 3, 43, 21, 0,
		101, 102, 3, 81, 40, 0, 102, 103, 3, 51, 25, 0, 103, 2, 1, 0, 0, 0, 104,
		105, 3, 43, 21, 0, 105, 106, 3, 65, 32, 0, 106, 107, 3, 81, 40, 0, 107,
		108, 3, 51, 25, 0, 108, 109, 3, 77, 38, 0, 109, 4, 1, 0, 0, 0, 110, 111,
		3, 49, 24, 0, 111, 112, 3, 43, 21, 0, 112, 113, 3, 81, 40, 0, 113, 114,
		3, 43, 21, 0, 114, 115, 3, 45, 22, 0, 115, 116, 3, 43, 21, 0, 116, 117,
		3, 79, 39, 0, 117, 118, 3, 51, 25, 0, 118, 6, 1, 0, 0, 0, 119, 120, 3,
		81, 40, 0, 120, 121, 3, 43, 21, 0, 121, 122, 3, 45, 22, 0, 122, 123, 3,
		65, 32, 0, 123, 124, 3, 51, 25, 0, 124, 8, 1, 0, 0, 0, 125, 126, 3, 59,
		29, 0, 126, 127, 3, 53, 26, 0, 127, 10, 1, 0, 0, 0, 128, 129, 3, 69, 34,
		0, 129, 130, 3, 71, 35, 0, 130, 131, 3, 81, 40, 0, 131, 12, 1, 0, 0, 0,
		132, 133, 3, 51, 25, 0, 133, 134, 3, 89, 44, 0, 134, 135, 3, 59, 29, 0,
		135, 136, 3, 79, 39, 0, 136, 137, 3, 81, 40, 0, 137, 138, 3, 79, 39, 0,
		138, 14, 1, 0, 0, 0, 139, 140, 3, 71, 35, 0, 140, 141, 3, 69, 34, 0, 141,
		16, 1, 0, 0, 0, 142, 143, 3, 47, 23, 0, 143, 144, 3, 65, 32, 0, 144, 145,
		3, 83, 41, 0, 145, 146, 3, 79, 39, 0, 146, 147, 3, 81, 40, 0, 147, 148,
		3, 51, 25, 0, 148, 149, 3, 77, 38, 0, 149, 18, 1, 0, 0, 0, 150, 151, 3,
		51, 25, 0, 151, 152, 3, 69, 34, 0, 152, 153, 3, 55, 27, 0, 153, 154, 3,
		59, 29, 0, 154, 155, 3, 69, 34, 0, 155, 156, 3, 51, 25, 0, 156, 20, 1,
		0, 0, 0, 157, 158, 3, 47, 23, 0, 158, 159, 3, 71, 35, 0, 159, 160, 3, 67,
		33, 0, 160, 161, 3, 67, 33, 0, 161, 162, 3, 51, 25, 0, 162, 163, 3, 69,
		34, 0, 163, 164, 3, 81, 40, 0, 164, 22, 1, 0, 0, 0, 165, 166, 3, 43, 21,
		0, 166, 167, 3, 49, 24, 0, 167, 168, 3, 49, 24, 0, 168, 24, 1, 0, 0, 0,
		169, 170, 3, 47, 23, 0, 170, 171, 3, 71, 35, 0, 171, 172, 3, 65, 32, 0,
		172, 173, 3, 83, 41, 0, 173, 174, 3, 67, 33, 0, 174, 175, 3, 69, 34, 0,
		175, 26, 1, 0, 0, 0, 176, 180, 7, 0, 0, 0, 177, 179, 7, 1, 0, 0, 178, 177,
		1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 28, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 61, 0, 0,
		184, 30, 1, 0, 0, 0, 185, 186, 5, 40, 0, 0, 186, 32, 1, 0, 0, 0, 187, 188,
		5, 41, 0, 0, 188, 34, 1, 0, 0, 0, 189, 190, 5, 44, 0, 0, 190, 36, 1, 0,
		0, 0, 191, 192, 5, 46, 0, 0, 192, 38, 1, 0, 0, 0, 193, 195, 7, 2, 0, 0,
		194, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196,
		197, 1, 0, 0, 0, 197, 40, 1, 0, 0, 0, 198, 202, 5, 39, 0, 0, 199, 201,
		8, 3, 0, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 203, 1, 0,
		0, 0, 202, 200, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0,
		205, 206, 5, 39, 0, 0, 206, 42, 1, 0, 0, 0, 207, 208, 7, 4, 0, 0, 208,
		44, 1, 0, 0, 0, 209, 210, 7, 5, 0, 0, 210, 46, 1, 0, 0, 0, 211, 212, 7,
		6, 0, 0, 212, 48, 1, 0, 0, 0, 213, 214, 7, 7, 0, 0, 214, 50, 1, 0, 0, 0,
		215, 216, 7, 8, 0, 0, 216, 52, 1, 0, 0, 0, 217, 218, 7, 9, 0, 0, 218, 54,
		1, 0, 0, 0, 219, 220, 7, 10, 0, 0, 220, 56, 1, 0, 0, 0, 221, 222, 7, 11,
		0, 0, 222, 58, 1, 0, 0, 0, 223, 224, 7, 12, 0, 0, 224, 60, 1, 0, 0, 0,
		225, 226, 7, 13, 0, 0, 226, 62, 1, 0, 0, 0, 227, 228, 7, 14, 0, 0, 228,
		64, 1, 0, 0, 0, 229, 230, 7, 15, 0, 0, 230, 66, 1, 0, 0, 0, 231, 232, 7,
		16, 0, 0, 232, 68, 1, 0, 0, 0, 233, 234, 7, 17, 0, 0, 234, 70, 1, 0, 0,
		0, 235, 236, 7, 18, 0, 0, 236, 72, 1, 0, 0, 0, 237, 238, 7, 19, 0, 0, 238,
		74, 1, 0, 0, 0, 239, 240, 7, 20, 0, 0, 240, 76, 1, 0, 0, 0, 241, 242, 7,
		21, 0, 0, 242, 78, 1, 0, 0, 0, 243, 244, 7, 22, 0, 0, 244, 80, 1, 0, 0,
		0, 245, 246, 7, 23, 0, 0, 246, 82, 1, 0, 0, 0, 247, 248, 7, 24, 0, 0, 248,
		84, 1, 0, 0, 0, 249, 250, 7, 25, 0, 0, 250, 86, 1, 0, 0, 0, 251, 252, 7,
		26, 0, 0, 252, 88, 1, 0, 0, 0, 253, 254, 7, 27, 0, 0, 254, 90, 1, 0, 0,
		0, 255, 256, 7, 28, 0, 0, 256, 92, 1, 0, 0, 0, 257, 258, 7, 29, 0, 0, 258,
		94, 1, 0, 0, 0, 259, 260, 7, 30, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262,
		6, 47, 0, 0, 262, 96, 1, 0, 0, 0, 5, 0, 178, 180, 196, 202, 1, 6, 0, 0,
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
	ClickHouseParserLexerALTER    = 2
	ClickHouseParserLexerDATABASE = 3
	ClickHouseParserLexerTABLE    = 4
	ClickHouseParserLexerIF       = 5
	ClickHouseParserLexerNOT      = 6
	ClickHouseParserLexerEXISTS   = 7
	ClickHouseParserLexerON       = 8
	ClickHouseParserLexerCLUSTER  = 9
	ClickHouseParserLexerENGINE   = 10
	ClickHouseParserLexerCOMMENT  = 11
	ClickHouseParserLexerADD      = 12
	ClickHouseParserLexerCOLUMN   = 13
	ClickHouseParserLexerID       = 14
	ClickHouseParserLexerEQUAL    = 15
	ClickHouseParserLexerLEFT_P   = 16
	ClickHouseParserLexerRIGHT_P  = 17
	ClickHouseParserLexerCOMMA    = 18
	ClickHouseParserLexerDOT      = 19
	ClickHouseParserLexerNUMBER   = 20
	ClickHouseParserLexerSTRING   = 21
	ClickHouseParserLexerB        = 22
	ClickHouseParserLexerC        = 23
	ClickHouseParserLexerD        = 24
	ClickHouseParserLexerE        = 25
	ClickHouseParserLexerF        = 26
	ClickHouseParserLexerG        = 27
	ClickHouseParserLexerH        = 28
	ClickHouseParserLexerI        = 29
	ClickHouseParserLexerJ        = 30
	ClickHouseParserLexerK        = 31
	ClickHouseParserLexerL        = 32
	ClickHouseParserLexerM        = 33
	ClickHouseParserLexerN        = 34
	ClickHouseParserLexerO        = 35
	ClickHouseParserLexerP        = 36
	ClickHouseParserLexerQ        = 37
	ClickHouseParserLexerR        = 38
	ClickHouseParserLexerS        = 39
	ClickHouseParserLexerT        = 40
	ClickHouseParserLexerU        = 41
	ClickHouseParserLexerV        = 42
	ClickHouseParserLexerW        = 43
	ClickHouseParserLexerX        = 44
	ClickHouseParserLexerY        = 45
	ClickHouseParserLexerZ        = 46
	ClickHouseParserLexerWS       = 47
)
