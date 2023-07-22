package collections

import "github.com/shivamMg/ppds/tree"

type (
	BinTree struct {
		root *BinTreeNode
	}
)

func NewBinTree() *BinTree {
	return &BinTree{}
}

// Insert -
// O (log n)
func (t *BinTree) Insert(value int) {
	leafNew := &BinTreeNode{Value: value}

	if t.root == nil {
		t.root = leafNew
		return
	}

	curr := t.root
loop:
	for {
		switch {
		case value > curr.Value && curr.ChildRight == nil:
			curr.ChildRight = leafNew

			break loop

		case value < curr.Value && curr.ChildLeft == nil:
			curr.ChildLeft = leafNew

			break loop

		case value > curr.Value:
			curr = curr.ChildRight

		case value < curr.Value:
			curr = curr.ChildLeft
		}
	}
}

// O (log n)
func (t *BinTree) find(value int) (node *BinTreeNode, parent *BinTreeNode, isLeftChild bool, err error) {
	if t.root == nil {
		return nil, nil, false, ErrEmpty
	}

	var (
		curr            = t.root
		parentCurr      *BinTreeNode
		isCurrLeftChild bool
	)
	for {
		switch {
		case value > curr.Value:
			if curr.ChildRight == nil {
				return nil, nil, false, NotFound
			}

			curr, parentCurr, isCurrLeftChild = curr.ChildRight, curr, false

		case value < curr.Value:
			if curr.ChildLeft == nil {
				return nil, nil, false, NotFound
			}

			curr, parentCurr, isCurrLeftChild = curr.ChildLeft, curr, true

		case curr.Value == value:
			return curr, parentCurr, isCurrLeftChild, nil
		}
	}
}

// Delete -
// O (log n)
func (t *BinTree) Delete(value int) error {
	if t.root == nil {
		return ErrEmpty
	}

	curr, currParent, isCurrLeftChild, err := t.find(value)
	if err != nil {
		return err
	}

	var successor *BinTreeNode

	switch {
	case curr.ChildRight == nil && curr.ChildLeft == nil:
		successor = nil

	case curr.ChildRight == nil:
		successor = curr.ChildLeft

	case curr.ChildLeft == nil:
		successor = curr.ChildRight

	case curr.ChildLeft != nil && curr.ChildRight != nil:
		successorParent := curr
		successor = curr.ChildRight
		isSuccessorRightChild := successor.ChildLeft == nil
		for successor.ChildLeft != nil {
			successor, successorParent = successor.ChildLeft, successor
		}

		if isSuccessorRightChild {
			successorParent.ChildRight = successor.ChildRight
		} else {
			successorParent.ChildLeft = successor.ChildRight
		}

		// children switch
		successor.ChildLeft = curr.ChildLeft
		successor.ChildRight = curr.ChildRight
	}

	// parent switch
	switch {
	case currParent == nil:
		t.root = successor

	case isCurrLeftChild:
		currParent.ChildLeft = successor

	case !isCurrLeftChild:
		currParent.ChildRight = successor
	}

	return nil
}

func (t *BinTree) Print() {
	tree.Print(t.root)
}

func (t *BinTree) TraverseInOrder(visit BinTreeVisitNodeFunc) {
	t.traverseInOrder(t.root, visit)
}

func (t *BinTree) traverseInOrder(node *BinTreeNode, visit BinTreeVisitNodeFunc) {
	if node != nil {
		t.traverseInOrder(node.ChildLeft, visit)
		visit(node)
		t.traverseInOrder(node.ChildRight, visit)
	}
}

func (t *BinTree) TraversePreOrder(visit BinTreeVisitNodeFunc) {
	t.traversePreOrder(t.root, visit)
}

func (t *BinTree) traversePreOrder(node *BinTreeNode, visit BinTreeVisitNodeFunc) {
	if node != nil {
		visit(node)
		t.traversePreOrder(node.ChildLeft, visit)
		t.traversePreOrder(node.ChildRight, visit)
	}
}

func (t *BinTree) TraversePostOrder(visit BinTreeVisitNodeFunc) {
	t.traversePostOrder(t.root, visit)
}

func (t *BinTree) traversePostOrder(node *BinTreeNode, visit BinTreeVisitNodeFunc) {
	if node != nil {
		t.traversePostOrder(node.ChildLeft, visit)
		t.traversePostOrder(node.ChildRight, visit)
		visit(node)
	}
}
