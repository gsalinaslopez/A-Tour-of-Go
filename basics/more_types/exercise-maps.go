package main

import (
    "golang.org/x/tour/wc"
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    for _, word := range strings.Fields(s) {
        _, isWordPresent := m[word]
        if (!isWordPresent) {
            m[word] = 0
        }
        m[word] = m[word] + 1
    }
    fmt.Println(m)
    return m
}

func main() {
    // WordCount("  foo bar  baz   ")
    wc.Test(WordCount)
}
