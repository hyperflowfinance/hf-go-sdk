package hyperflow

type argument struct {
	name     string
	gqlType  string
	value    any
	optional bool
}

type selectionNode struct {
	name              string
	fragment          string
	args              []argument
	children          []*selectionNode
	requiresSelection bool
}

type Document struct {
	operation string
	root      *selectionNode
}

func (n *selectionNode) field(name string, args ...argument) *selectionNode {
	child := &selectionNode{name: name, args: args}
	n.children = append(n.children, child)
	return child
}

func (n *selectionNode) object(name string, args ...argument) *selectionNode {
	child := n.field(name, args...)
	child.requiresSelection = true
	return child
}

func (n *selectionNode) inlineFragment(typeName string) *selectionNode {
	child := &selectionNode{fragment: typeName}
	n.children = append(n.children, child)
	return child
}

type optionalValue interface {
	graphOptional() (any, bool)
}
