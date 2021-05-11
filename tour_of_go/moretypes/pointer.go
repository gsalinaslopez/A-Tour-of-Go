package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p, p)
    *p = 21
    fmt.Println(i, p)

    p = &j
    *p = *p / 38
    fmt.Println(j, p)
}
