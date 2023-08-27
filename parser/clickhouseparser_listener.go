// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// ClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParserParser.
type ClickHouseParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCreateStatement is called when entering the createStatement production.
	EnterCreateStatement(c *CreateStatementContext)

	// EnterCreateDatabaseStatement is called when entering the createDatabaseStatement production.
	EnterCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterAlterTableColumn is called when entering the alterTableColumn production.
	EnterAlterTableColumn(c *AlterTableColumnContext)

	// EnterTalbeIdentifier is called when entering the talbeIdentifier production.
	EnterTalbeIdentifier(c *TalbeIdentifierContext)

	// EnterColumnType is called when entering the columnType production.
	EnterColumnType(c *ColumnTypeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCreateStatement is called when exiting the createStatement production.
	ExitCreateStatement(c *CreateStatementContext)

	// ExitCreateDatabaseStatement is called when exiting the createDatabaseStatement production.
	ExitCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitAlterTableColumn is called when exiting the alterTableColumn production.
	ExitAlterTableColumn(c *AlterTableColumnContext)

	// ExitTalbeIdentifier is called when exiting the talbeIdentifier production.
	ExitTalbeIdentifier(c *TalbeIdentifierContext)

	// ExitColumnType is called when exiting the columnType production.
	ExitColumnType(c *ColumnTypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)
}
