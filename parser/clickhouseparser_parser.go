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
		"statement", "createStatement", "createDatabaseStatement", "createTableStatement",
		"alterStatement", "alterTableColumn", "talbeIdentifier", "columnType",
		"expression", "parameters", "parameter",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 110, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 3, 0, 25, 8, 0, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 36, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 42, 8,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 58, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 70, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 77, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 83, 8, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 92, 8, 8, 1, 8, 3, 8, 95, 8, 8, 1, 9,
		1, 9, 1, 9, 5, 9, 100, 8, 9, 10, 9, 12, 9, 103, 9, 9, 1, 10, 1, 10, 1,
		10, 3, 10, 108, 8, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 0, 0, 113, 0, 24, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 30, 1, 0, 0,
		0, 6, 52, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 63, 1, 0, 0, 0, 12, 82, 1,
		0, 0, 0, 14, 86, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 96, 1, 0, 0, 0, 20,
		107, 1, 0, 0, 0, 22, 25, 3, 2, 1, 0, 23, 25, 3, 8, 4, 0, 24, 22, 1, 0,
		0, 0, 24, 23, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26, 29, 3, 4, 2, 0, 27, 29,
		3, 6, 3, 0, 28, 26, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 3, 1, 0, 0, 0,
		30, 31, 5, 1, 0, 0, 31, 35, 5, 3, 0, 0, 32, 33, 5, 5, 0, 0, 33, 34, 5,
		6, 0, 0, 34, 36, 5, 7, 0, 0, 35, 32, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36,
		37, 1, 0, 0, 0, 37, 41, 5, 14, 0, 0, 38, 39, 5, 8, 0, 0, 39, 40, 5, 9,
		0, 0, 40, 42, 5, 14, 0, 0, 41, 38, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42,
		46, 1, 0, 0, 0, 43, 44, 5, 10, 0, 0, 44, 45, 5, 15, 0, 0, 45, 47, 3, 16,
		8, 0, 46, 43, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 49,
		5, 11, 0, 0, 49, 51, 5, 21, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0,
		0, 51, 5, 1, 0, 0, 0, 52, 53, 5, 1, 0, 0, 53, 57, 5, 4, 0, 0, 54, 55, 5,
		5, 0, 0, 55, 56, 5, 6, 0, 0, 56, 58, 5, 7, 0, 0, 57, 54, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 3, 12, 6, 0, 60, 7, 1, 0, 0,
		0, 61, 62, 3, 10, 5, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 2, 0, 0, 64, 65,
		5, 4, 0, 0, 65, 69, 3, 12, 6, 0, 66, 67, 5, 8, 0, 0, 67, 68, 5, 9, 0, 0,
		68, 70, 5, 14, 0, 0, 69, 66, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1,
		0, 0, 0, 71, 72, 5, 12, 0, 0, 72, 76, 5, 13, 0, 0, 73, 74, 5, 5, 0, 0,
		74, 75, 5, 6, 0, 0, 75, 77, 5, 7, 0, 0, 76, 73, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 3, 14, 7, 0, 79, 11, 1, 0, 0, 0, 80,
		81, 5, 14, 0, 0, 81, 83, 5, 19, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0,
		0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 5, 14, 0, 0, 85, 13, 1, 0, 0, 0, 86,
		87, 3, 16, 8, 0, 87, 15, 1, 0, 0, 0, 88, 94, 5, 14, 0, 0, 89, 91, 5, 16,
		0, 0, 90, 92, 3, 18, 9, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 95, 5, 17, 0, 0, 94, 89, 1, 0, 0, 0, 94, 95, 1, 0,
		0, 0, 95, 17, 1, 0, 0, 0, 96, 101, 3, 20, 10, 0, 97, 98, 5, 18, 0, 0, 98,
		100, 3, 20, 10, 0, 99, 97, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1,
		0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 19, 1, 0, 0, 0, 103, 101, 1, 0, 0,
		0, 104, 108, 5, 20, 0, 0, 105, 108, 5, 21, 0, 0, 106, 108, 3, 16, 8, 0,
		107, 104, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108,
		21, 1, 0, 0, 0, 14, 24, 28, 35, 41, 46, 50, 57, 69, 76, 82, 91, 94, 101,
		107,
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
	ClickHouseParserParserEOF      = antlr.TokenEOF
	ClickHouseParserParserCREATE   = 1
	ClickHouseParserParserALTER    = 2
	ClickHouseParserParserDATABASE = 3
	ClickHouseParserParserTABLE    = 4
	ClickHouseParserParserIF       = 5
	ClickHouseParserParserNOT      = 6
	ClickHouseParserParserEXISTS   = 7
	ClickHouseParserParserON       = 8
	ClickHouseParserParserCLUSTER  = 9
	ClickHouseParserParserENGINE   = 10
	ClickHouseParserParserCOMMENT  = 11
	ClickHouseParserParserADD      = 12
	ClickHouseParserParserCOLUMN   = 13
	ClickHouseParserParserID       = 14
	ClickHouseParserParserEQUAL    = 15
	ClickHouseParserParserLEFT_P   = 16
	ClickHouseParserParserRIGHT_P  = 17
	ClickHouseParserParserCOMMA    = 18
	ClickHouseParserParserDOT      = 19
	ClickHouseParserParserNUMBER   = 20
	ClickHouseParserParserSTRING   = 21
	ClickHouseParserParserB        = 22
	ClickHouseParserParserC        = 23
	ClickHouseParserParserD        = 24
	ClickHouseParserParserE        = 25
	ClickHouseParserParserF        = 26
	ClickHouseParserParserG        = 27
	ClickHouseParserParserH        = 28
	ClickHouseParserParserI        = 29
	ClickHouseParserParserJ        = 30
	ClickHouseParserParserK        = 31
	ClickHouseParserParserL        = 32
	ClickHouseParserParserM        = 33
	ClickHouseParserParserN        = 34
	ClickHouseParserParserO        = 35
	ClickHouseParserParserP        = 36
	ClickHouseParserParserQ        = 37
	ClickHouseParserParserR        = 38
	ClickHouseParserParserS        = 39
	ClickHouseParserParserT        = 40
	ClickHouseParserParserU        = 41
	ClickHouseParserParserV        = 42
	ClickHouseParserParserW        = 43
	ClickHouseParserParserX        = 44
	ClickHouseParserParserY        = 45
	ClickHouseParserParserZ        = 46
	ClickHouseParserParserWS       = 47
)

