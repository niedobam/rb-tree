package rbtree

import "fmt"

type Node struct {
    value  int
    left   *Node
    right  *Node
    parent *Node
    color  string
}

func NewNode(value int, parent *Node, color string) *Node {
    var nillNode *Node = &Node{value: 0, left: nil, right: nil, parent: parent, color: "B"}
    return &Node{value: value, left: nillNode, right: nillNode, parent: parent, color: color}
}

func (n *Node) GetValue() int {
    return n.value
}

func (n *Node) GetLeft() *Node {
    return n.left
}

func (n *Node) GetRight() *Node {
    return n.right
}

func stringify(node *Node, level int) {
    if node.value != 0 {
        format := ""
        for i := 0; i < level; i++ {
            format += "       "
        }
        format += "---["
        level += 1
        stringify(node.right, level)
        fmt.Printf(format+"%s-%d\n", node.color, node.value)
        stringify(node.left, level)
    }
}
