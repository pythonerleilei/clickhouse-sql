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

type ClickHouseParser struct {
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
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	staticData.RuleNames = []string{
		"statement", "createStatement", "createDatabaseStatement", "parameters",
		"parameter",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 52, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 20, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 26, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 35, 8, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 5, 3, 45, 8, 3, 10, 3, 12, 3, 48, 9, 3, 1, 4, 1, 4, 1, 4, 0, 0, 5,
		0, 2, 4, 6, 8, 0, 1, 1, 0, 15, 16, 51, 0, 10, 1, 0, 0, 0, 2, 12, 1, 0,
		0, 0, 4, 14, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 49, 1, 0, 0, 0, 10, 11,
		3, 2, 1, 0, 11, 1, 1, 0, 0, 0, 12, 13, 3, 4, 2, 0, 13, 3, 1, 0, 0, 0, 14,
		15, 5, 1, 0, 0, 15, 19, 5, 2, 0, 0, 16, 17, 5, 3, 0, 0, 17, 18, 5, 4, 0,
		0, 18, 20, 5, 5, 0, 0, 19, 16, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21,
		1, 0, 0, 0, 21, 25, 5, 10, 0, 0, 22, 23, 5, 6, 0, 0, 23, 24, 5, 7, 0, 0,
		24, 26, 5, 10, 0, 0, 25, 22, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 36, 1,
		0, 0, 0, 27, 28, 5, 8, 0, 0, 28, 29, 5, 11, 0, 0, 29, 34, 5, 10, 0, 0,
		30, 31, 5, 12, 0, 0, 31, 32, 3, 6, 3, 0, 32, 33, 5, 13, 0, 0, 33, 35, 1,
		0, 0, 0, 34, 30, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 37, 1, 0, 0, 0, 36,
		27, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 9, 0,
		0, 39, 40, 5, 16, 0, 0, 40, 5, 1, 0, 0, 0, 41, 46, 3, 8, 4, 0, 42, 43,
		5, 14, 0, 0, 43, 45, 3, 8, 4, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0,
		46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 46, 1, 0,
		0, 0, 49, 50, 7, 0, 0, 0, 50, 9, 1, 0, 0, 0, 5, 19, 25, 34, 36, 46,
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

// ClickHouseParserInit initializes any static state used to implement ClickHouseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClickHouseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClickHouseParserInit() {
	staticData := &ClickHouseParserParserStaticData
	staticData.once.Do(clickhouseparserParserInit)
}

// NewClickHouseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClickHouseParser(input antlr.TokenStream) *ClickHouseParser {
	ClickHouseParserInit()
	this := new(ClickHouseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ClickHouseParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ClickHouseParser.g4"

	return this
}

// ClickHouseParser tokens.
const (
	ClickHouseParserEOF      = antlr.TokenEOF
	ClickHouseParserCREATE   = 1
	ClickHouseParserDATABASE = 2
	ClickHouseParserIF       = 3
	ClickHouseParserNOT      = 4
	ClickHouseParserEXISTS   = 5
	ClickHouseParserON       = 6
	ClickHouseParserCLUSTER  = 7
	ClickHouseParserENGINE   = 8
	ClickHouseParserCOMMENT  = 9
	ClickHouseParserID       = 10
	ClickHouseParserEQUAL    = 11
	ClickHouseParserLEFT_P   = 12
	ClickHouseParserRIGHT_P  = 13
	ClickHouseParserCOMMA    = 14
	ClickHouseParserNUMBER   = 15
	ClickHouseParserSTRING   = 16
	ClickHouseParserB        = 17
	ClickHouseParserC        = 18
	ClickHouseParserD        = 19
	ClickHouseParserE        = 20
	ClickHouseParserF        = 21
	ClickHouseParserG        = 22
	ClickHouseParserH        = 23
	ClickHouseParserI        = 24
	ClickHouseParserJ        = 25
	ClickHouseParserK        = 26
	ClickHouseParserL        = 27
	ClickHouseParserM        = 28
	ClickHouseParserN        = 29
	ClickHouseParserO        = 30
	ClickHouseParserP        = 31
	ClickHouseParserQ        = 32
	ClickHouseParserR        = 33
	ClickHouseParserS        = 34
	ClickHouseParserT        = 35
	ClickHouseParserU        = 36
	ClickHouseParserV        = 37
	ClickHouseParserW        = 38
	ClickHouseParserX        = 39
	ClickHouseParserY        = 40
	ClickHouseParserZ        = 41
)

// ClickHouseParser rules.
const (
	ClickHouseParserRULE_statement               = 0
	ClickHouseParserRULE_createStatement         = 1
	ClickHouseParserRULE_createDatabaseStatement = 2
	ClickHouseParserRULE_parameters              = 3
	ClickHouseParserRULE_parameter               = 4
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
	p.RuleIndex = ClickHouseParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserRULE_statement

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

func (p *ClickHouseParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClickHouseParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
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
	p.RuleIndex = ClickHouseParserRULE_createStatement
	return p
}

func InitEmptyCreateStatementContext(p *CreateStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_createStatement
}

func (*CreateStatementContext) IsCreateStatementContext() {}

func NewCreateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateStatementContext {
	var p = new(CreateStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserRULE_createStatement

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

func (p *ClickHouseParser) CreateStatement() (localctx ICreateStatementContext) {
	localctx = NewCreateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClickHouseParserRULE_createStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
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

	// GetComment returns the comment token.
	GetComment() antlr.Token

	// SetDb_name sets the db_name token.
	SetDb_name(antlr.Token)

	// SetCluster_name sets the cluster_name token.
	SetCluster_name(antlr.Token)

	// SetComment sets the comment token.
	SetComment(antlr.Token)

	// GetEngine_params returns the engine_params rule contexts.
	GetEngine_params() IParametersContext

	// SetEngine_params sets the engine_params rule contexts.
	SetEngine_params(IParametersContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	STRING() antlr.TerminalNode
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	ON() antlr.TerminalNode
	CLUSTER() antlr.TerminalNode
	ENGINE() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	Parameters() IParametersContext

	// IsCreateDatabaseStatementContext differentiates from other interfaces.
	IsCreateDatabaseStatementContext()
}

type CreateDatabaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	db_name       antlr.Token
	cluster_name  antlr.Token
	engine_params IParametersContext
	comment       antlr.Token
}

func NewEmptyCreateDatabaseStatementContext() *CreateDatabaseStatementContext {
	var p = new(CreateDatabaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_createDatabaseStatement
	return p
}

func InitEmptyCreateDatabaseStatementContext(p *CreateDatabaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_createDatabaseStatement
}

func (*CreateDatabaseStatementContext) IsCreateDatabaseStatementContext() {}

func NewCreateDatabaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDatabaseStatementContext {
	var p = new(CreateDatabaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserRULE_createDatabaseStatement

	return p
}

func (s *CreateDatabaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDatabaseStatementContext) GetDb_name() antlr.Token { return s.db_name }

func (s *CreateDatabaseStatementContext) GetCluster_name() antlr.Token { return s.cluster_name }

func (s *CreateDatabaseStatementContext) GetComment() antlr.Token { return s.comment }

func (s *CreateDatabaseStatementContext) SetDb_name(v antlr.Token) { s.db_name = v }

func (s *CreateDatabaseStatementContext) SetCluster_name(v antlr.Token) { s.cluster_name = v }

func (s *CreateDatabaseStatementContext) SetComment(v antlr.Token) { s.comment = v }

func (s *CreateDatabaseStatementContext) GetEngine_params() IParametersContext {
	return s.engine_params
}

func (s *CreateDatabaseStatementContext) SetEngine_params(v IParametersContext) { s.engine_params = v }

func (s *CreateDatabaseStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserCREATE, 0)
}

func (s *CreateDatabaseStatementContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserDATABASE, 0)
}

func (s *CreateDatabaseStatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserCOMMENT, 0)
}

func (s *CreateDatabaseStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ClickHouseParserID)
}

func (s *CreateDatabaseStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserID, i)
}

func (s *CreateDatabaseStatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserSTRING, 0)
}

