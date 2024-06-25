package ui

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "rb-tree/rbtree"
    "rb-tree/common"
)

func PrintHelp() {
    fmt.Printf("\n============================================\n")
    fmt.Printf("\nYou can enter following commands:\n")
    fmt.Printf("\n--------------------------------------------")
    fmt.Printf("\n\ninsert <value1> <value2> ... <value_n>\n\n--> will insert new node of a specified value into the tree\n\n")
    fmt.Printf("Example use case: --> insert 10\nThis will insert node with a value of 10 into the tree\n")
    fmt.Printf("\nExample use case --> insert 10 20 30\nThis will insert nodes with values of 10, 20 and 30 into the tree\n")
    fmt.Printf("\n--------------------------------------------")
    fmt.Printf("\n\ndelete(or remove) <value1> <value2> ... <value_n>\n\n--> will delete node(s) of a specified value(s) from the tree\n\n")
    fmt.Printf("Example use case: --> delete 10\nThis will delete node with a value of 10 into the tree\n")
    fmt.Printf("\nExample use case --> remove 10 20 30\nThis will remove nodes with values of 10, 20 and 30 from the tree\n")
    fmt.Printf("\n--------------------------------------------\n")
    fmt.Printf("\nmax(or get-max)\n\n--> will print maximum value contained in the tree\n\n")
    fmt.Printf("Example use case: --> max\nThis will print maximum value contained in the tree\n")
    fmt.Printf("\n--------------------------------------------\n")
    fmt.Printf("\nmin(or get-min)\n\n--> will print minimum value contained in the tree\n\n")
    fmt.Printf("Example use case: --> min\nThis will print minimum value contained in the tree")
    fmt.Printf("\n\n--------------------------------------------\n")
    fmt.Printf("\nroot(or get-root)\n\n--> will print value of the current tree root\n\n")
    fmt.Printf("Example use case: --> get-root\nThis will print value of the current tree root")
    fmt.Printf("\n\n--------------------------------------------\n")
    fmt.Printf("\nprint(or print-tree)\n\n--> will print current tree\n\n")
    fmt.Printf("Example use case: --> print-tree\nThis will print current tree")
    fmt.Printf("\n\n--------------------------------------------\n")
    fmt.Printf("\nsize(or get-size)\n\n--> will print current tree size\n\n")
    fmt.Printf("Example use case: --> get-size\nThis will print current tree size")
    fmt.Printf("\n\n--------------------------------------------\n")
    fmt.Printf("\ndelete-root(or remove-root)\n\n--> will remove root node from the current tree\n\n")
    fmt.Printf("Example use case: --> remove-root\nThis will remove current root node")
    fmt.Printf("\n\n--------------------------------------------\n")
    fmt.Printf("\ndelete-min(or remove-min)\n\n--> will remove node containing the minimum value from the tree\n\n")
    fmt.Printf("Example use case: --> remove-min\nThis will remove node containing the minimum value from the tree\n\n")
    fmt.Printf("\n--------------------------------------------\n")
    fmt.Printf("\ndelete-max(or remove-max)\n\n--> will remove node containing the maximum value from the tree\n\n")
    fmt.Printf("Example use case: --> remove-max\nThis will remove node containing the maximum value from the tree\n\n")
    fmt.Printf("--------------------------------------------\n")
    fmt.Printf("\nexit(or quit)\n\n--> will exit the program\n\n")
    fmt.Printf("Example use case: --> quit\nThis will terminate the program\n\n")
    fmt.Printf("--------------------------------------------\n")
    fmt.Printf("\nclear-tree(or delete-tree)\n\n--> will delete all nodes present in the current tree\n\n")
    fmt.Printf("Example use case: --> delete-tree\nThis will delete all nodes present in the current tree\n\n")
    fmt.Printf("--------------------------------------------\n")
    fmt.Printf("\nclear(or clc)\n\n--> will clear terminal screen\n\n")
    fmt.Printf("Example use case: --> clc\nThis will clear the terminal screen\n\n")
    fmt.Printf("--------------------------------------------\n")
}

