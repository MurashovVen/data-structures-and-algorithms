package collections

import "fmt"

type (
	linkedList struct {
		head *linkedListNode

		tail *linkedListNode
	}

	linkedListNode struct {
		value    int
		next     *linkedListNode
		previous *linkedListNode
	}
)

func newLinkedListNode(value int) *linkedListNode {
	return &linkedListNode{value: value}
}

func (l *linkedList) insertEnd(node *linkedListNode) {
	if l.tail == nil {
		l.head = node
		l.tail = node

		return
	}

	if l.tail == l.head {
		l.tail = node
		l.tail.previous = l.head
		l.head.next = l.tail

		return
	}

	node.previous = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *linkedList) insertStart(node *linkedListNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	if l.tail == l.head {
		l.head = node
		l.tail.previous = l.head
		l.head.next = l.tail

		return
	}

	node.next = l.head
	l.head.previous = node
	l.head = node
}

func (l *linkedList) deleteEnd() {
	if l.tail == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil

		return
	}

	l.tail = l.tail.previous
}

func (l *linkedList) deleteStart() {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil

		return
	}

	l.head = l.head.next
}

func (l *linkedList) print() {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr != nil {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
	fmt.Println()
}

func (l *linkedList) printReverse() {
	if l.tail == nil {
		return
	}

	curr := l.tail
	for curr != nil {
		fmt.Print(curr.value, " ")
		curr = curr.previous
	}
	fmt.Println()
}
