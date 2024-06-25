package rbtree

func (tree *Tree) balanceInsert(insertedNode *Node) {
    var insertedNodeGrandparent *Node = insertedNode.parent.parent
    var temp *Node = NewNode(0, nil, "B")

    for insertedNode.parent.color == "R" {
        if insertedNodeGrandparent.right != nil && insertedNode.parent == insertedNodeGrandparent.right {
            temp = insertedNodeGrandparent.left
            if temp != nil && temp.color == "R" {
                temp.color = "B"
                insertedNode.parent.color = "B"
                insertedNodeGrandparent.color = "R"
                insertedNode = insertedNodeGrandparent
            } else {
                if insertedNode == insertedNode.parent.left {
                    insertedNode = insertedNode.parent
                    tree.rotateRight(insertedNode)
                }
                insertedNode.parent.color = "B"
                insertedNodeGrandparent.color = "R"
                tree.rotateLeft(insertedNodeGrandparent)
            }
        } else {
            if insertedNodeGrandparent.right != nil && insertedNodeGrandparent.right.color == "R" {
                temp = insertedNodeGrandparent.right
                temp.color = "B"
                insertedNode.parent.color = "B"
                insertedNodeGrandparent.color = "R"
                insertedNode = insertedNodeGrandparent
            } else {
                if insertedNode == insertedNode.parent.right {
                    insertedNode = insertedNode.parent
                    tree.rotateLeft(insertedNode)
                }
                insertedNode.parent.color = "B"
                insertedNodeGrandparent.color = "R"
                tree.rotateRight(insertedNodeGrandparent)
            }
        }

        if insertedNode == tree.root {
            break
        }
    }
    tree.root.color = "B"
}

func (tree *Tree) balanceDelete(balancedNode *Node) {
    var sibling *Node
    for balancedNode != tree.root && balancedNode.color == "B" {
        if balancedNode == balancedNode.parent.left {
            sibling = balancedNode.parent.right
            if sibling.color == "R" {
                sibling.color = "B"
                balancedNode.parent.color = "R"
                tree.rotateLeft(balancedNode.parent)
                sibling = balancedNode.parent.right
            }
            if sibling.left.color == "B" && sibling.right.color == "B" {
                sibling.color = "R"
                balancedNode = balancedNode.parent
            } else {
                if sibling.right.color == "B" {
                    sibling.left.color = "B"
                    sibling.color = "R"
                    tree.rotateRight(sibling)
                    sibling = balancedNode.parent.right
                }
                sibling.color = balancedNode.parent.color
                balancedNode.parent.color = "B"
                sibling.right.color = "B"
                tree.rotateLeft(balancedNode.parent)
                balancedNode = tree.root
            }
        } else {
            sibling = balancedNode.parent.left
            if sibling.color == "R" {
                sibling.color = "B"
                balancedNode.parent.color = "R"
                tree.rotateRight(balancedNode.parent)
                sibling = balancedNode.parent.left
            }
            if sibling.left.color == "B" && sibling.right.color == "B" {
                sibling.color = "R"
                balancedNode = balancedNode.parent
            } else {
                if sibling.left.color == "B" {
                    sibling.right.color = "B"
                    sibling.color = "R"
                    tree.rotateLeft(sibling)
                    sibling = balancedNode.parent.left
                }
                sibling.color = balancedNode.parent.color
                balancedNode.parent.color = "B"
                sibling.left.color = "B"
                tree.rotateRight(balancedNode.parent)
                balancedNode = tree.root
            }
        }
    }
    balancedNode.color = "B"
}

func (tree *Tree) transplant(deleteNode *Node, deleteNodeChild *Node) {
    if deleteNode.parent == nil {
        tree.root = deleteNodeChild
    } else if deleteNode == deleteNode.parent.left {
        deleteNode.parent.left = deleteNodeChild
    } else {
        deleteNode.parent.right = deleteNodeChild
    }
    deleteNodeChild.parent = deleteNode.parent
}

func (tree *Tree) getSubTreeMin(node *Node) *Node {
    for node.left.value != 0 {
        node = node.left
    }
    return node
}
