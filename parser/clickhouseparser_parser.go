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
		"", "", "", "", "", "", "", "", "", "", "", "'='", "'('", "')'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "CREATE", "DATABASE", "IF", "NOT", "EXISTS", "ON", "CLUSTER", "ENGINE",
		"COMMENT", "ID", "EQUAL", "LEFT_P", "RIGHT_P", "COMMA", "NUMBER", "STRING",
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "WS",
	}
	staticData.RuleNames = []string{
		"statement", "createStatement", "createDatabaseStatement", "parameter",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 50, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 18, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 24, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5,
		2, 33, 8, 2, 10, 2, 12, 2, 36, 9, 2, 1, 2, 1, 2, 3, 2, 40, 8, 2, 3, 2,
		42, 8, 2, 1, 2, 1, 2, 3, 2, 46, 8, 2, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2,
		4, 6, 0, 1, 1, 0, 15, 16, 51, 0, 8, 1, 0, 0, 0, 2, 10, 1, 0, 0, 0, 4, 12,
		1, 0, 0, 0, 6, 47, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 1, 1, 0, 0, 0, 10,
		11, 3, 4, 2, 0, 11, 3, 1, 0, 0, 0, 12, 13, 5, 1, 0, 0, 13, 17, 5, 2, 0,
		0, 14, 15, 5, 3, 0, 0, 15, 16, 5, 4, 0, 0, 16, 18, 5, 5, 0, 0, 17, 14,
		1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 23, 5, 10, 0, 0,
		20, 21, 5, 6, 0, 0, 21, 22, 5, 7, 0, 0, 22, 24, 5, 10, 0, 0, 23, 20, 1,
		0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 41, 1, 0, 0, 0, 25, 26, 5, 8, 0, 0, 26,
		27, 5, 11, 0, 0, 27, 39, 5, 10, 0, 0, 28, 29, 5, 12, 0, 0, 29, 34, 3, 6,
		3, 0, 30, 31, 5, 14, 0, 0, 31, 33, 3, 6, 3, 0, 32, 30, 1, 0, 0, 0, 33,
		36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 37, 1, 0, 0,
		0, 36, 34, 1, 0, 0, 0, 37, 38, 5, 13, 0, 0, 38, 40, 1, 0, 0, 0, 39, 28,
		1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 25, 1, 0, 0, 0,
		41, 42, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 44, 5, 9, 0, 0, 44, 46, 5,
		16, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 5, 1, 0, 0, 0, 47,
		48, 7, 0, 0, 0, 48, 7, 1, 0, 0, 0, 6, 17, 23, 34, 39, 41, 45,
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
	ClickHouseParserParserDATABASE = 2
	ClickHouseParserParserIF       = 3
	ClickHouseParserParserNOT      = 4
	ClickHouseParserParserEXISTS   = 5
	ClickHouseParserParserON       = 6
	ClickHouseParserParserCLUSTER  = 7
	ClickHouseParserParserENGINE   = 8
	ClickHouseParserParserCOMMENT  = 9
	ClickHouseParserParserID       = 10
	ClickHouseParserParserEQUAL    = 11
	ClickHouseParserParserLEFT_P   = 12
	ClickHouseParserParserRIGHT_P  = 13
	ClickHouseParserParserCOMMA    = 14
	ClickHouseParserParserNUMBER   = 15
	ClickHouseParserParserSTRING   = 16
	ClickHouseParserParserB        = 17
	ClickHouseParserParserC        = 18
	ClickHouseParserParserD        = 19
	ClickHouseParserParserE        = 20
	ClickHouseParserParserF        = 21
	ClickHouseParserParserG        = 22
	ClickHouseParserParserH        = 23
	ClickHouseParserParserI        = 24
	ClickHouseParserParserJ        = 25
	ClickHouseParserParserK        = 26
	ClickHouseParserParserL        = 27
	ClickHouseParserParserM        = 28
	ClickHouseParserParserN        = 29
	ClickHouseParserParserO        = 30
	ClickHouseParserParserP        = 31
	ClickHouseParserParserQ        = 32
	ClickHouseParserParserR        = 33
	ClickHouseParserParserS        = 34
	ClickHouseParserParserT        = 35
	ClickHouseParserParserU        = 36
	ClickHouseParserParserV        = 37
	ClickHouseParserParserW        = 38
	ClickHouseParserParserX        = 39
	ClickHouseParserParserY        = 40
	ClickHouseParserParserZ        = 41
	ClickHouseParserParserWS       = 42
)

