package main

import (
	"learnGo/tree"
	"fmt"
	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

//组合的形式拓展已有类型
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
func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(10000000000000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000000000000))
	//10000000000000






}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	//new()直接就是地址，所以不用&
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println("root1")
	var root1 tree.Node
	root1.Left = &tree.Node{1, nil, nil}
	root1.Value = 3
	root1.Right = &tree.Node{2, nil, nil}

	root1.Print()
	fmt.Print()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	testSparse()
}
