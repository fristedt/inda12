package main

import (
	"fmt"
	"math/cmplx"
)

const limit float64 = 10e-100

func Cbrt(x complex128) complex128 {
	z := complex(1, 1)
	i := complex(3, 3)
	for cmplx.Abs(i - z) > limit {
		i = z
		z = z - (1 / 3) * (2 * z + (x / cmplx.Pow(z, 2)))
	}
	return z		
}

func main() {
	fmt.Println(Cbrt(8))
}
