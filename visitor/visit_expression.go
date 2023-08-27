package visitor

import "clickhouse-sql/parser"

func (v *BaseAstClickhouseParserVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	expression := &Expression{}
	expression.name = ctx.ID().GetText()
	if(ctx.Parameters() != nil) {
		var children = v.Visit(ctx.Parameters()).([]TreeNode)
		expression.SetChildren(children)
	}
	return expression
}

func (v *BaseAstClickhouseParserVisitor) VisitParameters(ctx *parser.ParametersContext) interface{} {
	var expressions []TreeNode
	params := ctx.AllParameter()
	for _, param := range params {
		expressions = append(expressions, v.Visit(param).(TreeNode))
	}
	return expressions
}

func (v *BaseAstClickhouseParserVisitor) VisitParameter(ctx *parser.ParameterContext) interface{} {
	if (ctx.NUMBER() != nil || ctx.STRING() != nil) {
		return NewLiteralExpression(ctx.GetText())
	}
	return v.Visit(ctx.Expression())
}