# rb-tree

![Screenshot from 2024-06-25 13-21-58](https://github.com/niedobam/rb-tree/assets/127039716/21303475-83ab-4863-a909-6f0eb54a4581)

## General Information
An implementation of a [Red-black tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) data structure with a command-line interface for interacting with the tree written in Go. This project demonstrates the creation, insertion, deletion, and traversal of a Red-Black Tree.

## Prerequisites
* [Go (Golang)](https://go.dev/) programming language

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/niedobam/rb-tree.git
    cd rb-tree
    ```

2. Build the project:

    ```sh
    go build -o rb-tree ./cmd
    ```

## Usage

Run the compiled program:

```sh
./rb-tree
```
## Available Commands
* insert \<value1> \<value2> ... \<value_n> - Insert new nodes with the specified values into the tree.

* delete \<value1> \<value2> ... \<value_n> - Delete nodes with the specified values from the tree.

* print (or print-tree) - Print the current state of the tree.

* max (or get-max) - Print the maximum value in the tree.

* min (or get-min) - Print the minimum value in the tree.

* root - Print the root value of the tree.

* size (or get-size) - Print the current state of the tree.

* delete-root (or remove-root) - Deletes current root node.

* delete-max (or remove-max) - Deletes node with the max value.

* delete-min (or remove-min) - Deletes node with the min value.

* exit (or quit) - Exits the program.

* clear-tree (or delete-tree) - Removes all nodes from the tree.

* clear (or clc) - Clears the screen.
