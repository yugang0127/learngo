package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	root.Value = 3
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)


	fmt.Println("In-order traversal.")
	root.Traverse()
	fmt.Println()

	fmt.Println("post-order tranversal .")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		maxNode = node.Value
	}
	fmt.Println("Max node value:", maxNode)
	root.Right.Left.SetValue(6)
	root.Right.Left.Print()
	fmt.Println()

	root.Print()
	root.SetValue(100)

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
}