package collections

import "github.com/shivamMg/ppds/tree"

type (
	BinTreeNode struct {
		Value int

		ChildLeft  *BinTreeNode
		ChildRight *BinTreeNode
	}

	BinTreeVisitNodeFunc func(node *BinTreeNode)
)

var _ tree.Node = (*BinTreeNode)(nil)

func (b *BinTreeNode) Data() interface{} {
	if b == nil {
		return nil
	}

	return b.Value
}

func (b *BinTreeNode) Children() []tree.Node {
	if b == nil {
		return nil
	}

	var nodes []tree.Node

	if b.ChildLeft != nil {
		nodes = append(nodes, b.ChildLeft)
	}

	if b.ChildRight != nil {
		nodes = append(nodes, b.ChildRight)
	}

	return nodes
}
