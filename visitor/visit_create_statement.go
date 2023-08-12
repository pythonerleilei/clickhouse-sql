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
	EngineParams	[]string
	Comment			string
}

func (v *BaseAstClickhouseParserVisitor) VisitCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) interface{} {
	createDatabase := &CreateDatabase{}
	if ctx.EXISTS() != nil {
		createDatabase.FlagNotExists = true
	}

	createDatabase.DatabaseName = ctx.GetDb_name().GetText();
	if (ctx.GetCluster_name() != nil) {
		createDatabase.ClusterName = ctx.GetCluster_name().GetText()
	}
	
	if (ctx.GetEngine_name() != nil ){
		createDatabase.EngineName = ctx.GetEngine_name().GetText()
	}
	
	if ctx.GetEngine_params() != nil {
		var params []string
		for _, param := range ctx.GetEngine_params() {
			params = append(params, v.VisitParameter(param.(*parser.ParameterContext)).(string))
		}
		createDatabase.EngineParams = params;
	}
	if (ctx.COMMENT() != nil ){
		createDatabase.Comment = ctx.GetComment().GetText();
	}
	return createDatabase
}

func (v *BaseAstClickhouseParserVisitor) VisitParameter(ctx *parser.ParameterContext) interface{} {
	return ctx.GetText()
}