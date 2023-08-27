// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ClickHouseParserParser.
type ClickHouseParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ClickHouseParserParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#createStatement.
	VisitCreateStatement(ctx *CreateStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#createDatabaseStatement.
	VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#alterTableColumn.
	VisitAlterTableColumn(ctx *AlterTableColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#talbeIdentifier.
	VisitTalbeIdentifier(ctx *TalbeIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#columnType.
	VisitColumnType(ctx *ColumnTypeContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by ClickHouseParserParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}
}