func (s *CreateDatabaseStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserIF, 0)
}

func (s *CreateDatabaseStatementContext) NOT() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserNOT, 0)
}

func (s *CreateDatabaseStatementContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserEXISTS, 0)
}

func (s *CreateDatabaseStatementContext) ON() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserON, 0)
}

func (s *CreateDatabaseStatementContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserCLUSTER, 0)
}

func (s *CreateDatabaseStatementContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserENGINE, 0)
}

func (s *CreateDatabaseStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserEQUAL, 0)
}

func (s *CreateDatabaseStatementContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserLEFT_P, 0)
}

func (s *CreateDatabaseStatementContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserRIGHT_P, 0)
}

func (s *CreateDatabaseStatementContext) Parameters() IParametersContext {
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

func (p *ClickHouseParser) CreateDatabaseStatement() (localctx ICreateDatabaseStatementContext) {
	localctx = NewCreateDatabaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ClickHouseParserRULE_createDatabaseStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Match(ClickHouseParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(15)
		p.Match(ClickHouseParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserIF {
		{
			p.SetState(16)
			p.Match(ClickHouseParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Match(ClickHouseParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(18)
			p.Match(ClickHouseParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(21)

		var _m = p.Match(ClickHouseParserID)

		localctx.(*CreateDatabaseStatementContext).db_name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserON {
		{
			p.SetState(22)
			p.Match(ClickHouseParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(ClickHouseParserCLUSTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)

			var _m = p.Match(ClickHouseParserID)

			localctx.(*CreateDatabaseStatementContext).cluster_name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ClickHouseParserENGINE {
		{
			p.SetState(27)
			p.Match(ClickHouseParserENGINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.Match(ClickHouseParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(ClickHouseParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ClickHouseParserLEFT_P {
			{
				p.SetState(30)
				p.Match(ClickHouseParserLEFT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(31)

				var _x = p.Parameters()

				localctx.(*CreateDatabaseStatementContext).engine_params = _x
			}
			{
				p.SetState(32)
				p.Match(ClickHouseParserRIGHT_P)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(38)
		p.Match(ClickHouseParserCOMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)

		var _m = p.Match(ClickHouseParserSTRING)

		localctx.(*CreateDatabaseStatementContext).comment = _m
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
	p.RuleIndex = ClickHouseParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserRULE_parameters

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
	return s.GetTokens(ClickHouseParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ClickHouseParserCOMMA, i)
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

func (p *ClickHouseParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ClickHouseParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Parameter()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ClickHouseParserCOMMA {
		{
			p.SetState(42)
			p.Match(ClickHouseParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Parameter()
		}

		p.SetState(48)
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
	p.RuleIndex = ClickHouseParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ClickHouseParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClickHouseParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserNUMBER, 0)
}

func (s *ParameterContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClickHouseParserSTRING, 0)
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

func (p *ClickHouseParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ClickHouseParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ClickHouseParserNUMBER || _la == ClickHouseParserSTRING) {
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
