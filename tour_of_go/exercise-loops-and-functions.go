package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for i := 0; i < 10; i++ {
        fmt.Printf("z: %v at iteration: %v\n", z*z, i)
        z -= (z*z - x) / (2*z)
    }
    return z
}

func Sqrt2(x float64) float64 {
    var z float64 = 1.0
    var z_prev float64 = 0.0
    i := 0
    for math.Abs(z_prev - z) > 0.000000000000001 {
        fmt.Printf("z: %v, z^2: %v at iteration: %v\n", z, z*z, i)
        z_prev = z
        z -= (z*z - x) / (2*z)
        i++
        fmt.Println(math.Abs(z_prev - z))
    }
    return z
}

func main() {
    fmt.Println(Sqrt2(2), math.Sqrt(2))
    fmt.Println(Sqrt2(3), math.Sqrt(3))
}
