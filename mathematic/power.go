package mathematic

import "math"

func PowerH(base, power float64) float64 {
	sum := float64(1)
	for i := float64(0); i < power; i++ {
		sum *= base
	}

	return sum
}

func PowerMath(base, power float64) float64 {
	return math.Pow(base, power)
}

func PowerSlice(base, power float64) float64 {
	result := make([]float64, int(power))

	sum := float64(1)
	for range result {
		sum *= base
	}

	return sum
}

func PowerSlice2(base, power float64) float64 {
	result := make([]float64, int(power)/2)
	result2 := make([]float64, int(power)-(int(power)/2))

	sum := float64(1)
	for range result {
		sum *= base
	}
	for range result2 {
		sum *= base
	}

	return sum
}
