// Code generated from ClickHouseParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

type BaseClickHouseParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseClickHouseParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateStatement(ctx *CreateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterStatement(ctx *AlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableColumn(ctx *AlterTableColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTalbeIdentifier(ctx *TalbeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnType(ctx *ColumnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}
