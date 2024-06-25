package main

import (
    "fmt"
    "rb-tree/rbtree"
    "rb-tree/ui"
)

func main() {
    var myTree *rbtree.Tree = rbtree.NewTree()

    fmt.Printf("\nWelcome to Red-Black Tree program!\n")
    fmt.Printf("=======================================================")
    fmt.Printf("\nTo display information about all possible commands type 'help'\n")
    fmt.Printf("Type a command or type 'exit' to terminate program\n\n")

    for {
        command, values := ui.GetUserInput()

        if command == "exit" || command == "quit" {
            fmt.Println("\n\nExiting program")
            break
        }

        ui.ExecuteCommand(myTree, command, values)
    }

    fmt.Printf("\n============================================\n")
    fmt.Println("-->Program exited successfully<--")
}
