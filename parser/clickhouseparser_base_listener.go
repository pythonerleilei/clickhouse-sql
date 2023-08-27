// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// BaseClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParserParser.
type BaseClickHouseParserListener struct{}

var _ ClickHouseParserListener = &BaseClickHouseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClickHouseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClickHouseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClickHouseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClickHouseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseClickHouseParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseClickHouseParserListener) ExitStatement(ctx *StatementContext) {}

// EnterCreateStatement is called when production createStatement is entered.
func (s *BaseClickHouseParserListener) EnterCreateStatement(ctx *CreateStatementContext) {}

// ExitCreateStatement is called when production createStatement is exited.
func (s *BaseClickHouseParserListener) ExitCreateStatement(ctx *CreateStatementContext) {}

// EnterCreateDatabaseStatement is called when production createDatabaseStatement is entered.
func (s *BaseClickHouseParserListener) EnterCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {
}

// ExitCreateDatabaseStatement is called when production createDatabaseStatement is exited.
func (s *BaseClickHouseParserListener) ExitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {
}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseClickHouseParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseClickHouseParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseClickHouseParserListener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseClickHouseParserListener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterAlterTableColumn is called when production alterTableColumn is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableColumn(ctx *AlterTableColumnContext) {}

// ExitAlterTableColumn is called when production alterTableColumn is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableColumn(ctx *AlterTableColumnContext) {}

// EnterTalbeIdentifier is called when production talbeIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterTalbeIdentifier(ctx *TalbeIdentifierContext) {}

// ExitTalbeIdentifier is called when production talbeIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitTalbeIdentifier(ctx *TalbeIdentifierContext) {}

// EnterColumnType is called when production columnType is entered.
func (s *BaseClickHouseParserListener) EnterColumnType(ctx *ColumnTypeContext) {}

// ExitColumnType is called when production columnType is exited.
func (s *BaseClickHouseParserListener) ExitColumnType(ctx *ColumnTypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseClickHouseParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseClickHouseParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseClickHouseParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseClickHouseParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseClickHouseParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseClickHouseParserListener) ExitParameter(ctx *ParameterContext) {}
