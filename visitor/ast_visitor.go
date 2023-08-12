package visitor

import (
	parser "clickhouse-sql/parser"

	"github.com/antlr4-go/antlr/v4"
)

type BaseAstClickhouseParserVisitor struct {
	*parser.BaseClickHouseParserVisitor

	Stream	*antlr.InputStream
	Lexer	*parser.ClickHouseParserLexer
	Parser	*parser.ClickHouseParserParser

}

func NewBaseAstClickhouseParserVisitor(sql string) *BaseAstClickhouseParserVisitor {
	stream := antlr.NewInputStream(sql)
	lexer := parser.NewClickHouseParserLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewClickHouseParserParser(tokenStream)
	return &BaseAstClickhouseParserVisitor{
		Stream: stream,
		Lexer: lexer,
		Parser: parser,
	}
}

func (v *BaseAstClickhouseParserVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *BaseAstClickhouseParserVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	child := node.GetChild(0).(antlr.ParseTree)
	return child.Accept(v)
}

func (v *BaseAstClickhouseParserVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}