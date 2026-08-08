package main

import (
	"fmt"
	"math"
)

func earlyEqual(a, b, epsilon float32) bool {
	absA := float32(math.Abs(float64(a)))
	absB := float32(math.Abs(float64(b)))
	diff := float32(math.Abs(float64(a - b)))

	if a == b {
		// Shortcut, handles infinities
		return true
	} else if a == 0 || b == 0 || (absA+absB < float32(math.SmallestNonzeroFloat32)) {
		// a or b is zero or both are extremely close to it
		// relative error is less meaningful here
		return diff < (epsilon * float32(math.SmallestNonzeroFloat32))
	} else {
		// Use relative error
		sum := absA + absB
		if sum > float32(math.MaxFloat32) {
			sum = float32(math.MaxFloat32)
		}

		return diff/sum < epsilon
	}
}

func main() {
	a := float32(1.0)
	b := float32(1.000001)
	epsilon := float32(0.00001)

	fmt.Println(earlyEqual(a, b, epsilon))
}
