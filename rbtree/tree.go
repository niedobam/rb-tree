package rbtree

import "fmt"

type Tree struct {
    root *Node
    size int
}

func NewTree() *Tree {
    return &Tree{root: nil, size: 0}
}

func (tree *Tree) GetTreeSize() int {
    return tree.size
}

func (tree *Tree) GetRoot() *Node {
    return tree.root
}

func (tree *Tree) SetRoot(node *Node) {
    tree.root = node
}

func (tree *Tree) IncrementSize() {
    tree.size++
}

func (tree *Tree) DecrementSize() {
    tree.size--
}

func (tree *Tree) Insert(insertValue int) {
    if insertValue <= 0 {
        return
    }
    if tree.Search(insertValue) {
        return
    }
    if tree.root == nil {
        tree.root = NewNode(insertValue, nil, "B")
        tree.size++
        return
    }
    var insertNode *Node = tree.root
    tree.insertRecursive(insertNode, insertValue)
    tree.size++
}

func (tree *Tree) insertRecursive(insertNode *Node, insertValue int) {
    if insertValue > insertNode.value {
        if insertNode.right.value == 0 {
            insertNode.right = NewNode(insertValue, insertNode, "R")
            tree.balanceInsert(insertNode.right)
        } else {
            tree.insertRecursive(insertNode.right, insertValue)
        }
    } else if insertValue < insertNode.value {
        if insertNode.left.value == 0 {
            insertNode.left = NewNode(insertValue, insertNode, "R")
            tree.balanceInsert(insertNode.left)
        } else {
            tree.insertRecursive(insertNode.left, insertValue)
        }
    }
}

func (tree *Tree) DeleteNode(deleteValue int) {
    var deleteNode *Node = tree.FindNode(deleteValue)
    var child *Node
    var subTreeMin *Node
    var origColor string = deleteNode.color

    if deleteNode.left.value == 0 {
        child = deleteNode.right
        tree.transplant(deleteNode, deleteNode.right)
    } else if deleteNode.right.value == 0 {
        child = deleteNode.left
        tree.transplant(deleteNode, deleteNode.left)
    } else {
        subTreeMin = tree.getSubTreeMin(deleteNode.right)
        origColor = subTreeMin.color
        child = subTreeMin.right
        if subTreeMin.parent == deleteNode {
            child.parent = subTreeMin
        } else {
            tree.transplant(subTreeMin, subTreeMin.right)
            subTreeMin.right = deleteNode.right
            subTreeMin.right.parent = subTreeMin
        }
        tree.transplant(deleteNode, subTreeMin)
        subTreeMin.left = deleteNode.left
        subTreeMin.left.parent = subTreeMin
        subTreeMin.color = deleteNode.color
    }

    if origColor == "B" {
        tree.balanceDelete(child)
    }
    tree.size--
}

func (tree *Tree) Search(searchValue int) bool {
    var node *Node = tree.root
    var result bool = false

    if tree.root == nil {
        return false
    }

    for node.value != 0 {
        if node.value == searchValue {
            result = true
            break
        } else if searchValue > node.value {
            node = node.right
        } else {
            node = node.left
        }
    }
    return result
}

func (tree *Tree) FindNode(value int) *Node {
    var node *Node = tree.root
    var result *Node

    for node.value != 0 {
        if node.value == value {
            result = node
            break
        } else if value >= node.value {
            node = node.right
        } else {
            node = node.left
        }
    }
    return result
}

func (tree *Tree) PrintTree() {
    stringify(tree.root, 0)
}

func (tree *Tree) GetTreeMin() int {
    if tree.root == nil {
        fmt.Printf("\nBinary tree is empty")
        return 0
    }
    if tree.root.left.value == 0 && tree.root.right.value == 0 {
        return tree.root.value
    }
    var node *Node = tree.root
    if node.value != 0 {
        for node.left.value != 0 {
            node = node.left
        }
    }
    return node.value
}

func (tree *Tree) GetTreeMax() int {
    if tree.root == nil {
        return 0
    }
    if tree.root.left.value == 0 && tree.root.right.value == 0 {
        return tree.root.value
    }
    var node *Node = tree.root
    if node.value != 0 {
        for node.right.value != 0 {
            node = node.right
        }
    }
    return node.value
}
