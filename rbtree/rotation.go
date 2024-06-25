package rbtree

func (tree *Tree) rotateLeft(node *Node) {
    var rightChild *Node = node.right
    node.right = rightChild.left
    if rightChild.left.value != 0 {
        rightChild.left.parent = node
    }
    rightChild.parent = node.parent
    if node.parent == nil {
        tree.root = rightChild
    } else if node == node.parent.left {
        node.parent.left = rightChild
    } else {
        node.parent.right = rightChild
    }
    rightChild.left = node
    node.parent = rightChild
}

func (tree *Tree) rotateRight(node *Node) {
    var leftChild *Node = node.left
    node.left = leftChild.right
    if leftChild.right.value != 0 {
        leftChild.right.parent = node
    }
    leftChild.parent = node.parent
    if node.parent == nil {
        tree.root = leftChild
    } else if node == node.parent.right {
        node.parent.right = leftChild
    } else {
        node.parent.left = leftChild
    }
    leftChild.right = node
    node.parent = leftChild
}
