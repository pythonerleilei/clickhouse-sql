// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// ClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParser.
type ClickHouseParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCreateStatement is called when entering the createStatement production.
	EnterCreateStatement(c *CreateStatementContext)

	// EnterCreateDatabaseStatement is called when entering the createDatabaseStatement production.
	EnterCreateDatabaseStatement(c *CreateDatabaseStatementContext)

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

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)
}
