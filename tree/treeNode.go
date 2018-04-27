package tree

import "fmt"

/*type tree struct {
	value string
	name string
}*/
type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

//两种写法都可以
func (node Node) Print() {
	fmt.Println(node.Value)
}
func print(node Node) {
	fmt.Print(node.Value, "")
}
func (node *Node) SetValue(Value int) {
	if node ==nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
	}

	node.Value = Value

}

