package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

func Walk(t *tree.Tree, ch chan int) {
    //
    var visited_nodes_stack []*tree.Tree
    current_node := t
    for len(visited_nodes_stack) != 0 || current_node != nil {
        if current_node == nil {
            // pop the stack
            current_node = visited_nodes_stack[len(visited_nodes_stack) - 1]
            ch<-current_node.Value
            current_node = current_node.Right
            visited_nodes_stack = visited_nodes_stack[:len(visited_nodes_stack) - 1]
        } else {
            visited_nodes_stack = append(visited_nodes_stack, current_node)
            current_node = current_node.Left
        }
    }
    close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
    tree_one_ch := make(chan int)
    tree_two_ch := make(chan int)
    go Walk(t1, tree_one_ch)
    go Walk(t2, tree_two_ch)
    for {
        v1, ok1 := <-tree_one_ch
        v2, ok2 := <-tree_two_ch

        if !ok1 && !ok2 {
        // If reached the end without returning yet...
            return true
        } else if v1 != v2 {
        // If at any point the consumed tree nodes are different
            return false
        }

        // let the goroutines walk
    }
    return false
}

func main() {
    c := make(chan int)
    go Walk(tree.New(1), c)
    for i := range c {
        fmt.Println(i)
    }

    fmt.Println(Same(tree.New(1), tree.New(2)))
    fmt.Println(Same(tree.New(2), tree.New(2)))
}
