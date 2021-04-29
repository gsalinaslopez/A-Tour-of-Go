package main

import (
    "fmt"
    "log"
    "rsc.io/quote"

    "example.com/greetings"
)


func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
