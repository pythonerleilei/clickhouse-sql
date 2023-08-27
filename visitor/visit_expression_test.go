package visitor

import "testing"

func TestVisitParameterNumber(t *testing.T) {
	sql := "18"
	astVisitor := NewBaseAstClickhouseParserVisitor(sql)
	
	result := astVisitor.Visit(astVisitor.Parser.Parameter())

	le, ok := result.(*LiteralExpression)
	if !ok {
		t.Fatalf("paser parameter[%v] failed", sql)
	}

	if le.GetValue() != "18" {
		t.Fatalf("paser parameter[%v] failed, expected[%v], get[%v]", sql, sql, le.GetValue())
	}
}

func TestVisitParameterString(t *testing.T) {
	sql := `'host:port'`
	astVisitor := NewBaseAstClickhouseParserVisitor(sql)
	
	result := astVisitor.Visit(astVisitor.Parser.Parameter())

	le, ok := result.(*LiteralExpression)
	if !ok {
		t.Fatalf("paser parameter[%v] failed", sql)
	}

	if le.GetValue() != "'host:port'" {
		t.Fatalf("paser parameter[%v] failed, expected[%v], get[%v]", sql, sql, le.GetValue())
	}
}

func TestVisitParameters(t *testing.T) {
	sql := `'host:port', 18`
	astVisitor := NewBaseAstClickhouseParserVisitor(sql)
	
	result := astVisitor.Visit(astVisitor.Parser.Parameters())

	es, ok := result.([]TreeNode)
	if !ok {
		t.Fatalf("paser parameter[%v] failed", sql)
	}

	
	if es[0].(*LiteralExpression).GetValue() != "'host:port'" {
		t.Fatalf("paser parameter[%v] failed, expected[%v], get[%v]", sql, "'host:port'", es[0].(*LiteralExpression).GetValue())
	}

	if es[1].(*LiteralExpression).GetValue() != "18" {
		t.Fatalf("paser parameter[%v] failed, expected[%v], get[%v]", sql, "18", es[0].(*LiteralExpression).GetValue())
	}
}

func TestVisitExpression(t *testing.T) {
	sql := "Map(String, Nullable(FixString(32)))"
	astVisitor := NewBaseAstClickhouseParserVisitor(sql)

	result := astVisitor.Visit(astVisitor.Parser.Expression())

	e, ok := result.(*Expression)
	if !ok {
		t.Fatalf("parse expression[%v] failed", sql)
	}
	
	if e.Name() != "Map" {
		t.Fatalf("parse expression[%v] failed, expected[%v], get[%v]", sql, "Map", e.Name())
	}
}

func TestVisitExpressionTuple(t *testing.T) {
	sql := "(String, Nullable(FixString(32))"
	astVisitor := NewBaseAstClickhouseParserVisitor(sql)

	result := astVisitor.Visit(astVisitor.Parser.Expression())

	e, ok := result.(*Expression)
	if !ok {
		t.Fatalf("parse expression[%v] failed", sql)
	}
	
	if e.Name() != "tuple" {
		t.Fatalf("parse expression[%v] failed, expected[%v], get[%v]", sql, "Map", e.Name())
	}
}