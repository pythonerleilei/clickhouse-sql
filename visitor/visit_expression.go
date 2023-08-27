package visitor

import "clickhouse-sql/parser"

func (v *BaseAstClickhouseParserVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	var name string
	if ctx.ID() == nil || len(ctx.ID().GetText()) == 0{
		name = "tuple"
	} else {
		name = ctx.ID().GetText()
	}
	expression := &Expression{}
	expression.name = name
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