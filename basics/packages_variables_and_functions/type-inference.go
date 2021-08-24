package main

import "fmt"

func main() {
    v := 42
    fmt.Printf("type %T\n", v)
    c := 3.142
    fmt.Printf("type %T\n", c)
    b := 0.867 + 0.5i
    fmt.Printf("type %T\n", b)
}
