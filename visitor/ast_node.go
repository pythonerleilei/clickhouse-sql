package visitor

type TreeNode interface {
	Children() []TreeNode
	SetChildren(children []TreeNode)
	Name() string
}


type Expression struct {
	children	[]*Expression
	name		string
}


func (e *Expression) Children() []TreeNode {
	if e.children == nil || len(e.children) == 0 {
		return nil
	}
	var nodes []TreeNode
	for _, e := range e.children {
		nodes = append(nodes, e)
	}
	return nodes
}


func (e *Expression) SetChildren(children []TreeNode) {
	if len(children) == 0 {
		return
	}
	var expressions []*Expression
	for _, c := range children {
		if _, ok := c.(*Expression); ok {
			expressions = append(expressions, c.(*Expression))
		}
	}
	e.children = expressions
}

func (e *Expression) SetName(name string) {
	e.name = name
}

func (e *Expression) Name() string {
	return e.name
}

type LeafExpression struct {
	Expression
}

func (le *LeafExpression) Children() []TreeNode {
	return nil;
}

func (le *LeafExpression) SetChildren(children []TreeNode) {
	
}

type LiteralExpression struct {
	LeafExpression
	value			string
}

func NewLiteralExpression(value string) *LiteralExpression {
	return &LiteralExpression{value: value}
}

func (le *LiteralExpression) GetValue() string {
	return le.value
}

func (le *LiteralExpression) SetValue(value string) {
	le.value = value
}


type UnaryExpression struct {
	Expression
	child		*Expression
}

func (ue *UnaryExpression) Chlid() *Expression {
	return ue.child
}

func (ue *UnaryExpression) SetChlid(e *Expression){
	ue.child = e
}

func (ue *UnaryExpression) Children() []TreeNode {
	var nodes []TreeNode
	nodes = append(nodes, ue.child)
	return nodes;
}


func (ue *UnaryExpression) SetChildren(children []TreeNode) {
	ue.child = children[0].(*Expression)
}





