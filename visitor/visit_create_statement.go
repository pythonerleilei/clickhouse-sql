package visitor

import parser "clickhouse-sql/parser"

func (v *BaseAstClickhouseParserVisitor) VisitCreateStatement(ctx *parser.CreateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

type CreateDatabase struct {
	FlagNotExists	bool
	DatabaseName	string
	ClusterName		string
	EngineName		string
	EngineParams	[]*Expression
	Comment			string
}

// create database
func (v *BaseAstClickhouseParserVisitor) VisitCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) interface{} {
	createDatabase := &CreateDatabase{}
	if ctx.EXISTS() != nil {
		createDatabase.FlagNotExists = true
	}

	createDatabase.DatabaseName = ctx.GetDb_name().GetText();
	if (ctx.GetCluster_name() != nil) {
		createDatabase.ClusterName = ctx.GetCluster_name().GetText()
	}
	
	if ctx.GetEngine_name() != nil {
		e := v.Visit(ctx.GetEngine_name()).(*Expression)
		createDatabase.EngineParams = e.children
		createDatabase.EngineName = e.Name()
	}
	
	if (ctx.COMMENT() != nil ){
		createDatabase.Comment = ctx.GetComment().GetText();
	}
	return createDatabase
}

//