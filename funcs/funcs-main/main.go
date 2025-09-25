package main

import (
	ds "github.com/pseudoelement/go-sandbox/funcs/data-structures"
)

func main() {
	tree := ds.NewHtmlTree()

	node2 := ds.NewTreeNode(2, "html-2")
	node3 := ds.NewTreeNode(3, "html-3")

	tree.Root.AppendChild(node2)
	tree.Root.AppendChild(node3)

	node4 := ds.NewTreeNode(4, "html-4")
	node5 := ds.NewTreeNode(5, "html-5")
	node6 := ds.NewTreeNode(6, "html-6")

	node2.AppendChild(node4)
	node2.AppendChild(node5)
	node2.AppendChild(node6)

	node7 := ds.NewTreeNode(7, "html-7")
	node8 := ds.NewTreeNode(8, "html-8")
	node9 := ds.NewTreeNode(9, "html-9")

	node4.AppendChild(node7)

	node5.AppendChild(node8)
	node5.AppendChild(node9)

	node10 := ds.NewTreeNode(10, "html-10")
	node11 := ds.NewTreeNode(11, "html-11")

	node7.AppendChild(node10)
	node7.AppendChild(node11)

	// node, stepCount := tree.FindByIdDFSRecursive(tree.Root, "html-8", new(int))

	// fmt.Printf("node %+v, stepCount %v\n", node, stepCount)

}
