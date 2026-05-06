//go:build ignore

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0

	for {
		old := z
		z -= (z*z - x) / (2 * x)

		if math.Abs(z-old) < 1e-10 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}
