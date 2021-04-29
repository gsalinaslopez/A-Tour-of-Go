package main

import (
    "fmt"

    "rsc.io/quote"

    "example.com/greetings"
)


func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    fmt.Println(greetings.Hello("Klaus"))
}
