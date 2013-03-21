package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for tmp := 0.0; math.Abs(tmp - z) > 0.00000001; {
        tmp = z
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}

