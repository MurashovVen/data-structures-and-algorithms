package main

import (
	"awesomeProject/collections"
)

func main() {
	tree := collections.NewBinTree()

	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(8)

	//tree.Insert(1)

	//tree.Insert(9)

	tree.Print()
}
