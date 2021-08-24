package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return float64(x), ErrNegativeSqrt(x)
    }

    z := float64(1)
    for i := 0; i < 10; i++ {
        fmt.Printf("z: %v at iteration: %v\n", z*z, i)
        z -= (z*z - x) / (2*z)
    }
    return z, nil
}

func main() {
    i := ErrNegativeSqrt(-4)
    fmt.Println(i)
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
