package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
    r := ""
    for i, b := range ip {
        r += fmt.Sprint(b)
        if i < len(ip) - 1 {
            r += "."
        }
    }
    return r
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