func ExecuteCommand(myTree *rbtree.Tree, command string, values []int) {
    switch command {
    case "insert":
        for _, value := range values {
            if myTree.GetRoot() == nil {
                fmt.Printf("\n============================================\n")
                fmt.Printf("\nTree was empty\n")
                fmt.Printf("Tree after inserting -->%d<--\n\n", value)
                myTree.Insert(value)
                myTree.PrintTree()
                fmt.Printf("\n============================================\n")
                continue
            }

            if myTree.GetTreeSize() == 0 {
                fmt.Printf("\nTree was empty\n")
                fmt.Printf("Tree after inserting -->%d<--\n\n", value)
                myTree.Insert(value)
                myTree.PrintTree()
                fmt.Printf("\n============================================\n")
                continue
            }

            if myTree.Search(value) {
                fmt.Printf("\n-->Value %d is already in the tree<--\n", value)
                fmt.Printf("Insert values that are not already present in the tree\n")
                fmt.Printf("\n============================================\n")
                continue
            }

            fmt.Printf("\nTree before inserting -->%d<--\n\n", value)
            myTree.PrintTree()
            myTree.Insert(value)
            fmt.Printf("\n\nTree after inserting -->%d<--\n\n", value)
            myTree.PrintTree()
            fmt.Printf("\n============================================\n")
        }
    case "delete", "remove":
        for _, value := range values {
            if myTree.GetTreeSize() == 0 {
                fmt.Printf("\nTree is empty --> enter some values before running delete command\n")
                fmt.Printf("\n============================================\n")
                continue
            }

            if !myTree.Search(value) {
                fmt.Printf("\n-->Value %d is not present in the tree<--\n", value)
                fmt.Printf("\nEnter value that is present in the tree\n")
                fmt.Printf("\n============================================\n")
                continue
            }
            fmt.Printf("\nTree before deleting -->%d<--\n\n", value)
            myTree.PrintTree()
            myTree.DeleteNode(value)
            fmt.Printf("\n\nTree after deleting -->%d<--\n\n", value)
            myTree.PrintTree()
            fmt.Printf("\n============================================\n")
        }
    case "size", "get-size":
        fmt.Printf("\nCurrent tree size -->%d<--\n", myTree.GetTreeSize())
        fmt.Printf("\n============================================\n")
    case "max", "get-max":
        if myTree.GetTreeSize() == 0 {
            fmt.Printf("\nTree is empty --> enter some values before running get-max command\n")
            fmt.Printf("\n============================================\n")
            return
        }
        if myTree.GetRoot().GetValue() == 0 {
            fmt.Printf("\n============================================\n")
            fmt.Printf("\nTree is empty --> enter some values before running get-max command\n")
            fmt.Printf("\n============================================\n")
            return
        }
        if myTree.GetRoot().GetLeft().GetValue() == 0 && myTree.GetRoot().GetRight().GetValue() == 0 {
            fmt.Printf("\nRoot is the only value in the tree --> value -->%d<-- is both min and max value\n", myTree.GetTreeMax())
            fmt.Printf("\n============================================\n")
            return
        }
        fmt.Printf("\nMaximum value is -->%d<--\n", myTree.GetTreeMax())
        fmt.Printf("\n============================================\n")
    case "clear-tree", "delete-tree":
        if myTree.GetTreeSize() == 0 {
            fmt.Printf("\n-->Tree is already empty - there is nothing to clear<--\n")
            fmt.Printf("\nEnter some values first\n")
            fmt.Printf("\n============================================\n")
            return
        }
        fmt.Printf("\nTree before clearing:\n")
        myTree.PrintTree()
        for myTree.GetTreeSize() != 0 {
            myTree.DeleteNode(myTree.GetRoot().GetValue())
        }
        fmt.Printf("\n\n-->Tree was cleared successfully<--\n")
        fmt.Printf("\n============================================\n")
    case "min", "get-min":
        if myTree.GetTreeSize() == 0 {
            fmt.Printf("\nTree is empty --> enter some values before running get-min command\n")
            fmt.Printf("\n============================================\n")
            return
        }
        if myTree.GetRoot().GetValue() == 0 {
            fmt.Printf("\nTree is empty --> enter some values before running get-min command\n")
            fmt.Printf("\n============================================\n")
            return
        }
        if myTree.GetRoot().GetLeft().GetValue() == 0 && myTree.GetRoot().GetRight().GetValue() == 0 {
            fmt.Printf("\nRoot is the only value in the tree --> value -->%d<-- is both min and max value\n", myTree.GetTreeMax())
            fmt.Printf("\n============================================\n")
            return
        }
        fmt.Printf("\nMinimum value is -->%d<--\n", myTree.GetTreeMin())
        fmt.Printf("\n============================================\n")
    case "root", "get-root":
        if myTree.GetRoot() == nil {
            fmt.Printf("\nTree is empty --> root is nil\n")
            fmt.Printf("\n============================================\n")
            return
        }
        if myTree.GetRoot().GetValue() == 0 {
            fmt.Printf("\nTree is empty --> root is nil\n")
            fmt.Printf("\n============================================\n")
            return
        }
        fmt.Printf("\nRoot value is -->%d<--\n", myTree.GetRoot().GetValue())
        fmt.Printf("\n============================================\n")
    case "print", "print-tree":
        if myTree.GetTreeSize() == 0 {
            fmt.Printf("\nTree is empty --> Enter some values before running the print-tree command\n")
            fmt.Printf("\n============================================\n")
            return
        }
        fmt.Printf("\n")
        myTree.PrintTree()
        fmt.Printf("\n============================================\n")
    case "clear", "clc":
        common.CallClear()
    case "help":
        PrintHelp()
    default:
        fmt.Printf("\n============================================\n")
        fmt.Printf("\n-->Unknown command - enter valid command<--\nType 'help' to list all available commands\n")
        fmt.Printf("\n============================================\n")
    }
}

func GetUserInput() (string, []int) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("\n>> ")
    input, _ := reader.ReadString('\n')
    inputSplit := strings.Fields(input)
    if len(inputSplit) == 0 {
        return "", nil
    }
    var command string = inputSplit[0]
    var values []int
    for i := 1; i < len(inputSplit); i++ {
        currentValue, err := strconv.Atoi(inputSplit[i])
        if err == nil && currentValue > 0 {
            values = append(values, currentValue)
        } else if err != nil {
            fmt.Printf("\n============================================\n")
            fmt.Printf("\nFollowing error occurred: %s", err)
            fmt.Printf("\n\n-->Enter valid non-float numeric values that are larger than zero\n")
            fmt.Printf("\n============================================\n")
            break
        } else {
            fmt.Printf("\n============================================\n")
            fmt.Printf("\nEntered value was smaller or equal to zero")
            fmt.Printf("\n\n-->Enter valid non-float numeric values that are larger than zero\n")
            fmt.Printf("\n============================================\n")
            break
        }
    }
    return command, values
}
