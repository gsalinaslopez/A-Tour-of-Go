package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int

func fibonacci() func() int {
    var i int = int(-1)
    var j int = int(1)
    return func() int {
        r := i + j
        i = j
        j = r
        return r
    }
}

// i = 1, j = 0
// r = 1 + 0 = 1, i = 0, j = 1
// r = 0 + 1 = 1, i = 1, j = 1

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