// ClickHouseParserParser rules.
const (
	ClickHouseParserParserRULE_statement               = 0
	ClickHouseParserParserRULE_createStatement         = 1
	ClickHouseParserParserRULE_createDatabaseStatement = 2
	ClickHouseParserParserRULE_createTableStatement    = 3
	ClickHouseParserParserRULE_alterStatement          = 4
	ClickHouseParserParserRULE_alterTableColumn        = 5
	ClickHouseParserParserRULE_talbeIdentifier         = 6
	ClickHouseParserParserRULE_columnType              = 7
	ClickHouseParserParserRULE_expression              = 8
	ClickHouseParserParserRULE_parameters              = 9
	ClickHouseParserParserRULE_parameter               = 10
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClickHouseParserParserCREATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.CreateStatement()
		}

	case ClickHouseParserParserALTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.CreateDatabaseStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
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
		p.SetState(30)
		p.Match(ClickHouseParserParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.Match(ClickHouseParserParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(32)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(37)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*CreateDatabaseStatementContext).db_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(38)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*CreateDatabaseStatementContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserENGINE {
		{
			p.SetState(43)
			p.Match(ClickHouseParserParserENGINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(ClickHouseParserParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)

			var _x = p.Expression()

			localctx.(*CreateDatabaseStatementContext).engine_name = _x
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMENT {
		{
			p.SetState(48)
			p.Match(ClickHouseParserParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)

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

	// GetTable_name returns the table_name rule contexts.
	GetTable_name() ITalbeIdentifierContext

	// SetTable_name sets the table_name rule contexts.
	SetTable_name(ITalbeIdentifierContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TalbeIdentifier() ITalbeIdentifierContext
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	table_name ITalbeIdentifierContext
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

func (s *CreateTableStatementContext) GetTable_name() ITalbeIdentifierContext { return s.table_name }

func (s *CreateTableStatementContext) SetTable_name(v ITalbeIdentifierContext) { s.table_name = v }

func (s *CreateTableStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCREATE, 0)
}

func (s *CreateTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserTABLE, 0)
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

func (s *CreateTableStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserIF, 0)
}

func (s *CreateTableStatementContext) NOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserNOT, 0)
}

func (s *CreateTableStatementContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserEXISTS, 0)
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
		p.SetState(52)
		p.Match(ClickHouseParserParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(ClickHouseParserParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(54)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(59)

		var _x = p.TalbeIdentifier()

		localctx.(*CreateTableStatementContext).table_name = _x
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
	p.EnterRule(localctx, 8, ClickHouseParserParserRULE_alterStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
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
	p.EnterRule(localctx, 10, ClickHouseParserParserRULE_alterTableColumn)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(ClickHouseParserParserALTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(ClickHouseParserParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)

		var _x = p.TalbeIdentifier()

		localctx.(*AlterTableColumnContext).table_name = _x
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(66)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*AlterTableColumnContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(71)
		p.Match(ClickHouseParserParserADD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(ClickHouseParserParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(73)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(78)

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
	p.EnterRule(localctx, 12, ClickHouseParserParserRULE_talbeIdentifier)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(80)
			p.Match(ClickHouseParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
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
		p.SetState(84)
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
	p.EnterRule(localctx, 14, ClickHouseParserParserRULE_columnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
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
	p.EnterRule(localctx, 16, ClickHouseParserParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ClickHouseParserParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
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
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3162112) != 0 {
			{
				p.SetState(90)
				p.Parameters()
			}

		}
		{
			p.SetState(93)
			p.Match(ClickHouseParserParserRIGHT_P)
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
	p.EnterRule(localctx, 18, ClickHouseParserParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Parameter()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ClickHouseParserParserCOMMA {
		{
			p.SetState(97)
			p.Match(ClickHouseParserParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Parameter()
		}

		p.SetState(103)
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
	p.EnterRule(localctx, 20, ClickHouseParserParserRULE_parameter)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ClickHouseParserParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(ClickHouseParserParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ClickHouseParserParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(ClickHouseParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ClickHouseParserParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
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