// ClickHouseParserParser rules.
const (
	ClickHouseParserParserRULE_statement               = 0
	ClickHouseParserParserRULE_createStatement         = 1
	ClickHouseParserParserRULE_createDatabaseStatement = 2
	ClickHouseParserParserRULE_parameter               = 3
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateStatement() ICreateStatementContext

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.CreateStatement()
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.CreateDatabaseStatement()
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

	// GetEngine_name returns the engine_name token.
	GetEngine_name() antlr.Token

	// GetComment returns the comment token.
	GetComment() antlr.Token

	// SetDb_name sets the db_name token.
	SetDb_name(antlr.Token)

	// SetCluster_name sets the cluster_name token.
	SetCluster_name(antlr.Token)

	// SetEngine_name sets the engine_name token.
	SetEngine_name(antlr.Token)

	// SetComment sets the comment token.
	SetComment(antlr.Token)

	// Get_parameter returns the _parameter rule contexts.
	Get_parameter() IParameterContext

	// Set_parameter sets the _parameter rule contexts.
	Set_parameter(IParameterContext)

	// GetEngine_params returns the engine_params rule context list.
	GetEngine_params() []IParameterContext

	// SetEngine_params sets the engine_params rule context list.
	SetEngine_params([]IParameterContext)

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
	STRING() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCreateDatabaseStatementContext differentiates from other interfaces.
	IsCreateDatabaseStatementContext()
}

type CreateDatabaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	db_name       antlr.Token
	cluster_name  antlr.Token
	engine_name   antlr.Token
	_parameter    IParameterContext
	engine_params []IParameterContext
	comment       antlr.Token
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

func (s *CreateDatabaseStatementContext) GetEngine_name() antlr.Token { return s.engine_name }

func (s *CreateDatabaseStatementContext) GetComment() antlr.Token { return s.comment }

func (s *CreateDatabaseStatementContext) SetDb_name(v antlr.Token) { s.db_name = v }

func (s *CreateDatabaseStatementContext) SetCluster_name(v antlr.Token) { s.cluster_name = v }

func (s *CreateDatabaseStatementContext) SetEngine_name(v antlr.Token) { s.engine_name = v }

func (s *CreateDatabaseStatementContext) SetComment(v antlr.Token) { s.comment = v }

func (s *CreateDatabaseStatementContext) Get_parameter() IParameterContext { return s._parameter }

func (s *CreateDatabaseStatementContext) Set_parameter(v IParameterContext) { s._parameter = v }

func (s *CreateDatabaseStatementContext) GetEngine_params() []IParameterContext {
	return s.engine_params
}

func (s *CreateDatabaseStatementContext) SetEngine_params(v []IParameterContext) { s.engine_params = v }

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

func (s *CreateDatabaseStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserSTRING, 0)
}

func (s *CreateDatabaseStatementContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserLEFT_P, 0)
}

func (s *CreateDatabaseStatementContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserRIGHT_P, 0)
}

func (s *CreateDatabaseStatementContext) AllParameter() []IParameterContext {
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

func (s *CreateDatabaseStatementContext) Parameter(i int) IParameterContext {
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

func (s *CreateDatabaseStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserParserCOMMA)
}

func (s *CreateDatabaseStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserParserCOMMA, i)
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
		p.SetState(12)
		p.Match(ClickHouseParserParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(13)
		p.Match(ClickHouseParserParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserIF {
		{
			p.SetState(14)
			p.Match(ClickHouseParserParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(15)
			p.Match(ClickHouseParserParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(16)
			p.Match(ClickHouseParserParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(19)

		var _m = p.Match(ClickHouseParserParserID)

		localctx.(*CreateDatabaseStatementContext).db_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserON {
		{
			p.SetState(20)
			p.Match(ClickHouseParserParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.Match(ClickHouseParserParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*CreateDatabaseStatementContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserENGINE {
		{
			p.SetState(25)
			p.Match(ClickHouseParserParserENGINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Match(ClickHouseParserParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)

			var _m = p.Match(ClickHouseParserParserID)

			localctx.(*CreateDatabaseStatementContext).engine_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ClickHouseParserParserLEFT_P {
			{
				p.SetState(28)
				p.Match(ClickHouseParserParserLEFT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(29)

				var _x = p.Parameter()

				localctx.(*CreateDatabaseStatementContext)._parameter = _x
			}
			localctx.(*CreateDatabaseStatementContext).engine_params = append(localctx.(*CreateDatabaseStatementContext).engine_params, localctx.(*CreateDatabaseStatementContext)._parameter)
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ClickHouseParserParserCOMMA {
				{
					p.SetState(30)
					p.Match(ClickHouseParserParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(31)
					p.Parameter()
				}

				p.SetState(36)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(37)
				p.Match(ClickHouseParserParserRIGHT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserParserCOMMENT {
		{
			p.SetState(43)
			p.Match(ClickHouseParserParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)

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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

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
	p.EnterRule(localctx, 6, ClickHouseParserParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ClickHouseParserParserNUMBER || _la == ClickHouseParserParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
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
