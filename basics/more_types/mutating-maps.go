package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["answer"] = 42
    fmt.Println("The value:", m["answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    fmt.Println("Map print with fmt: ", m)

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    delete(m, "Aanswer")

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
