// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ClickHouseParser.
type ClickHouseParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ClickHouseParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#createStatement.
	VisitCreateStatement(ctx *CreateStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#createDatabaseStatement.
	VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}
}
